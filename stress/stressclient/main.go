package main

import (
	_ "steve/stress/stressclient/core"
	"steve/stress/stressclient/core"
	"steve/stress/stressclient/plugin"
)

func main() {
	core.SetGetSpriteFunc(plugin.GetPluginSprite)
	select {}
}
