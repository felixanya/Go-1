// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hall.proto

package hall

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common1 "steve/client_pb/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// HallGetPlayerInfoReq 获取玩家信息请求
type HallGetPlayerInfoReq struct {
	Reserve          *int32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HallGetPlayerInfoReq) Reset()                    { *m = HallGetPlayerInfoReq{} }
func (m *HallGetPlayerInfoReq) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerInfoReq) ProtoMessage()               {}
func (*HallGetPlayerInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *HallGetPlayerInfoReq) GetReserve() int32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// HallGetPlayerInfoRsp 获取玩家信息应答
type HallGetPlayerInfoRsp struct {
	ErrCode          *uint32              `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	ErrDesc          *string              `protobuf:"bytes,2,opt,name=err_desc,json=errDesc" json:"err_desc,omitempty"`
	NickName         *string              `protobuf:"bytes,3,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Coin             *uint64              `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	PlayerState      *common1.PlayerState `protobuf:"varint,5,opt,name=player_state,json=playerState,enum=common.PlayerState" json:"player_state,omitempty"`
	GameId           *common1.GameId      `protobuf:"varint,6,opt,name=game_id,json=gameId,enum=common.GameId" json:"game_id,omitempty"`
	RealnameStatus   *uint32              `protobuf:"varint,7,opt,name=realname_status,json=realnameStatus" json:"realname_status,omitempty"`
	ShowUid          *uint64              `protobuf:"varint,8,opt,name=show_uid,json=showUid" json:"show_uid,omitempty"`
	Gender           *uint32              `protobuf:"varint,9,opt,name=gender" json:"gender,omitempty"`
	Avator           *string              `protobuf:"bytes,10,opt,name=avator" json:"avator,omitempty"`
	RealnameReward   *uint64              `protobuf:"varint,11,opt,name=realname_reward,json=realnameReward" json:"realname_reward,omitempty"`
	Phone            *string              `protobuf:"bytes,12,opt,name=phone" json:"phone,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *HallGetPlayerInfoRsp) Reset()                    { *m = HallGetPlayerInfoRsp{} }
func (m *HallGetPlayerInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerInfoRsp) ProtoMessage()               {}
func (*HallGetPlayerInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *HallGetPlayerInfoRsp) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *HallGetPlayerInfoRsp) GetErrDesc() string {
	if m != nil && m.ErrDesc != nil {
		return *m.ErrDesc
	}
	return ""
}

func (m *HallGetPlayerInfoRsp) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *HallGetPlayerInfoRsp) GetCoin() uint64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *HallGetPlayerInfoRsp) GetPlayerState() common1.PlayerState {
	if m != nil && m.PlayerState != nil {
		return *m.PlayerState
	}
	return common1.PlayerState_PS_IDLE
}

func (m *HallGetPlayerInfoRsp) GetGameId() common1.GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return common1.GameId_GAMEID_XUELIU
}

func (m *HallGetPlayerInfoRsp) GetRealnameStatus() uint32 {
	if m != nil && m.RealnameStatus != nil {
		return *m.RealnameStatus
	}
	return 0
}

func (m *HallGetPlayerInfoRsp) GetShowUid() uint64 {
	if m != nil && m.ShowUid != nil {
		return *m.ShowUid
	}
	return 0
}

func (m *HallGetPlayerInfoRsp) GetGender() uint32 {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return 0
}

func (m *HallGetPlayerInfoRsp) GetAvator() string {
	if m != nil && m.Avator != nil {
		return *m.Avator
	}
	return ""
}

func (m *HallGetPlayerInfoRsp) GetRealnameReward() uint64 {
	if m != nil && m.RealnameReward != nil {
		return *m.RealnameReward
	}
	return 0
}

func (m *HallGetPlayerInfoRsp) GetPhone() string {
	if m != nil && m.Phone != nil {
		return *m.Phone
	}
	return ""
}

// HallUpdatePlayerInfoReq 更新玩家信息请求
type HallUpdatePlayerInfoReq struct {
	Uid              *uint64               `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	NickName         *string               `protobuf:"bytes,2,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Avator           *string               `protobuf:"bytes,3,opt,name=avator" json:"avator,omitempty"`
	Gender           *common1.PlayerGender `protobuf:"varint,4,opt,name=gender,enum=common.PlayerGender" json:"gender,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *HallUpdatePlayerInfoReq) Reset()                    { *m = HallUpdatePlayerInfoReq{} }
func (m *HallUpdatePlayerInfoReq) String() string            { return proto.CompactTextString(m) }
func (*HallUpdatePlayerInfoReq) ProtoMessage()               {}
func (*HallUpdatePlayerInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *HallUpdatePlayerInfoReq) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *HallUpdatePlayerInfoReq) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *HallUpdatePlayerInfoReq) GetAvator() string {
	if m != nil && m.Avator != nil {
		return *m.Avator
	}
	return ""
}

