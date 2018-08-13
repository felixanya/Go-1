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
	GameId           *common1.GameId      `protobuf:"varint,6,opt,name=game_id,json=gameId,enum=common.GameId" json:"game_id,omitempty"`
	UserData         *uint64              `protobuf:"varint,2,opt,name=user_data,json=userData" json:"user_data,omitempty"`
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
}

func init() { proto.RegisterFile("hall.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xd2, 0xae, 0xed, 0x9c, 0x76, 0x1b, 0x66, 0x6c, 0x06, 0x0e, 0x94, 0x5c, 0xd6, 0xc3,
	0xb4, 0xa1, 0x21, 0x38, 0x72, 0x58, 0x27, 0x8d, 0x49, 0x63, 0x42, 0x9e, 0x26, 0x24, 0x2e, 0x91,
	0x97, 0xbc, 0x75, 0x11, 0x89, 0x9d, 0x39, 0x6e, 0xbb, 0x9d, 0xb9, 0x73, 0x43, 0xfc, 0x0a, 0xfe,
	0x02, 0x7f, 0x8c, 0x0b, 0xf2, 0x6b, 0xd2, 0x35, 0xa3, 0x2b, 0x12, 0x42, 0x9c, 0x6a, 0x7f, 0x2f,
	0x9f, 0xdf, 0xf7, 0xec, 0xef, 0xbd, 0x12, 0x72, 0x29, 0x92, 0x64, 0x27, 0xd3, 0xca, 0x28, 0x5a,
	0xb7, 0xeb, 0x27, 0xed, 0x50, 0xa5, 0xa9, 0x92, 0x13, 0xcc, 0x7f, 0x41, 0xd6, 0xdf, 0x8a, 0x24,
	0x39, 0x04, 0xf3, 0x3e, 0x11, 0x37, 0xa0, 0x8f, 0xe4, 0x85, 0xe2, 0x70, 0x45, 0x19, 0x69, 0x6a,
	0xc8, 0x41, 0x8f, 0x80, 0x39, 0x5d, 0xa7, 0xb7, 0xc4, 0xcb, 0xad, 0xff, 0xd3, 0x9d, 0x47, 0xc9,
	0x33, 0xfa, 0x98, 0xb4, 0x40, 0xeb, 0x20, 0x54, 0xd1, 0x84, 0xd3, 0xe1, 0x4d, 0xd0, 0xba, 0xaf,
	0x22, 0x28, 0x43, 0x11, 0xe4, 0x21, 0x73, 0xbb, 0x4e, 0x6f, 0x19, 0x43, 0x07, 0x90, 0x87, 0xf4,
	0x29, 0x59, 0x96, 0x71, 0xf8, 0x29, 0x90, 0x22, 0x05, 0x56, 0xc3, 0x58, 0xcb, 0x02, 0x27, 0x22,
	0x05, 0x4a, 0x49, 0x3d, 0x54, 0xb1, 0x64, 0xf5, 0xae, 0xd3, 0xab, 0x73, 0x5c, 0xd3, 0xd7, 0xa4,
	0x9d, 0x61, 0xde, 0x20, 0x37, 0xc2, 0x00, 0x5b, 0xea, 0x3a, 0xbd, 0x95, 0xbd, 0x87, 0x3b, 0x45,
	0x59, 0x13, 0x4d, 0xa7, 0x36, 0xc4, 0xbd, 0xec, 0x76, 0x43, 0xb7, 0x48, 0x73, 0x20, 0x52, 0x08,
	0xe2, 0x88, 0x35, 0x90, 0xb2, 0x52, 0x52, 0x0e, 0x45, 0x0a, 0x47, 0x11, 0x6f, 0x0c, 0xf0, 0x97,
	0x6e, 0x91, 0x55, 0x0d, 0x22, 0xb1, 0x82, 0x30, 0xc5, 0x30, 0x67, 0x4d, 0x2c, 0x67, 0xa5, 0x84,
	0x4f, 0x11, 0xb5, 0x55, 0xe5, 0x97, 0x6a, 0x1c, 0x0c, 0xe3, 0x88, 0xb5, 0x50, 0x61, 0xd3, 0xee,
	0xcf, 0xe2, 0x88, 0x6e, 0x90, 0xc6, 0x00, 0x64, 0x04, 0x9a, 0x2d, 0x23, 0xb5, 0xd8, 0x59, 0x5c,
	0x8c, 0x84, 0x51, 0x9a, 0x11, 0x2c, 0xb5, 0xd8, 0x55, 0x72, 0x6a, 0x18, 0x0b, 0x1d, 0x31, 0x0f,
	0x4f, 0x9c, 0xe6, 0xe4, 0x88, 0xfa, 0x5f, 0x1c, 0xb2, 0x69, 0x6f, 0xff, 0x2c, 0x8b, 0x84, 0x81,
	0xea, 0x9b, 0xad, 0x91, 0x9a, 0x95, 0xe2, 0x20, 0xd1, 0x2e, 0xab, 0x97, 0xeb, 0xde, 0xb9, 0xdc,
	0x5b, 0x2d, 0xb5, 0x8a, 0x96, 0xed, 0xa9, 0xf6, 0x3a, 0xde, 0xd3, 0x7a, 0xf5, 0x6a, 0x0f, 0x31,
	0x56, 0x56, 0xe4, 0x7f, 0xbb, 0x4f, 0xd0, 0x62, 0x47, 0xfc, 0x07, 0x65, 0x27, 0xe4, 0x51, 0xc5,
	0xa7, 0x13, 0x4f, 0x2c, 0xf2, 0xb6, 0x55, 0x35, 0xcc, 0x41, 0x07, 0x91, 0x30, 0x02, 0x55, 0xd5,
	0x79, 0xcb, 0x02, 0x07, 0xc2, 0x08, 0xff, 0xab, 0x33, 0xf7, 0xc0, 0x3c, 0xfb, 0xcd, 0x92, 0xce,
	0xbf, 0xb6, 0xe4, 0x42, 0x5d, 0x7b, 0x64, 0xa3, 0x90, 0x65, 0x59, 0xc7, 0x71, 0x6e, 0xfe, 0xdc,
	0xc4, 0xdf, 0x9d, 0xf9, 0xa4, 0xc5, 0x8f, 0xf6, 0x92, 0x78, 0xa8, 0x37, 0x54, 0xf2, 0x22, 0x1e,
	0x30, 0xb7, 0x5b, 0xeb, 0x79, 0x7b, 0x74, 0x56, 0x73, 0x1f, 0x23, 0x9c, 0x0c, 0xa6, 0x6b, 0xda,
	0x27, 0x0f, 0x90, 0x94, 0xc0, 0x08, 0x92, 0x92, 0x5a, 0x43, 0xea, 0xe6, 0x2c, 0xf5, 0xd8, 0xc6,
	0x0b, 0xfe, 0xea, 0xa0, 0x0a, 0xf8, 0x6f, 0xc8, 0xaa, 0x95, 0xcb, 0x41, 0x24, 0x27, 0xd8, 0x0c,
	0x57, 0x76, 0x36, 0xa0, 0x79, 0x1c, 0xb4, 0x08, 0xae, 0xe9, 0x26, 0x69, 0xc6, 0x51, 0x10, 0xda,
	0xf6, 0x99, 0x78, 0xaa, 0x11, 0x47, 0x7d, 0xdb, 0x36, 0x9f, 0x9d, 0x3b, 0x07, 0xfc, 0xf5, 0xbc,
	0x7a, 0x46, 0x3c, 0x3b, 0x86, 0xca, 0x2e, 0xad, 0xe1, 0x63, 0x10, 0x0b, 0x4d, 0x3a, 0xd4, 0x72,
	0x25, 0x8c, 0x83, 0x99, 0xb9, 0xd5, 0x94, 0x30, 0xee, 0xab, 0x58, 0xfa, 0x67, 0x84, 0x55, 0x0c,
	0x84, 0xaf, 0x7c, 0x6f, 0xf3, 0xce, 0xb8, 0xc3, 0x5d, 0xe4, 0x0e, 0xff, 0x87, 0x7b, 0xdf, 0xb9,
	0x8b, 0xab, 0x2c, 0x52, 0xba, 0x73, 0x53, 0xd6, 0x16, 0x1a, 0xf2, 0x39, 0x69, 0x1b, 0x65, 0x44,
	0x12, 0x9c, 0x0f, 0x35, 0x88, 0x21, 0x16, 0xda, 0xe1, 0x1e, 0x62, 0xfb, 0x08, 0xd9, 0x4f, 0xc6,
	0xb1, 0x94, 0xb1, 0x1c, 0x04, 0xba, 0x9c, 0xd3, 0x2e, 0xf7, 0x0a, 0x8c, 0x5b, 0xff, 0x6f, 0x13,
	0x9a, 0x8a, 0xeb, 0xa0, 0xfc, 0x2c, 0x37, 0x1a, 0x44, 0x8a, 0xad, 0xd0, 0xe1, 0x6b, 0xa9, 0xb8,
	0xfe, 0x30, 0x09, 0x9c, 0x22, 0x6e, 0x0f, 0xb4, 0x5f, 0xa7, 0xc3, 0xc4, 0xc4, 0x59, 0x02, 0xc5,
	0x50, 0xf6, 0x52, 0x71, 0xfd, 0xae, 0x80, 0xe8, 0x2b, 0xd2, 0xc1, 0x3e, 0xc9, 0xb4, 0xca, 0x40,
	0x9b, 0x1b, 0xd6, 0x42, 0x9f, 0xad, 0x4d, 0x3b, 0xb1, 0xc0, 0x79, 0xdb, 0x7e, 0x56, 0xee, 0xf6,
	0x37, 0x3e, 0xae, 0xe7, 0x06, 0x46, 0xb0, 0x1b, 0x26, 0x31, 0x48, 0x13, 0x64, 0xe7, 0xbb, 0xf6,
	0xaf, 0xf2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x57, 0xda, 0x8a, 0x3d, 0x07, 0x00, 0x00,
}
