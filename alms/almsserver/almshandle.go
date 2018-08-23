package almsserver

import (
	"fmt"
	"steve/alms/data"
	"steve/alms/packsack/packsack_gold"
	client_alms "steve/client_pb/alms"
	"steve/client_pb/msgid"
	"steve/common/constant"
	"steve/external/configclient"
	"steve/external/goldclient"
	"steve/server_pb/gold"
	"steve/structs/exchanger"
	"steve/structs/proto/gate_rpc"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

// HandleGetAlmsReq 获取救济金请求(玩家申请救济金)，如果期间配置发生改变，发送新的配置
func HandleGetAlmsReq(playerID uint64, header *steve_proto_gaterpc.Header, req client_alms.AlmsGetGoldReq) (rspMsg []exchanger.ResponseMsg) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "HandleGetAlmsReq",
		"playerID":  playerID,
		"request":   req,
	})
	// 返回消息
	response := &client_alms.AlmsGetGoldRsp{
		Result: proto.Bool(true),
	}
	rspMsg = []exchanger.ResponseMsg{
		exchanger.ResponseMsg{
			MsgID: uint32(msgid.MsgID_ALMS_GET_GOLD_RSP),
			Body:  response,
		},
	}
	// 校验玩家背包是否有金豆 有返回false
	Pkgold, err := packsack_gold.GetGoldMgr().GetGold(playerID)
	if err != nil {
		response.Result = proto.Bool(false)
		entry.WithError(err).Debugln("背包金币获取")
		return
	}
	if Pkgold != 0 {
		response.Result = proto.Bool(false)
		entry.Debugln("背包金币不为0")
		return
	}
	version := req.GetVersion() // 版本号
	// 获取配置用于验证
	ac, err := data.GetAlmsConfigByPlayerID(playerID)
	if err != nil {
		entry.WithError(err).Errorln(fmt.Sprintf("根据玩家ID获取救济金配置失败 playerID(%d)", playerID))
		response.Result = proto.Bool(false)
		return
	}
	response.NewVersion = proto.Int32(int32(ac.Version))
	// 版本不对应,配置发生改变,发送新的配置信息
	if version != int32(ac.Version) {
		entry.Debugln(fmt.Sprintf("版本号不一致 currVersion(%d)", ac.Version))
		newAlmsConfig := &client_alms.AlmsConfig{
			AlmsGetNorm:      proto.Int64(ac.GetNorm),                 // 救济金
			AlmsGetTimes:     proto.Int32(int32(ac.GetTimes)),         // 配置每个玩家可以领取次数
			AlmsGetNumber:    proto.Int64(ac.GetNumber),               // 领取数量
			AlmsCountDonw:    proto.Int32(int32(ac.AlmsCountDonw)),    //救济倒计时
			DepositCountDonw: proto.Int32(int32(ac.DepositCountDonw)), //救济倒计时
		}
		response.NewAlmsConfig = newAlmsConfig
	}

	//是否还有领取次数
	if ac.PlayerGotTimes >= 3 {
		entry.Debugln(fmt.Sprintf("领取次数已满 times(%v) ", ac.PlayerGotTimes))
		response.Result = proto.Bool(false)
		return
	}

	reqType := req.GetReqType()
	// 请求类型
	isExist := map[client_alms.AlmsReqType]bool{
		client_alms.AlmsReqType_LOGIN:    true, // 登录
		client_alms.AlmsReqType_SELECTED: true, // 选场
		client_alms.AlmsReqType_INGAME:   true, // 游戏中
	}[reqType]

	if !isExist {
		entry.Debugln(fmt.Sprintf("救济金请求类型不存在 reqType(%d)", reqType))
		response.Result = proto.Bool(false)
		return
	}

	//从金币服获取玩家身上金币，
	playerGold, err := goldclient.GetGold(playerID, int16(gold.GoldType_GOLD_COIN))
	if err != nil {
		entry.WithError(err).Debugln(fmt.Sprintf("playerID(%d) - 从金币服获取玩家身上金币失败", playerID))
		response.Result = proto.Bool(false)
		return
	}

	// 除了游戏中才判断是否救济金达标，必须低于救济线
	if reqType != client_alms.AlmsReqType_INGAME && playerGold >= ac.GetNorm {
		entry.Debugln(fmt.Sprintf("玩家身上的金币没有达到救济线以内  playerGold(%d) - GetNorm(%d) reqType(%v)", playerGold, ac.GetNorm, reqType))
		response.Result = proto.Bool(false)
		return
	}

	gameID := int(req.GetGameId())
	levelID := int(req.GetLevelId())
	// 选场判断游戏场次ID是否开启救济
	if reqType == client_alms.AlmsReqType_SELECTED {
		// 获取救济金配置
		gameLeveConfigMaps, err := configclient.GetAllGameLevelConfig()
		if err != nil {
			response.Result = proto.Bool(false)
			logrus.WithError(err).Debugln("获取救济金配置失败")
			return
		}
		flag := true
		for _, gameLeveConfigMap := range gameLeveConfigMaps {
			if gameLeveConfigMap.GameID == gameID && gameLeveConfigMap.LevelID == levelID {
				flag = false
				if gameLeveConfigMap.IsAlms == 0 {
					response.Result = proto.Bool(false)
					logrus.Debugln(fmt.Sprintf("该游戏场次未开启救济金  gameID(%d) - levelID(%d) - IsAlms(%d)", gameID, levelID, gameLeveConfigMap.IsAlms))
					return
				}
				// 如果玩家身上的金币已经够了，不应该发救济金
				if int64(gameLeveConfigMap.LowScores) <= playerGold {
					logrus.Debugln(fmt.Sprintf("玩家身上的金币，足够该场次的下限 playerGold(%d) -- gameLeveConfigMap(%v)", playerGold, gameLeveConfigMap))
					response.Result = proto.Bool(false)
					return
				}
				logrus.Debugln(fmt.Sprintf("该游戏场次配置  gameID(%d) - levelID(%d) - lowScores(%d)", gameID, levelID, gameLeveConfigMap.LowScores))
				break
			}
		}
		if flag {
			logrus.Debugln(fmt.Sprintf("该游戏场次不存在 gameID(%d) - levelID(%d)", gameID, levelID))
			response.Result = proto.Bool(false)
			flag = false
			return
		}
	}

	// 游戏中,玩家破产才领
	if reqType == client_alms.AlmsReqType_INGAME && playerGold != 0 {
		entry.Debugln(fmt.Sprintf("玩家金币不为0,没有破产 gold(%d)", playerGold))
		response.Result = proto.Bool(false)
		return
	}

	//验证通过，玩家领取次数加1
	if err := data.UpdatePlayerGotTimesByPlayerID(playerID, ac.PlayerGotTimes+1); err != nil {
		entry.WithError(err).Debugf(fmt.Sprintf(" playerID(%d) times(%d)- 更改玩家救济金领取次数失败", playerID, ac.PlayerGotTimes))
		response.Result = proto.Bool(false)
		return
	}
	// 更改玩家身上的金币 TODO almsFuncID 渠道ID
	almschannl := int64(0)
	changeGold, err := goldclient.AddGold(playerID, int16(gold.GoldType_GOLD_COIN), ac.GetNumber, int32(constant.ALMSFUNC), almschannl, 0, 0)
	if err != nil || changeGold != playerGold+ac.GetNumber {
		entry.WithError(err).Debugln(fmt.Sprintf("playerID(%d) - 设置玩家身上金币失败 changeGold(%d) , needGold(%d)", playerID, changeGold, (playerGold + ac.GetNumber)))
		response.Result = proto.Bool(false)
		return
	}
	response.PlayerAlmsTimes = proto.Int32(int32(ac.PlayerGotTimes + 1)) // 玩家已经领取的次数
	response.ChangeGold = proto.Int64(changeGold)
	entry.WithFields(logrus.Fields{"PlayerAlmsTimes": *response.PlayerAlmsTimes,
		"GetNumber": ac.GetNumber, "GetNorm": ac.GetNorm,
		"oldGold": playerGold, "newGold": changeGold}).Debugln("申请救济成功")
	return
}
