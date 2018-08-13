package packsack

import (
	"fmt"
	"steve/structs"

	"github.com/Sirupsen/logrus"
)

//SaveGoldToDB 将玩家金币同步到DB
func SaveGoldToDB(uid uint64, goldValue int64, changeValue int64) error {

	exposer := structs.GetGlobalExposer()
	engine, err := exposer.MysqlEngineMgr.GetEngine(dbName)
	if err != nil {
		return fmt.Errorf("connect db error")
	}

	strCol := ""
	if changeValue >= 0 {

		strCol += fmt.Sprintf("%d", changeValue)

	} else {
		strCol += fmt.Sprintf("%d", -changeValue)
	}

	sql := fmt.Sprintf("update t_player_currency set %s  where playerID=?;", strCol)
	res, err := engine.Exec(sql, uid)
	if err != nil {
		logrus.Errorf("exec sql err:sql=%s,err=%s", sql, err)
		return err
	}
	if aff, err := res.RowsAffected(); aff == 0 {
		logrus.Errorf("exec sql Affect err:sql=%s,err=%s", sql, err)
		return err
	}
	return nil
}
