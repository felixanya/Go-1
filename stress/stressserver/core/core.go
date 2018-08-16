package core

import (
	"github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
	"github.com/spf13/viper"
	"fmt"
	"github.com/boltdb/bolt"
)

var Grpc *grpc.Server
var db *bolt.DB

func Init() {
	startHttp()
	initDB()
	Grpc = createRPCServer("", "")
	addr := viper.GetString("rpc_addr")
	port := viper.GetInt("rpc_port")
	runRPCServer(addr, port)
}

func initDB() {
	var err error
	file := fmt.Sprintf("%s/stress.db", viper.GetString("db_path"))
	db, err = bolt.Open(file, 0600, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		logrus.Info("")
		b, err := tx.CreateBucketIfNotExists([]byte("clients"))
		if err != nil {
			return err
		}
		b.Put([]byte("dd"), []byte("dddddd"))
		b.NextSequence()
		return nil
	})
}
