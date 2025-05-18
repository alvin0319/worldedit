package command

import (
	"github.com/alvin0319/worldedit/session"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Replace struct {
	From cmd.Varargs `cmd:"from"`
	To   cmd.Varargs `cmd:"to"`
}

func (r Replace) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	if _, ok := source.(*player.Player); !ok {
		output.Error("This command can only be run by a player.")
		return
	}
	p := source.(*player.Player)
	s := session.GetSession(p)
	if s == nil {
		output.Error("Session not found.")
		return
	}
	// TODO...
}
