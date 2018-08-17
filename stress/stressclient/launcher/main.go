package main

import (
	"errors"
	"steve/stress/stressclient/core"
	_ "steve/stress/stressclient/core"

	"steve/stress/sprites/login"
	"steve/stress/common"
	"github.com/Sirupsen/logrus"
)

func main() {
	core.SetGetSpriteFunc(getSpriteByName)
	<-common.Waitc
	logrus.Info("EXIT")
}

func getSpriteByName(name string) (core.Sprite, error) {
	switch name {
	case "login":
		return login.GetSprite(), nil
	}
	return nil, errors.New("no sprite")
}
