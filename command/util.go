package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

func playerOnly(s cmd.Source) bool {
	_, ok := s.(*player.Player)
	return ok
}

var allower = func(s cmd.Source) bool {
	return playerOnly(s)
}

func SetAllower(f func(s cmd.Source) bool) {
	allower = f
}
