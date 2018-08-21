package idclient

/*
	功能：idserver的Client API封装,实现调用
	作者： SkyWang
	日期： 2018-8-20
*/

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"steve/server_pb/idserver"
	"steve/structs"
)

/*
方法：生成一个playerid和showid
参数：
返回：playerid,showid,错误信息
*/
func NewPlayerShowId() (uint64, uint64, error) {

	// 得到服务连接
	con, err := getMyServer()
	if err != nil || con == nil {
		return 0, 0, errors.New("no idserver connection")
	}

	req := new(idsvr.NewPlayerShowIdReq)

	// 新建Client
	client := idsvr.NewIdserviceClient(con)
	// 调用RPC方法
	rsp, err := client.NewPlayerShowId(context.Background(), req)
	// 检测返回值
	if err != nil {
		return 0, 0, err
	}

	if rsp.ErrCode != idsvr.ResultStat_SUCCEED {
		return 0, 0, errors.New("NewPlayerShowId failed")
	}
	return rsp.GetPlayerId(), rsp.GetShowId(), nil
}

/*
方法：生成一个playerid
参数：
返回：playerid,错误信息
*/
func NewPlayerId() (uint64, error) {

	// 得到服务连接
	con, err := getMyServer()
	if err != nil || con == nil {
		return 0,  errors.New("no idserver connection")
	}

	req := new(idsvr.NewPlayerIdReq)

	// 新建Client
	client := idsvr.NewIdserviceClient(con)
	// 调用RPC方法
	rsp, err := client.NewPlayerId(context.Background(), req)
	// 检测返回值
	if err != nil {
		return 0,  err
	}

	if rsp.ErrCode != idsvr.ResultStat_SUCCEED {
		return 0,  errors.New("NewPlayerId failed")
	}
	return rsp.GetNewId(), nil
}

/*
方法：生成一个showid
参数：
返回：showid,错误信息
*/
func NewShowId() (uint64, error) {

	// 得到服务连接
	con, err := getMyServer()
	if err != nil || con == nil {
		return 0,  errors.New("no idserver connection")
	}

	req := new(idsvr.NewShowIdReq)

	// 新建Client
	client := idsvr.NewIdserviceClient(con)
	// 调用RPC方法
	rsp, err := client.NewShowId(context.Background(), req)
	// 检测返回值
	if err != nil {
		return 0,  err
	}

	if rsp.ErrCode != idsvr.ResultStat_SUCCEED {
		return 0,  errors.New("NewShowId failed")
	}
	return rsp.GetNewId(), nil
}


// 根据服务名生成服务连接获取方式
func getMyServer() (*grpc.ClientConn, error) {

	e := structs.GetGlobalExposer()
	// 得到服务连接
	con, err := e.RPCClient.GetConnectByServerName("idserver")
	if err != nil || con == nil {
		return nil, errors.New("no connection")
	}

	return con, nil
}
