package data

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"steve/structs"
	"strconv"
)

/*
	功能： 服务数据保存到Mysql.
	作者： SkyWang
	日期： 2018-8-15

CREATE TABLE `t_show_id` (
  `n_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '递增ID',
  `n_showid` bigint(20) NOT NULL COMMENT 'show id 值',
  `n_isUse` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被使用',
  PRIMARY KEY (`n_id`),
  UNIQUE KEY `t_show_id_UN_showid` (`n_showid`),
  KEY `t_show_id_n_isUse_IDX` (`n_isUse`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='showid记录表'

CREATE TABLE `t_player_id` (
  `n_id` bigint(20) NOT NULL COMMENT '通用变量ID',
  `n_value` bigint(20) DEFAULT '0' COMMENT '变量值',
  `n_des` varchar(255) DEFAULT NULL COMMENT '变量描述',
  PRIMARY KEY (`n_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='playerid表'

*/

const dbName = "player"

const TagPlayerId = 1

// 开始号码ID
var playerIdBegin = 100000




// 获取目前可用号码总数
func GetCanUseSumFromDB() (uint64, error) {
	exposer := structs.GetGlobalExposer()
	engine, err := exposer.MysqlEngineMgr.GetEngine(dbName)
	if err != nil {
		return 0, fmt.Errorf("connect db error")
	}

	sql := fmt.Sprintf("select count(n_id) as 'can_sum' from t_show_id  where n_isUse='%d';", 0)
	res, err := engine.QueryString(sql)
	if err != nil {
		return 0, err
	}
	if len(res) == 0 {
		return 0, fmt.Errorf("can sum param no exist")
	}
	row := res[0]

	sum, err := strconv.ParseUint(row["can_sum"], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("can sum parse error")
	}
	logrus.Debugf("GetCanUseSumFromDB: sum=%d", sum)
	return sum, nil
}

func initPlayerIdToDB() (uint64, error) {
	exposer := structs.GetGlobalExposer()
	engine, err := exposer.MysqlEngineMgr.GetEngine(dbName)
	if err != nil {
		return 0, fmt.Errorf("connect db error")
	}

	sql := fmt.Sprintf("insert into t_player_id (n_id, n_value, n_des) values('%d', '%d','%s');", TagPlayerId, playerIdBegin+1, "playerId计数")
	resUpdate, err := engine.Exec(sql)
	if err != nil {
		return 0, err
	}
	if aff, err := resUpdate.RowsAffected(); aff == 0 {
		// 如果插入行=0，表明插入失败
		return 0, fmt.Errorf("init player id failed:%v", err)
	}

	id := uint64(playerIdBegin)
	logrus.Debugf("get init player id  win: id=%d", id)

	return id, nil

}

// 从DB生成一个playerId
func NewPlayerIdFromDB() (uint64, error) {

	exposer := structs.GetGlobalExposer()
	engine, err := exposer.MysqlEngineMgr.GetEngine(dbName)
	if err != nil {
		return 0, fmt.Errorf("connect db error")
	}

	sql := fmt.Sprintf("select n_value from t_player_id  where n_id='%d';", TagPlayerId)
	res, err := engine.QueryString(sql)
	if err != nil {
		return 0, err
	}
	if len(res) == 0 {
		// 如果没有初始化数据，直接插入第一个ID
		return initPlayerIdToDB()
	}
	row := res[0]

	id, err := strconv.ParseUint(row["n_value"], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("player id parse error")
	}

	sql = fmt.Sprintf("update t_player_id set n_value=n_value+1 where n_id='%d';", TagPlayerId)

	resUpdate, err := engine.Exec(sql)
	if err != nil {
		return 0, err
	}
	if aff, err := resUpdate.RowsAffected(); aff == 0 {
		// 如果影响行=0，表明修改失败
		return 0, fmt.Errorf("inc player id failed:%v", err)
	}
	id = uint64(id)
	logrus.Debugf("get player id  win: id=%d", id)

	return id, nil
}

// 从DB生成一个showId
func NewShowIdFromDB() (uint64, error) {

	exposer := structs.GetGlobalExposer()
	engine, err := exposer.MysqlEngineMgr.GetEngine(dbName)
	if err != nil {
		return 0, fmt.Errorf("connect db error")
	}

	sql := fmt.Sprintf("select n_showid  from t_show_id  where n_isUse='%d' limit 1;", 0)
	res, err := engine.QueryString(sql)
	if err != nil {
		return 0, err
	}
	if len(res) == 0 {
		return 0, fmt.Errorf("good show id is empty")
	}
	row := res[0]

	id, err := strconv.ParseUint(row["n_showid"], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("show id parse error")
	}

	sql = fmt.Sprintf("update t_show_id set n_isUse=1 where n_showid='%d';", id)
	resUpdate, err := engine.Exec(sql)
	if err != nil {
		return 0, err
	}
	if aff, err := resUpdate.RowsAffected(); aff == 0 {
		// 如果插入行=0，表明插入失败
		return 0, fmt.Errorf("inc show id failed:%v", err)
	}

	id = uint64(id)
	logrus.Debugf("get show id  win: id=%d", id)

	return id, nil
}

// 将新增ShowId同步到DB
func InsertShowId(uids []string) error {

	exposer := structs.GetGlobalExposer()
	engine, err := exposer.MysqlEngineMgr.GetEngine(dbName)
	if err != nil {
		return fmt.Errorf("connect db error")
	}

	cols := ""
	for _, id := range uids {
		if len(cols) > 0 {
			cols += ","
		}
		cols += fmt.Sprintf("('%s')", id)
	}

	sql := fmt.Sprintf("insert into t_show_id (n_showid) values%s;", cols)
	logrus.Debugf("InsertShowId sql:%s", sql)
	res, err := engine.Exec(sql)
	if err != nil {
		return err
	}

	if aff, err := res.RowsAffected(); aff == 0 {
		// 如果插入行=0，表明插入失败
		return err
	}
	return nil
}
