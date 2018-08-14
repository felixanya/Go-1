package data

import (
	"errors"
	"steve/external/configclient"

	"github.com/Sirupsen/logrus"
)

//MyAlmsConfig  救济金配置
type MyAlmsConfig struct {
	GetNorm          int64 // 救济线
	GetTimes         int   // 最多领取次数
	GetNumber        int64 // 领取数量
	AlmsCountDonw    int   // 救济倒计时，时间是秒
	DepositCountDonw int   // 快充倒计时，时间是秒
	PlayerGotTimes   int   // 玩家已领取数量
	Version          int   // 救济金配置表版本号,初始1
}

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
	entry.Debugln("修改玩家已经领取数量")
	return nil
}

//GetPlayerGotTimesByPlayerID 根据玩家ID获取救济金已领取数量，先从redis,不存在从db，再存入redis
func GetPlayerGotTimesByPlayerID(playerID uint64) (int, error) {
	entry := logrus.WithFields(logrus.Fields{
		"func_name": "GetPlayerGotTimesByPlayerID",
	})
	// redis获取当前玩家救济领取数量
	times, err := GetAlmsPlayerGotTimes(playerID)
	if err == nil {
		return times, nil
	}
	entry.WithError(err).Debugln("重新从db获取已领取次数")
	dbtimes, dberr := getMysqlPlayerGotTimesByPlayerID(playerID) //t_hall_info 可能不存在该玩家id,
	if dberr != nil {
		entry.WithError(dberr).Errorln("获取救济金已领取数量失败")
		return 0, dberr
	}
	// 存入redis
	if err := UpdateAlmsPlayerGotTimes(playerID, dbtimes, RedisTimeOut); err != nil {
		entry.WithError(err).Errorln("存储玩家救济金领取次数失败")
	}
	return dbtimes, nil
}

//GetAlmsConfigByPlayerID 根据玩家ID获取救济金配置
func GetAlmsConfigByPlayerID(playerID uint64) (*MyAlmsConfig, error) {
	// 获取玩家已经领取次数
	gotTimes, err := GetPlayerGotTimesByPlayerID(playerID)
	if err != nil {
		return nil, err
	}
	// 获取救济金配置
	almsConfigMap, err := configclient.GetAlmsConfigMap()
	if err != nil {
		logrus.WithError(err).Debugln("获取救济金配置失败")
		return nil, err
	}
	if len(almsConfigMap) == 0 {
		return nil, errors.New("救济金配置为0")
	}
	acf := almsConfigMap[0]
	resultAlmsConfig := &MyAlmsConfig{
		GetNorm:          int64(acf.GetNorm),
		GetNumber:        int64(acf.GetNumber),
		GetTimes:         acf.GetTimes,
		AlmsCountDonw:    acf.AlmsCountDown,
		DepositCountDonw: acf.DepositCountDown,
		Version:          acf.Version,
		PlayerGotTimes:   gotTimes,
	}
	return resultAlmsConfig, nil
}
