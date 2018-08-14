package user

import (
	"fmt"
	"steve/client_pb/hall"
	"steve/entity/db"
	"testing"

	"github.com/Sirupsen/logrus"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
)

/*
由于验证码发送和校验不容易 mock， 所以此单元测试主要通过输出来查看逻辑是否正确
*/

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	// 测试不加载配置
	loadBindPhoneRewardConfig = func(int) {}

}

func TestHandleSendAuthCodeReq(t *testing.T) {
	viper.SetDefault("send_code_url", "http://192.168.7.26:8086/mock/24/account/sendCode")
	fmt.Printf("%#v", HandleSendAuthCodeReq(1, nil, hall.AuthCodeReq{
		CellphoneNum: proto.Uint64(13521653125),
		SendCase:     hall.AuthCodeSendScene_BIND_PHONE.Enum(),
	}))
}

func TestCheckAuthCodeReq(t *testing.T) {
	viper.SetDefault("check_code_url", "http://192.168.7.26:8086/mock/24/account/checkCode")
	HandleCheckAuthCodeReq(1, nil, hall.CheckAuthCodeReq{
		SendCase: hall.AuthCodeSendScene_BIND_WECHAT.Enum(),
		Code:     proto.String("2125"),
		Phone:    proto.String("13512321231"),
	})
}

func TestHandleBindPhoneReq(t *testing.T) {
	viper.SetDefault("bind_phone_url", "http://192.168.7.26:8086/mock/24/account/bindPhone")
	dbPlayerGetter = func(playerID uint64, fields ...string) (*db.TPlayer, error) {
		fmt.Printf("获取玩家信息:playerID=%d, fields=%v\n", playerID, fields)
		return &db.TPlayer{Accountid: 100, Phone: ""}, nil
	}
	dbPlayerSetter = func(playerID uint64, fields []string, dbPlayer *db.TPlayer) error {
		fmt.Printf("设置玩家信息：playerID=%d, fields=%v, player=%#v\n", playerID, fields, dbPlayer)
		return nil
	}
	goldAdder = func(uid uint64, goldType int16, changeValue int64, funcId int32, channel int64, gameId int32, level int32) (int64, error) {
		fmt.Printf("添加金币： uid=%d ,goldType=%d, changeValue=%d\n", uid, goldType, changeValue)
		return 0, nil
	}
	HandleBindPhoneReq(1, nil, hall.BindPhoneReq{
		Phone:    proto.String("13512321231"),
		DymcCode: proto.String("2252"),
		Passwd:   proto.String("abcd"),
	})
}

func TestHandleChangePhoneReq(t *testing.T) {
	viper.SetDefault("change_phone_url", "http://192.168.7.26:8086/mock/24/account/resetPhone")
	dbPlayerGetter = func(playerID uint64, fields ...string) (*db.TPlayer, error) {
		fmt.Printf("获取玩家信息:playerID=%d, fields=%v\n", playerID, fields)
		return &db.TPlayer{Accountid: 100, Phone: "13526785555"}, nil
	}
	dbPlayerSetter = func(playerID uint64, fields []string, dbPlayer *db.TPlayer) error {
		fmt.Printf("设置玩家信息：playerID=%d, fields=%v, player=%#v\n", playerID, fields, dbPlayer)
		return nil
	}
	HandleChangePhoneReq(1, nil, hall.ChangePhoneReq{
		OldPhone:     proto.String("13526785555"),
		NewPhone:     proto.String("13526785555"),
		OldPhoneCode: proto.String("1234"),
		NewPhoneCode: proto.String("1234"),
	})
}
