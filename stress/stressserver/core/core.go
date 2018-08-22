package core

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/coreos/bbolt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