func (m *HallUpdatePlayerInfoReq) GetGender() common1.PlayerGender {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return common1.PlayerGender_PG_NIL
}

// HallUpdatePlayerInfoRsp 更新玩家信应答
type HallUpdatePlayerInfoRsp struct {
	ErrCode          *uint32               `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	NickName         *string               `protobuf:"bytes,2,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Avator           *string               `protobuf:"bytes,3,opt,name=avator" json:"avator,omitempty"`
	Gender           *common1.PlayerGender `protobuf:"varint,4,opt,name=gender,enum=common.PlayerGender" json:"gender,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *HallUpdatePlayerInfoRsp) Reset()                    { *m = HallUpdatePlayerInfoRsp{} }
func (m *HallUpdatePlayerInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*HallUpdatePlayerInfoRsp) ProtoMessage()               {}
func (*HallUpdatePlayerInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *HallUpdatePlayerInfoRsp) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *HallUpdatePlayerInfoRsp) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *HallUpdatePlayerInfoRsp) GetAvator() string {
	if m != nil && m.Avator != nil {
		return *m.Avator
	}
	return ""
}

func (m *HallUpdatePlayerInfoRsp) GetGender() common1.PlayerGender {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return common1.PlayerGender_PG_NIL
}

// HallGetPlayerStateReq 获取玩家当前状态请求
type HallGetPlayerStateReq struct {
	Reserve          *int32  `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	UserData         *uint64 `protobuf:"varint,2,opt,name=user_data,json=userData" json:"user_data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallGetPlayerStateReq) Reset()                    { *m = HallGetPlayerStateReq{} }
func (m *HallGetPlayerStateReq) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerStateReq) ProtoMessage()               {}
func (*HallGetPlayerStateReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *HallGetPlayerStateReq) GetReserve() int32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

func (m *HallGetPlayerStateReq) GetUserData() uint64 {
	if m != nil && m.UserData != nil {
		return *m.UserData
	}
	return 0
}

// HallGetPlayerStateRsp 获取玩家当前状态响应
type HallGetPlayerStateRsp struct {
	PlayerState      *common1.PlayerState `protobuf:"varint,1,opt,name=player_state,json=playerState,enum=common.PlayerState" json:"player_state,omitempty"`
	GameId           *common1.GameId      `protobuf:"varint,2,opt,name=game_id,json=gameId,enum=common.GameId" json:"game_id,omitempty"`
	UserData         *uint64              `protobuf:"varint,3,opt,name=user_data,json=userData" json:"user_data,omitempty"`
	LevelId          *uint32              `protobuf:"varint,4,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *HallGetPlayerStateRsp) Reset()                    { *m = HallGetPlayerStateRsp{} }
func (m *HallGetPlayerStateRsp) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerStateRsp) ProtoMessage()               {}
func (*HallGetPlayerStateRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *HallGetPlayerStateRsp) GetPlayerState() common1.PlayerState {
	if m != nil && m.PlayerState != nil {
		return *m.PlayerState
	}
	return common1.PlayerState_PS_IDLE
}

func (m *HallGetPlayerStateRsp) GetGameId() common1.GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return common1.GameId_GAMEID_XUELIU
}

func (m *HallGetPlayerStateRsp) GetUserData() uint64 {
	if m != nil && m.UserData != nil {
		return *m.UserData
	}
	return 0
}

