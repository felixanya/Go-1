package packsack_utils

import (
	"fmt"
	"steve/entity/db"
	"steve/structs"

	"github.com/go-xorm/xorm"
)

const (
	dbName      = "player"
	pktableName = "t_player_packsack"
)

//Mydb 单元测试需要
var Mydb = getDBByName

func getDBByName(mysqlName string) (*xorm.Engine, error) {
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

// 初始化背包信息表
func initTPacksack(uid uint64) error {
	engine, err := Mydb(dbName)
	if err != nil {
		return err
	}
	tp := &db.TPlayerPacksack{
		Playerid: int64(uid),
		Gold:     0,
	}
	num, err := engine.Table(pktableName).Insert(tp)
	if err != nil {
		return err
	}
	if num == 0 {
		return fmt.Errorf("初始化背包信息表失败 playerID(%d)", uid)
	}
	return nil
}

//GetGoldFromDB 从DB获取玩家背包金币
func GetGoldFromDB(uid uint64) (int64, error) {
	engine, err := Mydb(dbName)
	if err != nil {
		return 0, err
	}
	tppk := &db.TPlayerPacksack{}
	session := engine.Table(pktableName).Where(fmt.Sprintf("playerID=%d", uid)).Select("gold")
	exist, err := session.Get(tppk)
	if err != nil {
		sql, _ := session.LastSQL()
		return 0, fmt.Errorf("从DB获取玩家背包金币金币：(%v), sql:(%s)", err, sql)
	}
	if !exist {
		return 0, initTPacksack(uid)
	}
	return int64(tppk.Gold), nil
}

//SaveGoldToDB 将玩家金币设置到DBtppk
func SaveGoldToDB(uid uint64, changeValue int64) error {
	engine, err := Mydb(dbName)
	if err != nil {
		return err
	}
	tpk := &db.TPlayerPacksack{
		Gold: int(changeValue),
	}
	session := engine.Table(pktableName).Where(fmt.Sprintf("playerID=%d", uid))
	sum, err := session.Update(tpk)
	if err != nil {
		sql, _ := session.LastSQL()
		return fmt.Errorf("将玩家金币设置到DB失败：(%v), sql:(%s)", err, sql)
	}
	if sum == 0 {
		return fmt.Errorf("修改背包金币失败 playerID(%d) changeValue(%d)", uid, changeValue)
	}
	return nil
}
