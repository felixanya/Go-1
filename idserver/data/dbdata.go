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

*/

const dbName = "player"

const TagPlayerId = 1
const TagShowId = 2

// 从DB生成一个playerId
func NewPlayerIdFromDB() (uint64, error) {

	exposer := structs.GetGlobalExposer()
	engine, err := exposer.MysqlEngineMgr.GetEngine(dbName)
	if err != nil {
		return 0, fmt.Errorf("connect db error")
	}

	sql := fmt.Sprintf("select n_value from t_param  where n_id='%d';", TagPlayerId)
	res, err := engine.QueryString(sql)
	if err != nil {
		return 0, err
	}
	if len(res) == 0 {
		return 0, fmt.Errorf("player id param no exist")
	}
	row := res[0]

	id, err := strconv.ParseUint(row["n_value"], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("player id parse error")
	}

	sql = fmt.Sprintf("update t_param set n_value=n_value+1 where n_id='%d';", TagPlayerId)

	resUpdate, err := engine.Exec(sql)
	if err != nil {
		return 0, err
	}
	if aff, err := resUpdate.RowsAffected(); aff == 0 {
		// 如果插入行=0，表明插入失败
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

	sql := fmt.Sprintf("select n_showid  from t_show_id  where n_isUse='%d';", 0)
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
