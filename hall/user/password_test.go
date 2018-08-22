package user

import (
	"fmt"
	"steve/client_pb/hall"
	"steve/entity/db"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"

	"github.com/Sirupsen/logrus"
)

// --------------------------------------------
// 由于平台接口不容易 mock， 所以主要查看输出来测试逻辑
// -------------------------------------------

func init() {
	logrus.SetLevel(logrus.DebugLevel)

}

// TestHandleChangePasswordReq 修改密码测试
func TestHandleChangePasswordReq(t *testing.T) {
	viper.SetDefault("change_password_url", "http://192.168.7.26:18101/account/checkPwd")
	dbPlayerGetter = func(playerID uint64, fields ...string) (*db.TPlayer, error) {
		fmt.Printf("获取玩家信息:playerID=%d, fields=%v\n", playerID, fields)
		return &db.TPlayer{Accountid: 100}, nil
	}
	HandleChangePasswordReq(1, nil, hall.ChangePasswordReq{
		OldPasswd: proto.String("123"), NewPasswd: proto.String("456"),
	})
}

// 重置密码测试
func TestHandleResetPasswordReq(t *testing.T) {
	viper.SetDefault("reset_password_url", "http://192.168.7.26:18101/account/checkPwd")
	dbPlayerGetter = func(playerID uint64, fields ...string) (*db.TPlayer, error) {
		fmt.Printf("获取玩家信息:playerID=%d, fields=%v\n", playerID, fields)
		return &db.TPlayer{Accountid: 100}, nil
	}
	HandleResetPasswordReq(1, nil, hall.ResetPasswordReq{
		Phone:     proto.String("15234562156"),
		DymcCode:  proto.String("1234"),
		NewPasswd: proto.String("989"),
	})
}

// 校验密码测试
func TestHandleCheckPasswordReq(t *testing.T) {
	viper.SetDefault("check_password_url", "http://192.168.7.26:18101/account/checkPwd")
	dbPlayerGetter = func(playerID uint64, fields ...string) (*db.TPlayer, error) {
		fmt.Printf("获取玩家信息:playerID=%d, fields=%v\n", playerID, fields)
		return &db.TPlayer{Accountid: 100}, nil
	}
	HandleCheckPasswordReq(1, nil, hall.CheckPasswordReq{
		Password: proto.String("123"),
	})
}
