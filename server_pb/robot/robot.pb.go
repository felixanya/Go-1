// Code generated by protoc-gen-go. DO NOT EDIT.
// source: robot.proto

package robot

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// GameConfig 游戏玩法信息
type GameConfig struct {
	GameId  uint32 `protobuf:"varint,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	LevelId uint32 `protobuf:"varint,2,opt,name=level_id,json=levelId" json:"level_id,omitempty"`
}

func (m *GameConfig) Reset()                    { *m = GameConfig{} }
func (m *GameConfig) String() string            { return proto.CompactTextString(m) }
func (*GameConfig) ProtoMessage()               {}
func (*GameConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GameConfig) GetGameId() uint32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *GameConfig) GetLevelId() uint32 {
	if m != nil {
		return m.LevelId
	}
	return 0
}

// GameWinRate 游戏对应的胜率
type GameWinRate struct {
	Game    *GameConfig `protobuf:"bytes,1,opt,name=game" json:"game,omitempty"`
	WinRate int32       `protobuf:"varint,2,opt,name=win_rate,json=winRate" json:"win_rate,omitempty"`
}

func (m *GameWinRate) Reset()                    { *m = GameWinRate{} }
func (m *GameWinRate) String() string            { return proto.CompactTextString(m) }
func (*GameWinRate) ProtoMessage()               {}
func (*GameWinRate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GameWinRate) GetGame() *GameConfig {
	if m != nil {
		return m.Game
	}
	return nil
}

func (m *GameWinRate) GetWinRate() int32 {
	if m != nil {
		return m.WinRate
	}
	return 0
}

// WinRateRange 胜率范围
type WinRateRange struct {
	High int32 `protobuf:"varint,1,opt,name=high" json:"high,omitempty"`
	Low  int32 `protobuf:"varint,2,opt,name=low" json:"low,omitempty"`
}

func (m *WinRateRange) Reset()                    { *m = WinRateRange{} }
func (m *WinRateRange) String() string            { return proto.CompactTextString(m) }
func (*WinRateRange) ProtoMessage()               {}
func (*WinRateRange) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *WinRateRange) GetHigh() int32 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *WinRateRange) GetLow() int32 {
	if m != nil {
		return m.Low
	}
	return 0
}

// CoinsRange 金币范围
type CoinsRange struct {
	High int64 `protobuf:"varint,1,opt,name=high" json:"high,omitempty"`
	Low  int64 `protobuf:"varint,2,opt,name=low" json:"low,omitempty"`
}

func (m *CoinsRange) Reset()                    { *m = CoinsRange{} }
func (m *CoinsRange) String() string            { return proto.CompactTextString(m) }
func (*CoinsRange) ProtoMessage()               {}
func (*CoinsRange) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CoinsRange) GetHigh() int64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *CoinsRange) GetLow() int64 {
	if m != nil {
		return m.Low
	}
	return 0
}

// GetLeisureRobotInfoReq 获取空闲机器人信息请求
type GetLeisureRobotInfoReq struct {
	Game         *GameConfig   `protobuf:"bytes,1,opt,name=game" json:"game,omitempty"`
	WinRateRange *WinRateRange `protobuf:"bytes,2,opt,name=win_rate_range,json=winRateRange" json:"win_rate_range,omitempty"`
	CoinsRange   *CoinsRange   `protobuf:"bytes,3,opt,name=coins_range,json=coinsRange" json:"coins_range,omitempty"`
}

func (m *GetLeisureRobotInfoReq) Reset()                    { *m = GetLeisureRobotInfoReq{} }
func (m *GetLeisureRobotInfoReq) String() string            { return proto.CompactTextString(m) }
func (*GetLeisureRobotInfoReq) ProtoMessage()               {}
func (*GetLeisureRobotInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *GetLeisureRobotInfoReq) GetGame() *GameConfig {
	if m != nil {
		return m.Game
	}
	return nil
}

func (m *GetLeisureRobotInfoReq) GetWinRateRange() *WinRateRange {
	if m != nil {
		return m.WinRateRange
	}
	return nil
}

func (m *GetLeisureRobotInfoReq) GetCoinsRange() *CoinsRange {
	if m != nil {
		return m.CoinsRange
	}
	return nil
}

// GetLeisureRobotInfoRsp 获取空闲机器人信息响应
type GetLeisureRobotInfoRsp struct {
	RobotPlayerId uint64  `protobuf:"varint,1,opt,name=robot_player_id,json=robotPlayerId" json:"robot_player_id,omitempty"`
	Coin          int64   `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	WinRate       float64 `protobuf:"fixed64,3,opt,name=win_rate,json=winRate" json:"win_rate,omitempty"`
	ErrCode       ErrCode `protobuf:"varint,4,opt,name=err_code,json=errCode,enum=robot.ErrCode" json:"err_code,omitempty"`
}

