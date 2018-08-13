package data

import (
	"github.com/Sirupsen/logrus"
)

//UpdatePlayerGotTimesByPlayerID 根据玩家ID修改玩家已经领取数量(db,redis)
func UpdatePlayerGotTimesByPlayerID(playerID uint64, changeTimes int) error {
	entry := logrus.WithFields(logrus.Fields{
		"func_name":   "getAlmsConfigByPlayerID",
		"changeTimes": changeTimes,
	})
	// redis
	if err := UpdateAlmsPlayerGotTimes(playerID, changeTimes, RedisTimeOut); err != nil {
		entry.WithError(err).Errorln("修改玩家已经领取数量 redis 失败 playerID(%v)", playerID)
		return err
	}
	//db
	if err := updateMysqlPlayerGotTimesByPlayerID(playerID, changeTimes); err != nil {
		entry.WithError(err).Errorln("修改玩家已经领取数量 DB 失败 playerID(%v)", playerID)
		return err
	}
	entry.Infoln("修改玩家已经领取数量")
	return nil
}

//GetAlmsConfigByPlayerID 根据玩家ID获取救济金配置
func GetAlmsConfigByPlayerID(playerID uint64) (*AlmsConfig, error) {
	// entry := logrus.WithFields(logrus.Fields{
	// 	"func_name": "getAlmsConfigByPlayerID",
	// })
	return nil, nil
}

//GetPlayerGotTimesByPlayerID 根据玩家ID获取救济金已领取数量，先从redis,不存在从db，再存入redis
func GetPlayerGotTimesByPlayerID(playerID uint64) (int, error) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "GetPlayerGotTimesByPlayerID",
	})
	// redis获取当前玩家救济领取数量
	times, err := GetAlmsPlayerGotTimes(playerID)
	if err != nil {
		entry.WithError(err).Warnf("警告:从redis获取失败,重新从db获取数据 playerID(%v)", playerID)
		// 从t_hall_info数据库取数据
		times, err := getMysqlPlayerGotTimesByPlayerID(playerID) //t_hall_info 可能不存在该玩家id,
		if err != nil {
			entry.WithError(err).Errorln("获取救济金已领取数量失败")
			return 0, err
		}
		// 存入redis
		if err = UpdateAlmsPlayerGotTimes(playerID, times, RedisTimeOut); err != nil {
			entry.WithError(err).Errorln("存储玩家救济金领取次数失败")
		}
	}
	return times, nil
}
