package main

import (
	"errors"
	"steve/stress/stressclient/core"
	_ "steve/stress/stressclient/core"
	"steve/stress/stressclient/sprite"

	"steve/stress/sprites/login"
)

func main() {
	core.SetGetSpriteFunc(getSpriteByName)
	select {}
}

func getSpriteByName(name string) (sprite.Sprite, error) {
	switch name {
	case "login":
		return login.GetSprite(), nil
	}
	return nil, errors.New("no sprite")
}
