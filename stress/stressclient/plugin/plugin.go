
package plugin

import (
	"strings"
	"steve/stress/stressclient/sprite"
	"plugin"
)

func GetPluginSprite(name string) (sprite.Sprite, error) {
	if !strings.HasSuffix(name, ".so") {
		name += ".so"
	}
	p, err := plugin.Open(name)
	if err != nil {
		return nil, err
	}
	f, err := p.Lookup("GetSprite")
	if err != nil {
		return nil, err
	}
	getter := f.(func() sprite.Sprite)
	sp := getter()
	return sp, nil
	//return nil, nil
}
