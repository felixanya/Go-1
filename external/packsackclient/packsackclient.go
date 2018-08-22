package packsackclient

import (
	"context"
	"errors"
	"fmt"

	"steve/server_pb/alms"
	"steve/structs"
	"steve/structs/common"

	"google.golang.org/grpc"
)

/*
	功能： 背包服的Client API封装,实现调用
	作者： wuhongwei
	日期： 2018-8-03
*/

// GetPacksackGold 获取背包金币
func GetPacksackGold(playerID uint64) (int64, error) {
	// 得到服务连接
	con, err := getPacksackServer()
	if err != nil || con == nil {
		return 0, errors.New("no connection")
	}
	// 新建Client
	client := alms.NewPacksackServerClient(con)
	// 调用RPC方法
	rsp, err := client.GetPacksackGold(context.Background(), &alms.PacksackGetGoldReq{
		PlayerId: playerID,
	})
	// 检测返回值
	if err != nil {
		return 0, err
	}
	if !rsp.Result {
		return 0, fmt.Errorf("获取失败")
	}
	return rsp.GetPacksackGold(), nil
}

func getPacksackServer() (*grpc.ClientConn, error) {
	e := structs.GetGlobalExposer()
	con, err := e.RPCClient.GetConnectByServerName(common.AlmsServiceName)
	if err != nil || con == nil {
		return nil, errors.New("no connection")
	}
	return con, nil
}
