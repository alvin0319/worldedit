package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

func playerOnly(s cmd.Source) bool {
	_, ok := s.(*player.Player)
	return ok
}
