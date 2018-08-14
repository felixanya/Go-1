package almsserver

import (
	"steve/alms/data"
	client_alms "steve/client_pb/alms"
	"steve/client_pb/msgid"
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

	// 校验玩家背包是否有金豆 TODO，有返回false

	gameID := int32(*req.GetGameId().Enum()) // 游戏ID
	levelID := req.GetLevelId()              // 场次ID
	almsGetType := req.GetAlmsApplyType()    // 救济金领取类型
	totalGold := req.GetTotalGold()          // 总共需要的金币
	version := req.GetVersion()              // 版本号
	// 获取配置用于验证
	ac, err := data.GetAlmsConfigByPlayerID(playerID)
	if err != nil {
		entry.WithError(err).Errorf("根据玩家ID获取救济金配置失败 playerID(%v)", playerID)
		response.Result = proto.Bool(false)
		return
	}
	response.NewVersion = proto.Int32(int32(ac.Version))
	// 版本不对应,配置发生改变,发送新的配置信息
	if version != int32(ac.Version) {
		newAlmsConfig := &client_alms.AlmsConfig{
			AlmsGetNorm:      proto.Int64(ac.GetNorm),                 // 救济金
			AlmsGetTimes:     proto.Int32(int32(ac.GetTimes)),         //次数
			AlmsGetNumber:    proto.Int64(ac.GetNumber),               // 领取数量
			AlmsCountDonw:    proto.Int32(int32(ac.AlmsCountDonw)),    //救济倒计时
			DepositCountDonw: proto.Int32(int32(ac.DepositCountDonw)), //救济倒计时
		}
		response.NewAlmsConfig = newAlmsConfig
	}
	//是否还有领取次数
	if ac.PlayerGotTimes >= 3 {
		entry.Errorf("领取次数已满 times(%v) ", ac.PlayerGotTimes)
		response.Result = proto.Bool(false)
		return
	}
	//从金币服获取玩家身上金币，
	playerGold, err := goldclient.GetGold(playerID, int16(gold.GoldType_GOLD_COIN))
	if err != nil {
		entry.WithError(err).Errorf(" playerID(%v) - 从金币服获取玩家身上金币失败", playerID)
		response.Result = proto.Bool(false)
		return
	}
	if playerGold > ac.GetNorm {
		entry.Errorf("玩家身上的金币没有达到救济线  playerGold(%v) ", playerGold)
		response.Result = proto.Bool(false)
		return
	}
	//不是登陆类型，才玩家身上金，加上救济数量，是否符合，总需求
	isExist := map[client_alms.AlmsApplyType]bool{
		client_alms.AlmsApplyType_AAT_LOGIN:      true, // 登陆
		client_alms.AlmsApplyType_AAT_SELECTIONS: true, // 选场
		client_alms.AlmsApplyType_AAT_GAME_OVER:  true, // 游戏结束，下一局
		client_alms.AlmsApplyType_AAT_IN_GAME:    true, // 游戏中
	}[almsGetType]
	if !isExist {
		entry.Errorf("救济金获取类型不存在 (%v) ", almsGetType)
		response.Result = proto.Bool(false)
		return
	}
	// 选场和游戏结束下一局时,有下限，判断下限参数是否正确
	if almsGetType == client_alms.AlmsApplyType_AAT_SELECTIONS || almsGetType == client_alms.AlmsApplyType_AAT_GAME_OVER {
		gameLevelConfigs, err := configclient.GetGameLevelConfigMap()
		if err != nil || len(gameLevelConfigs) == 0 {
			entry.WithError(err).Debugln("游戏场次配置获取失败")
			response.Result = proto.Bool(false)
			return
		}
		for _, glc := range gameLevelConfigs {
			if int32(glc.GameID) == gameID && int32(glc.LevelID) == levelID {
				if int64(glc.LowScores) != totalGold {
					entry.Errorf("needGold(%d) no eq lowScores(%d)", totalGold, glc.LowScores)
					response.Result = proto.Bool(false)
					return
				}
				break
			}
		}
	}
	// 不是登录情况，有所需金币，如果救济金+身上金还是不足所需金币，应该是弹快冲
	if almsGetType != client_alms.AlmsApplyType_AAT_LOGIN && totalGold > (playerGold+ac.GetNumber) {
		entry.Errorf("所需金币不足 totalGold(%v) playerGold(%v) GetNumber(%v) ", totalGold, playerGold, ac.GetNumber)
		response.Result = proto.Bool(false)
		return
	}
	//验证通过，玩家领取次数加1
	if err := data.UpdatePlayerGotTimesByPlayerID(playerID, ac.PlayerGotTimes+1); err != nil {
		entry.WithError(err).Errorf(" playerID(%d) times(%d)- 更改玩家救济金领取次数失败", playerID, ac.PlayerGotTimes)
		response.Result = proto.Bool(false)
		return
	}
	// 更改玩家身上的金币 TODO almsFuncID 渠道ID
	almsFuncID := int32(10)
	almschannl := int64(10)
	changeGold, err := goldclient.AddGold(playerID, int16(gold.GoldType_GOLD_COIN), ac.GetNumber, almsFuncID, almschannl, gameID, levelID)
	if err != nil || changeGold != playerGold+ac.GetNumber {
		entry.WithError(err).Errorf(" playerID(%v) - 设置玩家身上金币失败 changeGold(%v) , needGold(%v)", playerID, changeGold, playerGold+ac.GetNumber)
		response.Result = proto.Bool(false)
		return
	}
	response.PlayerAlmsTimes = proto.Int32(int32(ac.PlayerGotTimes + 1))
	response.ChangeGold = proto.Int64(changeGold)
	entry.WithFields(logrus.Fields{"playerID": playerID, "oldGold": playerGold, "newGold": changeGold}).Infoln("申请救济成功")
	return
}
