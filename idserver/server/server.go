package server

/*
	功能： RPC服务实现类，完成当前模块服务的所有RPC接口实现和处理
    作者： SkyWang
    日期： 2018-7-24

*/

import (
	"context"
	"github.com/Sirupsen/logrus"
	"steve/server_pb/idserver"
	"steve/idserver/logic"
)

// PropsServer 实现 props.PropsServer
type IdServer struct{}

var _ idsvr.IdserviceServer = new(IdServer)

// 生成一个新的playerid 和 showid
func (gs *IdServer) NewPlayerShowId(ctx context.Context, request *idsvr.NewPlayerShowIdReq) (response *idsvr.NewPlayerShowIdRsp, err error) {
	logrus.Debugln("NewPlayerShowId req", *request)
	response = &idsvr.NewPlayerShowIdRsp{}
	response.ErrCode = idsvr.ResultStat_FAILED

	// 参数检查

	// 调用逻辑实现代码
	pid, sid, err := logic.NewPlayerShowId()

	// 逻辑代码返回错误
	if err != nil {
		response.ErrCode = idsvr.ResultStat_FAILED
		logrus.WithError(err).Errorln("NewPlayerShowId resp", *response)
		return response, nil
	}
	// 设置返回值
	response.PlayerId = pid
	response.ShowId = sid

	response.ErrCode = idsvr.ResultStat_SUCCEED
	logrus.Debugln("NewPlayerShowId resp", *response)
	return response, nil
}

// 生成一个新的playerid
func (gs *IdServer) NewPlayerId(ctx context.Context, request *idsvr.NewPlayerIdReq) (response *idsvr.NewPlayerIdRsp, err error) {
	logrus.Debugln("NewPlayerId req", *request)
	response = &idsvr.NewPlayerIdRsp{}
	response.ErrCode = idsvr.ResultStat_FAILED

	// 参数检查

	// 调用逻辑实现代码
	id, err := logic.NewPlayerId()

	// 逻辑代码返回错误
	if err != nil {
		response.ErrCode = idsvr.ResultStat_FAILED
		logrus.WithError(err).Errorln("NewPlayerId resp", *response)
		return response, nil
	}
	// 设置返回值
	response.NewId = id

	response.ErrCode = idsvr.ResultStat_SUCCEED
	logrus.Debugln("NewPlayerId resp", *response)
	return response, nil
}

// 生成一个新的showid
func (gs *IdServer) NewShowId(ctx context.Context, request *idsvr.NewShowIdReq) (response *idsvr.NewShowIdRsp, err error) {
	logrus.Debugln("NewShowId req", *request)
	response = &idsvr.NewShowIdRsp{}
	response.ErrCode = idsvr.ResultStat_FAILED

	// 参数检查

	// 调用逻辑实现代码
	id, err := logic.NewShowId()

	// 逻辑代码返回错误
	if err != nil {
		response.ErrCode = idsvr.ResultStat_FAILED
		logrus.WithError(err).Errorln("NewShowId resp", *response)
		return response, nil
	}
	// 设置返回值
	response.NewId = id

	response.ErrCode = idsvr.ResultStat_SUCCEED
	logrus.Debugln("NewShowId resp", *response)
	return response, nil
}