func (m *HallGetPlayerStateRsp) GetLevelId() uint32 {
	if m != nil && m.LevelId != nil {
		return *m.LevelId
	}
	return 0
}

// HallGetGameListInfoReq 获取游戏列表信息请求
type HallGetGameListInfoReq struct {
	Reserve          *int32 `protobuf:"varint,1,opt,name=reserve" json:"reserve,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HallGetGameListInfoReq) Reset()                    { *m = HallGetGameListInfoReq{} }
func (m *HallGetGameListInfoReq) String() string            { return proto.CompactTextString(m) }
func (*HallGetGameListInfoReq) ProtoMessage()               {}
func (*HallGetGameListInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *HallGetGameListInfoReq) GetReserve() int32 {
	if m != nil && m.Reserve != nil {
		return *m.Reserve
	}
	return 0
}

// HallGetGameListInfoRsp 获取游戏列表信息响应
type HallGetGameListInfoRsp struct {
	ErrCode          *uint32                    `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	GameConfig       []*common1.GameConfig      `protobuf:"bytes,2,rep,name=game_config,json=gameConfig" json:"game_config,omitempty"`
	GameLevelConfig  []*common1.GameLevelConfig `protobuf:"bytes,3,rep,name=game_level_config,json=gameLevelConfig" json:"game_level_config,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *HallGetGameListInfoRsp) Reset()                    { *m = HallGetGameListInfoRsp{} }
func (m *HallGetGameListInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*HallGetGameListInfoRsp) ProtoMessage()               {}
func (*HallGetGameListInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *HallGetGameListInfoRsp) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *HallGetGameListInfoRsp) GetGameConfig() []*common1.GameConfig {
	if m != nil {
		return m.GameConfig
	}
	return nil
}

func (m *HallGetGameListInfoRsp) GetGameLevelConfig() []*common1.GameLevelConfig {
	if m != nil {
		return m.GameLevelConfig
	}
	return nil
}

// HallRealNameReq 实名认证请求
type HallRealNameReq struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	IdCard           *string `protobuf:"bytes,2,opt,name=id_card,json=idCard" json:"id_card,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallRealNameReq) Reset()                    { *m = HallRealNameReq{} }
func (m *HallRealNameReq) String() string            { return proto.CompactTextString(m) }
func (*HallRealNameReq) ProtoMessage()               {}
func (*HallRealNameReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *HallRealNameReq) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *HallRealNameReq) GetIdCard() string {
	if m != nil && m.IdCard != nil {
		return *m.IdCard
	}
	return ""
}

// HallRealNameRsp 实名认证响应
type HallRealNameRsp struct {
	ErrCode          *uint32 `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	ErrDesc          *string `protobuf:"bytes,2,opt,name=err_desc,json=errDesc" json:"err_desc,omitempty"`
	CoinReward       *uint64 `protobuf:"varint,3,opt,name=coin_reward,json=coinReward" json:"coin_reward,omitempty"`
	NewCoin          *uint64 `protobuf:"varint,4,opt,name=new_coin,json=newCoin" json:"new_coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallRealNameRsp) Reset()                    { *m = HallRealNameRsp{} }
func (m *HallRealNameRsp) String() string            { return proto.CompactTextString(m) }
func (*HallRealNameRsp) ProtoMessage()               {}
func (*HallRealNameRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *HallRealNameRsp) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *HallRealNameRsp) GetErrDesc() string {
	if m != nil && m.ErrDesc != nil {
		return *m.ErrDesc
	}
	return ""
}

func (m *HallRealNameRsp) GetCoinReward() uint64 {
	if m != nil && m.CoinReward != nil {
		return *m.CoinReward
	}
	return 0
}

func (m *HallRealNameRsp) GetNewCoin() uint64 {
	if m != nil && m.NewCoin != nil {
		return *m.NewCoin
	}
	return 0
}

