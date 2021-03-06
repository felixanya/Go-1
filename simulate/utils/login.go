package utils

import (
	"fmt"
	"steve/client_pb/hall"
	"steve/client_pb/login"
	msgid "steve/client_pb/msgid"
	"steve/simulate/config"
	"steve/simulate/connect"
	"steve/simulate/facade"

	"steve/simulate/global"
	"steve/simulate/interfaces"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type clientPlayer struct {
	playerID  uint64
	coin      uint64
	client    interfaces.Client
	usrName   string
	accountID uint64
	expectors map[msgid.MsgID]interfaces.MessageExpector
	token     string
}

func (p *clientPlayer) GetID() uint64 {
	return p.playerID
}
func (p *clientPlayer) GetCoin() uint64 {
	return p.coin
}

func (p *clientPlayer) GetClient() interfaces.Client {
	return p.client
}

func (p *clientPlayer) GetUsrName() string {
	return p.usrName
}

func (p *clientPlayer) GetAccountID() uint64 {
	return p.accountID
}

func (p *clientPlayer) GetToken() string {
	return p.token
}

func (p *clientPlayer) AddExpectors(msgIDs ...msgid.MsgID) {
	for _, msgID := range msgIDs {
		p.expectors[msgID], _ = p.client.ExpectMessage(msgID)
	}
}

func (p *clientPlayer) GetExpector(msgID msgid.MsgID) interfaces.MessageExpector {
	return p.expectors[msgID]
}

func loginPlayer(request *login.LoginAuthReq, expectedMsgs ...msgid.MsgID) (interfaces.ClientPlayer, error) {
	gateClient := connect.NewTestClient(config.GetGatewayServerAddr(), config.GetClientVersion())
	if gateClient == nil {
		return nil, fmt.Errorf("连接网关服失败")
	}
	expectors := make(map[msgid.MsgID]interfaces.MessageExpector, len(expectedMsgs))
	for _, msgID := range expectedMsgs {
		expector, err := gateClient.ExpectMessage(msgID)
		if err != nil {
			return nil, fmt.Errorf("expector 失败, %v", msgID)
		}
		expectors[msgID] = expector
	}
	response := &login.LoginAuthRsp{}
	err := facade.Request(gateClient, msgid.MsgID_LOGIN_AUTH_REQ, request,
		global.DefaultWaitMessageTime, msgid.MsgID_LOGIN_AUTH_RSP, response)
	if err != nil {
		return nil, fmt.Errorf("请求登录失败：%v", err)
	}
	if response.GetErrCode() != login.ErrorCode_SUCCESS {
		return nil, fmt.Errorf("登录失败，错误码：%v", response.GetErrCode())
	}
	playerInfoRsp := &hall.HallGetPlayerInfoRsp{}
	if err := facade.Request(gateClient, msgid.MsgID_HALL_GET_PLAYER_INFO_REQ, &hall.HallGetPlayerInfoReq{},
		global.DefaultWaitMessageTime, msgid.MsgID_HALL_GET_PLAYER_INFO_RSP, playerInfoRsp); err != nil {
		return nil, fmt.Errorf("请求用户信息失败：%v", err)
	}
	if playerInfoRsp.GetErrCode() != 0 {
		return nil, fmt.Errorf("请求用户信息失败，错误码:%d", playerInfoRsp.GetErrCode())
	}

	logrus.Infoln("登录成功", response)
	return &clientPlayer{
		playerID:  response.GetPlayerId(),
		accountID: request.GetAccountId(),
		coin:      playerInfoRsp.GetCoin(),
		client:    gateClient,
		usrName:   "",
		expectors: expectors,
		token:     response.GetToken(),
	}, nil
}

// LoginNewPlayer 自动分配账号 ID， 生成账号名称，然后登录
func LoginNewPlayer(expectedMsgs ...msgid.MsgID) (interfaces.ClientPlayer, error) {
	return loginPlayer(&login.LoginAuthReq{
		AccountId: proto.Uint64(global.AllocAccountID()),
	}, expectedMsgs...)
}

// LoginPlayerByToken 使用 token 登录玩家
func LoginPlayerByToken(playerID uint64, token string) (interfaces.ClientPlayer, error) {
	request := &login.LoginAuthReq{
		PlayerId: proto.Uint64(playerID),
		Token:    proto.String(token),
	}
	return loginPlayer(request)
}

// LoginPlayerYouke 登录游客
func LoginPlayerYouke(imei string) (interfaces.ClientPlayer, error) {
	loginData := login.LoginData{
		Type:     login.LoginType_VISITOR_LOGIN.Enum(),
		Channel:  proto.Int32(1),
		Username: proto.String(imei),
		ProId:    proto.Uint64(55555),
		CityId:   proto.Uint64(1111),
		BindInfo: &login.BindInfo{},
	}
	return loginPlayer(&login.LoginAuthReq{
		RequestData: &loginData,
	})
}

// Channel          *ChannelType `protobuf:"varint,2,opt,name=channel,enum=login.ChannelType" json:"channel,omitempty"`
// Username         *string      `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
// DymcCode         *string      `protobuf:"bytes,4,opt,name=dymc_code,json=dymcCode" json:"dymc_code,omitempty"`
// Password         *string      `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
// ProId            *uint64      `protobuf:"varint,6,opt,name=pro_id,json=proId" json:"pro_id,omitempty"`
// CityId           *uint64      `protobuf:"varint,7,opt,name=city_id,json=cityId" json:"city_id,omitempty"`
// BindInfo         *BindInfo    `protobuf:"bytes,8,opt,name=bind_info,json=bindInfo" json:"bind_info,omitempty"`

func UpdatePlayerClientInfo(client interfaces.Client, player interfaces.ClientPlayer, deskData *DeskData) {
	oldPlayer, exist := deskData.Players[player.GetID()]
	if !exist {
		return
	}

	newPlayer := DeskPlayer{
		Player:    player,
		Seat:      oldPlayer.Seat,
		Expectors: createPlayerExpectors(player.GetClient()),
	}
	deskData.Players[player.GetID()] = newPlayer
	return
}

func UpdateDDZPlayerClientInfo(client interfaces.Client, player interfaces.ClientPlayer, deskData *DeskData) {
	oldPlayer, exist := deskData.Players[player.GetID()]
	if !exist {
		return
	}

	newPlayer := DeskPlayer{
		Player:    player,
		Seat:      oldPlayer.Seat,
		Expectors: createDDZPlayerExpectors(player.GetClient()),
	}
	deskData.Players[player.GetID()] = newPlayer
	return
}

// GenerateAccountName 生成账号名字
func GenerateAccountName(accountID uint64) string {
	return fmt.Sprintf("account_%v", accountID)
}

// RequestAuth 请求认证
func RequestAuth(client interfaces.Client, accountID uint64, accountName string, expireDuration time.Duration) (*login.LoginAuthRsp, error) {
	request := &login.LoginAuthReq{
		AccountId: proto.Uint64(accountID),
	}
	response := &login.LoginAuthRsp{}
	err := facade.Request(client, msgid.MsgID_LOGIN_AUTH_REQ, request, global.DefaultWaitMessageTime, msgid.MsgID_LOGIN_AUTH_RSP, response)
	return response, err
}
