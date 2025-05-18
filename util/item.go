package util

import (
	"github.com/alvin0319/worldedit/item"
	"github.com/df-mc/dragonfly/server/world"
)

func InitCustomItem() {
	world.RegisterItem(item.Wand{})
}
