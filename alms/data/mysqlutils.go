package data

import (
	"fmt"
	"steve/entity/db"
	"steve/structs"

	"github.com/Sirupsen/logrus"
	"github.com/go-xorm/xorm"
)

const (
	// MysqlPlayerdbName 数据库名
	MysqlPlayerdbName = "player"
	hallInfoTableName = "t_hall_info" // 大厅信息表
)

//MysqlEnginefunc 单元测试需要
var MysqlEnginefunc = getMysqlEngineByName

func getMysqlEngineByName(mysqlName string) (*xorm.Engine, error) {
	e := structs.GetGlobalExposer()
	engine, err := e.MysqlEngineMgr.GetEngine(mysqlName)
	if err != nil {
		return nil, fmt.Errorf("获取 mysql 引擎失败：%v", err)
	}
	if err := engine.Ping(); err != nil {
		return nil, fmt.Errorf("engine ping 失败：%v", err)
	}
	return engine, nil
}

// 获取玩家救济金已领取次数
func getMysqlPlayerGotTimesByPlayerID(playerID uint64) (int, error) {
	engine, err := MysqlEnginefunc(MysqlPlayerdbName)
	if err != nil {
		return 0, err
	}
	var almsGotTimes int
	session := engine.Table(hallInfoTableName).Select("almsGotTimes").Where(fmt.Sprintf("playerID=%v", playerID))
	exist, err := session.Get(&almsGotTimes)
	if err != nil {
		sql, _ := session.LastSQL()
		return 0, fmt.Errorf("从数据库获取数据失败：(%v), sql:(%s)", err, sql)
	}
	if !exist { // 不存在插入新的
		hi := &db.THallInfo{}
		hi.Playerid = int64(playerID)
		hi.Almsgottimes = 0
		_, err := engine.Table(hallInfoTableName).Insert(hi)
		logrus.WithError(err).Debugln("插入新的数据 t_hal playerID(%d)", playerID)
		return 0, err
	}
	return almsGotTimes, nil
}

// 更改玩家救济金已领取次数
func updateMysqlPlayerGotTimesByPlayerID(playerID uint64, gotTimes int) error {
	engine, err := MysqlEnginefunc(MysqlPlayerdbName)
	if err != nil {
		return err
	}
	hi := &db.THallInfo{
		Almsgottimes: gotTimes,
	}
	session := engine.Table(hallInfoTableName).Select("almsGotTimes").Where(fmt.Sprintf("playerID=%v", playerID))
	num, err := session.Update(hi)
	if err != nil {
		sql, _ := session.LastSQL()
		return fmt.Errorf("从数据库获取数据失败：(%v), sql:(%s)", err, sql)
	}
	if num == 0 {
		return fmt.Errorf("修改玩家救济金已领取次数失败 : %v - %v", playerID, gotTimes)
	}
	return nil
}