func (m *GetLeisureRobotInfoRsp) Reset()                    { *m = GetLeisureRobotInfoRsp{} }
func (m *GetLeisureRobotInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*GetLeisureRobotInfoRsp) ProtoMessage()               {}
func (*GetLeisureRobotInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GetLeisureRobotInfoRsp) GetRobotPlayerId() uint64 {
	if m != nil {
		return m.RobotPlayerId
	}
	return 0
}

func (m *GetLeisureRobotInfoRsp) GetCoin() int64 {
	if m != nil {
		return m.Coin
	}
	return 0
}

func (m *GetLeisureRobotInfoRsp) GetWinRate() float64 {
	if m != nil {
		return m.WinRate
	}
	return 0
}

func (m *GetLeisureRobotInfoRsp) GetErrCode() ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return ErrCode_EC_SUCCESS
}

// SetRobotPlayerStateReq 設置机器人玩家状态請求
type SetRobotPlayerStateReq struct {
	RobotPlayerId uint64 `protobuf:"varint,1,opt,name=robot_player_id,json=robotPlayerId" json:"robot_player_id,omitempty"`
	NewState      bool   `protobuf:"varint,2,opt,name=new_state,json=newState" json:"new_state,omitempty"`
}

func (m *SetRobotPlayerStateReq) Reset()                    { *m = SetRobotPlayerStateReq{} }
func (m *SetRobotPlayerStateReq) String() string            { return proto.CompactTextString(m) }
func (*SetRobotPlayerStateReq) ProtoMessage()               {}
func (*SetRobotPlayerStateReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *SetRobotPlayerStateReq) GetRobotPlayerId() uint64 {
	if m != nil {
		return m.RobotPlayerId
	}
	return 0
}

func (m *SetRobotPlayerStateReq) GetNewState() bool {
	if m != nil {
		return m.NewState
	}
	return false
}

// SetRobotPlayerStateRsp 設置机器人玩家状态响应
type SetRobotPlayerStateRsp struct {
	Result  bool    `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	ErrCode ErrCode `protobuf:"varint,2,opt,name=err_code,json=errCode,enum=robot.ErrCode" json:"err_code,omitempty"`
}

func (m *SetRobotPlayerStateRsp) Reset()                    { *m = SetRobotPlayerStateRsp{} }
func (m *SetRobotPlayerStateRsp) String() string            { return proto.CompactTextString(m) }
func (*SetRobotPlayerStateRsp) ProtoMessage()               {}
func (*SetRobotPlayerStateRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *SetRobotPlayerStateRsp) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *SetRobotPlayerStateRsp) GetErrCode() ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return ErrCode_EC_SUCCESS
}

// UpdataRobotGameWinRateReq 更新机器人胜率請求
type UpdataRobotGameWinRateReq struct {
	RobotPlayerId uint64  `protobuf:"varint,1,opt,name=robot_player_id,json=robotPlayerId" json:"robot_player_id,omitempty"`
	GameId        int32   `protobuf:"varint,2,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	NewWinRate    float64 `protobuf:"fixed64,3,opt,name=newWinRate" json:"newWinRate,omitempty"`
}

