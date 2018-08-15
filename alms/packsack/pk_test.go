package packsack

import (
	"fmt"
	"steve/alms/packsack/packsack_gold"
	"testing"
)

func init() {
	// packsack_utils.Mydb = func(mysqlName string) (*xorm.Engine, error) {
	// 	conf := mysql.Config{
	// 		User:                 "root",
	// 		Passwd:               "123456",
	// 		Net:                  "tcp",
	// 		Addr:                 "192.168.8.210:3306",
	// 		DBName:               "steve",
	// 		Params:               map[string]string{"charset": "utf8"},
	// 		AllowNativePasswords: true,
	// 	}
	// 	engine, _ := xorm.NewEngine("mysql", conf.FormatDSN())

	// 	if err := engine.Ping(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	return engine, nil
	// }
	// packsack_utils.Myredis = func() *redis.Client {
	// 	myredis := redis.NewClient(&redis.Options{
	// 		Addr:     "127.0.0.1:6379",
	// 		Password: "",
	// 		DB:       0,
	// 	})
	// 	cmd := myredis.Ping()
	// 	if cmd.Err() != nil {
	// 		fmt.Println("连接 redis 失败")
	// 	}
	// 	return myredis
	// }
}

func Test_Get(t *testing.T) {
	gold, err := packsack_gold.GetGoldMgr().GetGold(11)
	fmt.Printf("err(%v)\n", err)
	fmt.Println(gold)
}

func Test_Add(t *testing.T) {
	gold, err := packsack_gold.GetGoldMgr().AddGold(12, -1500)
	fmt.Printf("err(%v)\n", err)
	fmt.Println(gold)
}
