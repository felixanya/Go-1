package almsserver

import (
	"context"
	"steve/alms/packsack/packsack_gold"
	s_alms "steve/server_pb/alms"

	"github.com/Sirupsen/logrus"
)

// PacksackServer 背包服
type PacksackServer struct{}

var _ s_alms.PacksackServerServer = new(PacksackServer)

// GetPacksackGold 根据账号获取玩家 ID
func (ps *PacksackServer) GetPacksackGold(ctx context.Context, req *s_alms.PacksackGetGoldReq) (*s_alms.PacksackGetGoldRsp, error) {
	logrus.Debugf("GetPacksackGold req: (%v)", *req)
	rsp := &s_alms.PacksackGetGoldRsp{}
	playerID := req.GetPlayerId()
	pkgold, err := packsack_gold.GetGoldMgr().GetGold(playerID)
	if err != nil {
		logrus.WithError(err).Debugln("获取背包金币失败")
		rsp.Result = false
		return rsp, err
	}
	rsp.Result = true
	rsp.PacksackGold = pkgold
	return rsp, nil
}