func (m *UpdataRobotGameWinRateReq) Reset()                    { *m = UpdataRobotGameWinRateReq{} }
func (m *UpdataRobotGameWinRateReq) String() string            { return proto.CompactTextString(m) }
func (*UpdataRobotGameWinRateReq) ProtoMessage()               {}
func (*UpdataRobotGameWinRateReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *UpdataRobotGameWinRateReq) GetRobotPlayerId() uint64 {
	if m != nil {
		return m.RobotPlayerId
	}
	return 0
}

func (m *UpdataRobotGameWinRateReq) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *UpdataRobotGameWinRateReq) GetNewWinRate() float64 {
	if m != nil {
		return m.NewWinRate
	}
	return 0
}

// UpdataRobotGameWinRateRsp 更新机器人胜率响应
type UpdataRobotGameWinRateRsp struct {
	Result  bool    `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	ErrCode ErrCode `protobuf:"varint,2,opt,name=err_code,json=errCode,enum=robot.ErrCode" json:"err_code,omitempty"`
}

func (m *UpdataRobotGameWinRateRsp) Reset()                    { *m = UpdataRobotGameWinRateRsp{} }
func (m *UpdataRobotGameWinRateRsp) String() string            { return proto.CompactTextString(m) }
func (*UpdataRobotGameWinRateRsp) ProtoMessage()               {}
func (*UpdataRobotGameWinRateRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *UpdataRobotGameWinRateRsp) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *UpdataRobotGameWinRateRsp) GetErrCode() ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return ErrCode_EC_SUCCESS
}

// IsRobotPlayerReq 判断是否时机器人請求
type IsRobotPlayerReq struct {
	RobotPlayerId uint64 `protobuf:"varint,1,opt,name=robot_player_id,json=robotPlayerId" json:"robot_player_id,omitempty"`
}

func (m *IsRobotPlayerReq) Reset()                    { *m = IsRobotPlayerReq{} }
func (m *IsRobotPlayerReq) String() string            { return proto.CompactTextString(m) }
func (*IsRobotPlayerReq) ProtoMessage()               {}
func (*IsRobotPlayerReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *IsRobotPlayerReq) GetRobotPlayerId() uint64 {
	if m != nil {
		return m.RobotPlayerId
	}
	return 0
}

// IsRobotPlayerRsp 判断是否时机器人响应
type IsRobotPlayerRsp struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *IsRobotPlayerRsp) Reset()                    { *m = IsRobotPlayerRsp{} }
func (m *IsRobotPlayerRsp) String() string            { return proto.CompactTextString(m) }
func (*IsRobotPlayerRsp) ProtoMessage()               {}
func (*IsRobotPlayerRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *IsRobotPlayerRsp) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

// UpdataRobotGoldReq 更新机器人金币請求
type UpdataRobotGoldReq struct {
	RobotPlayerId uint64 `protobuf:"varint,1,opt,name=robot_player_id,json=robotPlayerId" json:"robot_player_id,omitempty"`
	Gold          int64  `protobuf:"varint,2,opt,name=gold" json:"gold,omitempty"`
}

func (m *UpdataRobotGoldReq) Reset()                    { *m = UpdataRobotGoldReq{} }
func (m *UpdataRobotGoldReq) String() string            { return proto.CompactTextString(m) }
func (*UpdataRobotGoldReq) ProtoMessage()               {}
func (*UpdataRobotGoldReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *UpdataRobotGoldReq) GetRobotPlayerId() uint64 {
	if m != nil {
		return m.RobotPlayerId
	}
	return 0
}

func (m *UpdataRobotGoldReq) GetGold() int64 {
	if m != nil {
		return m.Gold
	}
	return 0
}

// UpdataRobotGoldRsp 更新机器人金币响应
type UpdataRobotGoldRsp struct {
	Result  bool    `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	ErrCode ErrCode `protobuf:"varint,2,opt,name=err_code,json=errCode,enum=robot.ErrCode" json:"err_code,omitempty"`
}