// HallGetPlayerGameInfoReq 获取玩家对应游戏的信息请求
type HallGetPlayerGameInfoReq struct {
	Uid              *uint64         `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	GameId           *common1.GameId `protobuf:"varint,2,opt,name=game_id,json=gameId,enum=common.GameId" json:"game_id,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *HallGetPlayerGameInfoReq) Reset()                    { *m = HallGetPlayerGameInfoReq{} }
func (m *HallGetPlayerGameInfoReq) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerGameInfoReq) ProtoMessage()               {}
func (*HallGetPlayerGameInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *HallGetPlayerGameInfoReq) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *HallGetPlayerGameInfoReq) GetGameId() common1.GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return common1.GameId_GAMEID_XUELIU
}

// HallGetPlayerGameInfoRsp 获取玩家对应游戏的信息响应
type HallGetPlayerGameInfoRsp struct {
	ErrCode          *uint32             `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	Uid              *uint64             `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	GameId           *common1.GameId     `protobuf:"varint,3,opt,name=game_id,json=gameId,enum=common.GameId" json:"game_id,omitempty"`
	TotalBureau      *uint32             `protobuf:"varint,4,opt,name=total_bureau,json=totalBureau" json:"total_bureau,omitempty"`
	WinningRate      *float32            `protobuf:"fixed32,5,opt,name=winning_rate,json=winningRate" json:"winning_rate,omitempty"`
	MaxWinningStream *uint32             `protobuf:"varint,6,opt,name=max_winning_stream,json=maxWinningStream" json:"max_winning_stream,omitempty"`
	MaxMultiple      *uint32             `protobuf:"varint,7,opt,name=max_multiple,json=maxMultiple" json:"max_multiple,omitempty"`
	UserProperty     []*common1.Property `protobuf:"bytes,8,rep,name=user_property,json=userProperty" json:"user_property,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *HallGetPlayerGameInfoRsp) Reset()                    { *m = HallGetPlayerGameInfoRsp{} }
func (m *HallGetPlayerGameInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*HallGetPlayerGameInfoRsp) ProtoMessage()               {}
func (*HallGetPlayerGameInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *HallGetPlayerGameInfoRsp) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *HallGetPlayerGameInfoRsp) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *HallGetPlayerGameInfoRsp) GetGameId() common1.GameId {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return common1.GameId_GAMEID_XUELIU
}

func (m *HallGetPlayerGameInfoRsp) GetTotalBureau() uint32 {
	if m != nil && m.TotalBureau != nil {
		return *m.TotalBureau
	}
	return 0
}

func (m *HallGetPlayerGameInfoRsp) GetWinningRate() float32 {
	if m != nil && m.WinningRate != nil {
		return *m.WinningRate
	}
	return 0
}

func (m *HallGetPlayerGameInfoRsp) GetMaxWinningStream() uint32 {
	if m != nil && m.MaxWinningStream != nil {
		return *m.MaxWinningStream
	}
	return 0
}

func (m *HallGetPlayerGameInfoRsp) GetMaxMultiple() uint32 {
	if m != nil && m.MaxMultiple != nil {
		return *m.MaxMultiple
	}
	return 0
}

func (m *HallGetPlayerGameInfoRsp) GetUserProperty() []*common1.Property {
	if m != nil {
		return m.UserProperty
	}
	return nil
}

// MoneyChangeNtf 货币数变化通知
type MoneyChangeNtf struct {
	PlayerId         *uint64        `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Money            *common1.Money `protobuf:"bytes,2,opt,name=money" json:"money,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *MoneyChangeNtf) Reset()                    { *m = MoneyChangeNtf{} }
func (m *MoneyChangeNtf) String() string            { return proto.CompactTextString(m) }
func (*MoneyChangeNtf) ProtoMessage()               {}
func (*MoneyChangeNtf) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *MoneyChangeNtf) GetPlayerId() uint64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *MoneyChangeNtf) GetMoney() *common1.Money {
	if m != nil {
		return m.Money
	}
	return nil
}

func init() {
	proto.RegisterType((*HallGetPlayerInfoReq)(nil), "hall.HallGetPlayerInfoReq")
	proto.RegisterType((*HallGetPlayerInfoRsp)(nil), "hall.HallGetPlayerInfoRsp")
	proto.RegisterType((*HallUpdatePlayerInfoReq)(nil), "hall.HallUpdatePlayerInfoReq")
	proto.RegisterType((*HallUpdatePlayerInfoRsp)(nil), "hall.HallUpdatePlayerInfoRsp")
	proto.RegisterType((*HallGetPlayerStateReq)(nil), "hall.HallGetPlayerStateReq")
	proto.RegisterType((*HallGetPlayerStateRsp)(nil), "hall.HallGetPlayerStateRsp")
	proto.RegisterType((*HallGetGameListInfoReq)(nil), "hall.HallGetGameListInfoReq")
	proto.RegisterType((*HallGetGameListInfoRsp)(nil), "hall.HallGetGameListInfoRsp")
	proto.RegisterType((*HallRealNameReq)(nil), "hall.HallRealNameReq")
	proto.RegisterType((*HallRealNameRsp)(nil), "hall.HallRealNameRsp")
	proto.RegisterType((*HallGetPlayerGameInfoReq)(nil), "hall.HallGetPlayerGameInfoReq")
	proto.RegisterType((*HallGetPlayerGameInfoRsp)(nil), "hall.HallGetPlayerGameInfoRsp")
	proto.RegisterType((*MoneyChangeNtf)(nil), "hall.MoneyChangeNtf")
}

func init() { proto.RegisterFile("hall.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 794 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0xe3, 0x36,
	0x10, 0x86, 0x2c, 0x27, 0x76, 0x28, 0xdb, 0x49, 0x59, 0x37, 0xe1, 0xb6, 0x87, 0xba, 0xea, 0x61,
	0x7d, 0x58, 0xec, 0x16, 0x2e, 0xda, 0x63, 0x0f, 0xeb, 0x05, 0x52, 0x03, 0xbb, 0x46, 0xc1, 0x20,
	0x28, 0xd0, 0x8b, 0xc0, 0x15, 0x27, 0xb2, 0x50, 0x89, 0xd4, 0x52, 0xb4, 0x9d, 0x9c, 0x7b, 0x2f,
	0xd0, 0x53, 0x9f, 0xa2, 0x40, 0x9f, 0xa0, 0xcf, 0x56, 0x70, 0x24, 0x79, 0xa3, 0xd4, 0x71, 0x7f,
	0x0e, 0x3d, 0x59, 0xf3, 0x0d, 0x67, 0xf8, 0xcd, 0xf0, 0x9b, 0x31, 0x21, 0x2b, 0x91, 0x65, 0xcf,
	0x0b, 0xa3, 0xad, 0xa6, 0x5d, 0xf7, 0xfd, 0xf1, 0x20, 0xd6, 0x79, 0xae, 0x55, 0x85, 0x85, 0x5f,
	0x90, 0xf1, 0xb7, 0x22, 0xcb, 0x2e, 0xc1, 0x7e, 0x97, 0x89, 0x3b, 0x30, 0x0b, 0x75, 0xa3, 0x39,
	0xbc, 0xa3, 0x8c, 0xf4, 0x0c, 0x94, 0x60, 0x36, 0xc0, 0xbc, 0x89, 0x37, 0x3d, 0xe2, 0x8d, 0x19,
	0xfe, 0xe2, 0xef, 0x0b, 0x29, 0x0b, 0xfa, 0x84, 0xf4, 0xc1, 0x98, 0x28, 0xd6, 0xb2, 0x8a, 0x19,
	0xf2, 0x1e, 0x18, 0x33, 0xd7, 0x12, 0x1a, 0x97, 0x84, 0x32, 0x66, 0x9d, 0x89, 0x37, 0x3d, 0x41,
	0xd7, 0x2b, 0x28, 0x63, 0xfa, 0x09, 0x39, 0x51, 0x69, 0xfc, 0x63, 0xa4, 0x44, 0x0e, 0xcc, 0x47,
	0x5f, 0xdf, 0x01, 0x4b, 0x91, 0x03, 0xa5, 0xa4, 0x1b, 0xeb, 0x54, 0xb1, 0xee, 0xc4, 0x9b, 0x76,
	0x39, 0x7e, 0xd3, 0xaf, 0xc9, 0xa0, 0xc0, 0x7b, 0xa3, 0xd2, 0x0a, 0x0b, 0xec, 0x68, 0xe2, 0x4d,
	0x47, 0xb3, 0x0f, 0x9f, 0xd7, 0x65, 0x55, 0x9c, 0xae, 0x9c, 0x8b, 0x07, 0xc5, 0x7b, 0x83, 0x3e,
	0x25, 0xbd, 0x44, 0xe4, 0x10, 0xa5, 0x92, 0x1d, 0x63, 0xc8, 0xa8, 0x09, 0xb9, 0x14, 0x39, 0x2c,
	0x24, 0x3f, 0x4e, 0xf0, 0x97, 0x3e, 0x25, 0xa7, 0x06, 0x44, 0xe6, 0x08, 0xe1, 0x15, 0xeb, 0x92,
	0xf5, 0xb0, 0x9c, 0x51, 0x03, 0x5f, 0x21, 0xea, 0xaa, 0x2a, 0x57, 0x7a, 0x1b, 0xad, 0x53, 0xc9,
	0xfa, 0xc8, 0xb0, 0xe7, 0xec, 0xeb, 0x54, 0xd2, 0x73, 0x72, 0x9c, 0x80, 0x92, 0x60, 0xd8, 0x09,
	0x86, 0xd6, 0x96, 0xc3, 0xc5, 0x46, 0x58, 0x6d, 0x18, 0xc1, 0x52, 0x6b, 0xab, 0x75, 0xa7, 0x81,
	0xad, 0x30, 0x92, 0x05, 0x98, 0x71, 0x77, 0x27, 0x47, 0x94, 0x8e, 0xc9, 0x51, 0xb1, 0xd2, 0x0a,
	0xd8, 0x00, 0xe3, 0x2b, 0x23, 0xfc, 0xd9, 0x23, 0x17, 0xee, 0x4d, 0xae, 0x0b, 0x29, 0x2c, 0xb4,
	0x5f, 0xf2, 0x8c, 0xf8, 0x8e, 0xa0, 0x87, 0xe9, 0xdc, 0x67, 0xbb, 0xe5, 0x9d, 0x07, 0x2d, 0x7f,
	0xcf, 0xd0, 0x6f, 0x31, 0x7c, 0xb6, 0xab, 0xa8, 0x8b, 0xdd, 0x1b, 0xb7, 0x1b, 0x7e, 0x89, 0xbe,
	0xa6, 0xce, 0xf0, 0xd7, 0xc7, 0x08, 0x1d, 0xd6, 0xc9, 0xff, 0xc0, 0x6c, 0x49, 0x3e, 0x6a, 0xa9,
	0xb7, 0x52, 0xca, 0x21, 0xc5, 0x3b, 0x56, 0xeb, 0x12, 0x4c, 0x24, 0x85, 0x15, 0xc8, 0xaa, 0xcb,
	0xfb, 0x0e, 0x78, 0x25, 0xac, 0x08, 0x7f, 0xf7, 0xf6, 0x26, 0x2c, 0x8b, 0xbf, 0x08, 0xd5, 0xfb,
	0xf7, 0x42, 0xed, 0x1c, 0x14, 0x6a, 0x8b, 0x97, 0xdf, 0xe6, 0xe5, 0xba, 0x9c, 0xc1, 0x06, 0x32,
	0x97, 0xa6, 0x5b, 0x75, 0x19, 0xed, 0x85, 0x0c, 0x67, 0xe4, 0xbc, 0x66, 0xec, 0x12, 0xbe, 0x4e,
	0x4b, 0xfb, 0xf7, 0x53, 0xff, 0x9b, 0xb7, 0x3f, 0xe8, 0xf0, 0x7b, 0x7e, 0x49, 0x02, 0x2c, 0x25,
	0xd6, 0xea, 0x26, 0x4d, 0x58, 0x67, 0xe2, 0x4f, 0x83, 0x19, 0xbd, 0x5f, 0xce, 0x1c, 0x3d, 0x9c,
	0x24, 0xbb, 0x6f, 0x3a, 0x27, 0x1f, 0x60, 0x50, 0x45, 0xbf, 0x0e, 0xf5, 0x31, 0xf4, 0xe2, 0x7e,
	0xe8, 0x6b, 0xe7, 0xaf, 0xe3, 0x4f, 0x93, 0x36, 0x10, 0x7e, 0x43, 0x4e, 0x1d, 0x5d, 0x0e, 0x22,
	0x5b, 0xe2, 0xf4, 0xbc, 0x73, 0xcb, 0x04, 0x75, 0xe5, 0xa1, 0x7a, 0xf0, 0x9b, 0x5e, 0x90, 0x5e,
	0x2a, 0xa3, 0xd8, 0xcd, 0x5b, 0x25, 0xb7, 0xe3, 0x54, 0xce, 0x85, 0x91, 0xe1, 0x4f, 0xde, 0x83,
	0x04, 0xff, 0x79, 0xc1, 0x7d, 0x4a, 0x02, 0xb7, 0xb7, 0x9a, 0xb1, 0xae, 0xde, 0x89, 0x38, 0xa8,
	0x1e, 0xe9, 0x27, 0xa4, 0xaf, 0x60, 0x1b, 0xdd, 0x5b, 0x74, 0x3d, 0x05, 0xdb, 0xb9, 0x4e, 0x55,
	0x78, 0x4d, 0x58, 0x4b, 0x5b, 0x28, 0x80, 0x47, 0xe7, 0xfa, 0x9f, 0x0a, 0x27, 0xfc, 0xa3, 0xf3,
	0x58, 0xde, 0xc3, 0x55, 0xd6, 0x57, 0x76, 0xf6, 0x5e, 0xe9, 0x1f, 0xd4, 0xea, 0x67, 0x64, 0x60,
	0xb5, 0x15, 0x59, 0xf4, 0x76, 0x6d, 0x40, 0xac, 0x6b, 0x49, 0x06, 0x88, 0xbd, 0x44, 0xc8, 0x1d,
	0xd9, 0xa6, 0x4a, 0xa5, 0x2a, 0x89, 0x4c, 0xb3, 0xd8, 0x3b, 0x3c, 0xa8, 0x31, 0xee, 0x46, 0xe3,
	0x19, 0xa1, 0xb9, 0xb8, 0x8d, 0x9a, 0x63, 0xa5, 0x35, 0x20, 0x72, 0x5c, 0xe7, 0x43, 0x7e, 0x96,
	0x8b, 0xdb, 0xef, 0x2b, 0xc7, 0x15, 0xe2, 0x2e, 0xa1, 0x3b, 0x9d, 0xaf, 0x33, 0x9b, 0x16, 0x19,
	0xd4, 0x5b, 0x3c, 0xc8, 0xc5, 0xed, 0x9b, 0x1a, 0xa2, 0x5f, 0x91, 0x21, 0x8e, 0x50, 0x61, 0x74,
	0x01, 0xc6, 0xde, 0xb1, 0x3e, 0xea, 0xec, 0x6c, 0x37, 0xa4, 0x35, 0xce, 0x07, 0xee, 0x58, 0x63,
	0x85, 0x9c, 0x8c, 0xde, 0x68, 0x05, 0x77, 0xf3, 0x95, 0x50, 0x09, 0x2c, 0xed, 0x8d, 0x9b, 0xc5,
	0x7a, 0xd8, 0x77, 0x6f, 0xd2, 0xaf, 0x80, 0x85, 0xa4, 0x9f, 0x93, 0xa3, 0xdc, 0x1d, 0xc7, 0xce,
	0x05, 0xb3, 0x61, 0x93, 0x1d, 0x73, 0xf0, 0xca, 0xf7, 0xf2, 0xfc, 0x87, 0x71, 0x69, 0x61, 0x03,
	0x2f, 0xe2, 0x2c, 0x05, 0x65, 0xa3, 0xe2, 0xed, 0x0b, 0xf7, 0x7f, 0xfd, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x13, 0x08, 0xeb, 0x85, 0xc2, 0x07, 0x00, 0x00,
}