func (m *UpdataRobotGoldRsp) Reset()                    { *m = UpdataRobotGoldRsp{} }
func (m *UpdataRobotGoldRsp) String() string            { return proto.CompactTextString(m) }
func (*UpdataRobotGoldRsp) ProtoMessage()               {}
func (*UpdataRobotGoldRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *UpdataRobotGoldRsp) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *UpdataRobotGoldRsp) GetErrCode() ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return ErrCode_EC_SUCCESS
}

func init() {
	proto.RegisterType((*GameConfig)(nil), "robot.GameConfig")
	proto.RegisterType((*GameWinRate)(nil), "robot.GameWinRate")
	proto.RegisterType((*WinRateRange)(nil), "robot.WinRateRange")
	proto.RegisterType((*CoinsRange)(nil), "robot.CoinsRange")
	proto.RegisterType((*GetLeisureRobotInfoReq)(nil), "robot.GetLeisureRobotInfoReq")
	proto.RegisterType((*GetLeisureRobotInfoRsp)(nil), "robot.GetLeisureRobotInfoRsp")
	proto.RegisterType((*SetRobotPlayerStateReq)(nil), "robot.SetRobotPlayerStateReq")
	proto.RegisterType((*SetRobotPlayerStateRsp)(nil), "robot.SetRobotPlayerStateRsp")
	proto.RegisterType((*UpdataRobotGameWinRateReq)(nil), "robot.UpdataRobotGameWinRateReq")
	proto.RegisterType((*UpdataRobotGameWinRateRsp)(nil), "robot.UpdataRobotGameWinRateRsp")
	proto.RegisterType((*IsRobotPlayerReq)(nil), "robot.IsRobotPlayerReq")
	proto.RegisterType((*IsRobotPlayerRsp)(nil), "robot.IsRobotPlayerRsp")
	proto.RegisterType((*UpdataRobotGoldReq)(nil), "robot.UpdataRobotGoldReq")
	proto.RegisterType((*UpdataRobotGoldRsp)(nil), "robot.UpdataRobotGoldRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RobotService service

type RobotServiceClient interface {
	GetLeisureRobotInfoByInfo(ctx context.Context, in *GetLeisureRobotInfoReq, opts ...grpc.CallOption) (*GetLeisureRobotInfoRsp, error)
	SetRobotPlayerState(ctx context.Context, in *SetRobotPlayerStateReq, opts ...grpc.CallOption) (*SetRobotPlayerStateRsp, error)
	UpdataRobotGameWinRate(ctx context.Context, in *UpdataRobotGameWinRateReq, opts ...grpc.CallOption) (*UpdataRobotGameWinRateRsp, error)
	IsRobotPlayer(ctx context.Context, in *IsRobotPlayerReq, opts ...grpc.CallOption) (*IsRobotPlayerRsp, error)
	UpdataRobotGold(ctx context.Context, in *UpdataRobotGoldReq, opts ...grpc.CallOption) (*UpdataRobotGoldRsp, error)
}

type robotServiceClient struct {
	cc *grpc.ClientConn
}

func NewRobotServiceClient(cc *grpc.ClientConn) RobotServiceClient {
	return &robotServiceClient{cc}
}

func (c *robotServiceClient) GetLeisureRobotInfoByInfo(ctx context.Context, in *GetLeisureRobotInfoReq, opts ...grpc.CallOption) (*GetLeisureRobotInfoRsp, error) {
	out := new(GetLeisureRobotInfoRsp)
	err := grpc.Invoke(ctx, "/robot.RobotService/GetLeisureRobotInfoByInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) SetRobotPlayerState(ctx context.Context, in *SetRobotPlayerStateReq, opts ...grpc.CallOption) (*SetRobotPlayerStateRsp, error) {
	out := new(SetRobotPlayerStateRsp)
	err := grpc.Invoke(ctx, "/robot.RobotService/SetRobotPlayerState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) UpdataRobotGameWinRate(ctx context.Context, in *UpdataRobotGameWinRateReq, opts ...grpc.CallOption) (*UpdataRobotGameWinRateRsp, error) {
	out := new(UpdataRobotGameWinRateRsp)
	err := grpc.Invoke(ctx, "/robot.RobotService/UpdataRobotGameWinRate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) IsRobotPlayer(ctx context.Context, in *IsRobotPlayerReq, opts ...grpc.CallOption) (*IsRobotPlayerRsp, error) {
	out := new(IsRobotPlayerRsp)
	err := grpc.Invoke(ctx, "/robot.RobotService/IsRobotPlayer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) UpdataRobotGold(ctx context.Context, in *UpdataRobotGoldReq, opts ...grpc.CallOption) (*UpdataRobotGoldRsp, error) {
	out := new(UpdataRobotGoldRsp)
	err := grpc.Invoke(ctx, "/robot.RobotService/UpdataRobotGold", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RobotService service

type RobotServiceServer interface {
	GetLeisureRobotInfoByInfo(context.Context, *GetLeisureRobotInfoReq) (*GetLeisureRobotInfoRsp, error)
	SetRobotPlayerState(context.Context, *SetRobotPlayerStateReq) (*SetRobotPlayerStateRsp, error)
	UpdataRobotGameWinRate(context.Context, *UpdataRobotGameWinRateReq) (*UpdataRobotGameWinRateRsp, error)
	IsRobotPlayer(context.Context, *IsRobotPlayerReq) (*IsRobotPlayerRsp, error)
	UpdataRobotGold(context.Context, *UpdataRobotGoldReq) (*UpdataRobotGoldRsp, error)
}

func RegisterRobotServiceServer(s *grpc.Server, srv RobotServiceServer) {
	s.RegisterService(&_RobotService_serviceDesc, srv)
}

func _RobotService_GetLeisureRobotInfoByInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeisureRobotInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).GetLeisureRobotInfoByInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.RobotService/GetLeisureRobotInfoByInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).GetLeisureRobotInfoByInfo(ctx, req.(*GetLeisureRobotInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_SetRobotPlayerState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRobotPlayerStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).SetRobotPlayerState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.RobotService/SetRobotPlayerState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).SetRobotPlayerState(ctx, req.(*SetRobotPlayerStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_UpdataRobotGameWinRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdataRobotGameWinRateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).UpdataRobotGameWinRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.RobotService/UpdataRobotGameWinRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).UpdataRobotGameWinRate(ctx, req.(*UpdataRobotGameWinRateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_IsRobotPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsRobotPlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).IsRobotPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.RobotService/IsRobotPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).IsRobotPlayer(ctx, req.(*IsRobotPlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_UpdataRobotGold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdataRobotGoldReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).UpdataRobotGold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.RobotService/UpdataRobotGold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).UpdataRobotGold(ctx, req.(*UpdataRobotGoldReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RobotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "robot.RobotService",
	HandlerType: (*RobotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeisureRobotInfoByInfo",
			Handler:    _RobotService_GetLeisureRobotInfoByInfo_Handler,
		},
		{
			MethodName: "SetRobotPlayerState",
			Handler:    _RobotService_SetRobotPlayerState_Handler,
		},
		{
			MethodName: "UpdataRobotGameWinRate",
			Handler:    _RobotService_UpdataRobotGameWinRate_Handler,
		},
		{
			MethodName: "IsRobotPlayer",
			Handler:    _RobotService_IsRobotPlayer_Handler,
		},
		{
			MethodName: "UpdataRobotGold",
			Handler:    _RobotService_UpdataRobotGold_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "robot.proto",
}

func init() { proto.RegisterFile("robot.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xc7, 0xeb, 0xd8, 0xb1, 0xdd, 0xf1, 0x47, 0xd2, 0x0d, 0x38, 0xb1, 0x4a, 0x8b, 0x59, 0x68,
	0x49, 0x7b, 0xc8, 0x41, 0xed, 0xa5, 0x3d, 0x95, 0x9a, 0x12, 0x04, 0x85, 0x86, 0x35, 0x21, 0x87,
	0xd0, 0x08, 0xc5, 0x9a, 0x38, 0x02, 0x59, 0xab, 0xac, 0xe4, 0x88, 0x40, 0x9f, 0xa4, 0x0f, 0xd1,
	0x43, 0x9f, 0xb0, 0xec, 0x6a, 0x65, 0xcb, 0xb6, 0xe4, 0xc6, 0x90, 0x8b, 0x99, 0xdd, 0xf1, 0xff,
	0x37, 0x1f, 0x3b, 0x63, 0x43, 0x4b, 0xf0, 0x6b, 0x1e, 0x9f, 0x84, 0x82, 0xc7, 0x9c, 0xec, 0xaa,
	0x83, 0x01, 0x18, 0xcc, 0xa6, 0xe9, 0x15, 0xfd, 0x02, 0x70, 0xea, 0x4c, 0x71, 0xc8, 0x83, 0x1b,
	0x6f, 0x42, 0x0e, 0xa1, 0x31, 0x71, 0xa6, 0x68, 0x7b, 0xee, 0x51, 0x65, 0x50, 0x39, 0xee, 0xb0,
	0xba, 0x3c, 0x5a, 0x2e, 0xe9, 0x43, 0xd3, 0xc7, 0x7b, 0xf4, 0xa5, 0x67, 0x47, 0x79, 0x1a, 0xea,
	0x6c, 0xb9, 0xf4, 0x07, 0xb4, 0x24, 0xe1, 0xc2, 0x0b, 0x98, 0x13, 0x23, 0x79, 0x03, 0x35, 0xa9,
	0x51, 0xfa, 0x96, 0xf9, 0xe2, 0x24, 0x8d, 0xbf, 0x88, 0xc1, 0x94, 0x5b, 0x02, 0x13, 0x2f, 0xb0,
	0x85, 0x13, 0xa3, 0x02, 0xee, 0xb2, 0x46, 0x92, 0x12, 0xe8, 0x47, 0x68, 0x6b, 0x18, 0x73, 0x82,
	0x09, 0x12, 0x02, 0xb5, 0x5b, 0x6f, 0x72, 0xab, 0x88, 0xbb, 0x4c, 0xd9, 0x64, 0x1f, 0xaa, 0x3e,
	0x4f, 0xb4, 0x52, 0x9a, 0xd4, 0x04, 0x18, 0x72, 0x2f, 0x88, 0xd6, 0x35, 0xd5, 0x75, 0x4d, 0x35,
	0xd5, 0xfc, 0xa9, 0x40, 0xef, 0x14, 0xe3, 0xef, 0xe8, 0x45, 0x33, 0x81, 0x4c, 0x66, 0x6a, 0x05,
	0x37, 0x9c, 0xe1, 0xdd, 0x63, 0xcb, 0xf8, 0x04, 0xdd, 0xac, 0x0c, 0x5b, 0xc8, 0xc8, 0x0a, 0xdf,
	0x32, 0x0f, 0xb4, 0x20, 0x5f, 0x08, 0x6b, 0x27, 0xf9, 0xb2, 0x4c, 0x68, 0x8d, 0x65, 0xc2, 0x5a,
	0x57, 0x5d, 0x0a, 0xb4, 0x28, 0x85, 0xc1, 0x78, 0x6e, 0xd3, 0xdf, 0x25, 0x09, 0x47, 0x21, 0x79,
	0x0b, 0x7b, 0x4a, 0x6a, 0x87, 0xbe, 0xf3, 0x80, 0x22, 0x7b, 0xc2, 0x1a, 0xeb, 0xa8, 0xeb, 0x33,
	0x75, 0x6b, 0xb9, 0xb2, 0x33, 0x12, 0xa8, 0xdb, 0xa0, 0xec, 0xa5, 0xc7, 0x90, 0x79, 0x54, 0xe6,
	0x8f, 0x41, 0xde, 0x41, 0x13, 0x85, 0xb0, 0xc7, 0xdc, 0xc5, 0xa3, 0xda, 0xa0, 0x72, 0xdc, 0x35,
	0xbb, 0x3a, 0xc5, 0x6f, 0x42, 0x0c, 0xb9, 0x8b, 0xac, 0x81, 0xa9, 0x41, 0x7f, 0x42, 0x6f, 0x84,
	0x31, 0x5b, 0x44, 0x1b, 0xc5, 0xb2, 0x58, 0xbc, 0x7b, 0x74, 0x6e, 0x2f, 0xe1, 0x79, 0x80, 0x89,
	0x1d, 0xc5, 0xd9, 0x54, 0x34, 0x59, 0x33, 0xc0, 0x44, 0x71, 0xe8, 0x65, 0x31, 0x3e, 0x0a, 0x49,
	0x0f, 0xea, 0x02, 0xa3, 0x99, 0x1f, 0x2b, 0x6a, 0x93, 0xe9, 0xd3, 0x52, 0xee, 0x3b, 0x9b, 0x73,
	0xff, 0x05, 0xfd, 0xf3, 0xd0, 0x75, 0x62, 0x47, 0xf1, 0x73, 0xf3, 0xbc, 0x4d, 0xfa, 0xb9, 0xed,
	0x49, 0x07, 0x33, 0xdb, 0x9e, 0xd7, 0x00, 0x01, 0x26, 0x9a, 0xa8, 0x3b, 0x9c, 0xbb, 0xa1, 0x57,
	0xa5, 0xd1, 0x9f, 0xa6, 0xba, 0xcf, 0xb0, 0x6f, 0x45, 0xb9, 0xce, 0x6d, 0x51, 0x14, 0x7d, 0xbf,
	0xaa, 0x2d, 0x4f, 0x89, 0x9e, 0x01, 0xc9, 0xd7, 0xc1, 0x7d, 0x77, 0x9b, 0xf6, 0x11, 0xa8, 0x4d,
	0xb8, 0xef, 0x66, 0x93, 0x29, 0x6d, 0x7a, 0xb1, 0x4e, 0x7c, 0x92, 0x96, 0x98, 0x7f, 0xab, 0xd0,
	0x56, 0xcc, 0x11, 0x8a, 0x7b, 0x6f, 0x8c, 0xe4, 0x12, 0xfa, 0x05, 0x9b, 0xf5, 0xf5, 0x41, 0x7e,
	0x92, 0x57, 0xd9, 0xfe, 0x17, 0xfe, 0x58, 0x18, 0x9b, 0xdc, 0x51, 0x48, 0x9f, 0x91, 0x73, 0x38,
	0x28, 0x98, 0xdd, 0x39, 0xb6, 0x78, 0x6d, 0x8c, 0x4d, 0x6e, 0x85, 0xbd, 0x82, 0x5e, 0xf1, 0xdc,
	0x90, 0x81, 0x96, 0x96, 0x0e, 0xb5, 0xf1, 0x9f, 0x6f, 0x28, 0xfe, 0x10, 0x3a, 0x4b, 0x6f, 0x4f,
	0x0e, 0xb5, 0x68, 0x75, 0x9a, 0x8c, 0x62, 0x87, 0x82, 0x58, 0xb0, 0xb7, 0xf2, 0x84, 0xa4, 0x5f,
	0x10, 0x3b, 0x1d, 0x16, 0xa3, 0xcc, 0x25, 0x51, 0xd7, 0x75, 0xf5, 0x9f, 0xf5, 0xe1, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x2a, 0xc9, 0xca, 0xcb, 0xd5, 0x06, 0x00, 0x00,
}
