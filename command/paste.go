package command

import (
	"github.com/alvin0319/worldedit/job"
	"github.com/alvin0319/worldedit/session"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"time"
)

type Paste struct {
}

func (Paste) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
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
	if !s.StructureReady() {
		output.Error("You must copy an area before you can paste it.")
		return
	}
	str := s.Structure()
	t := time.Now()
	paste := job.NewPaste(str, cube.PosFromVec3(p.Position()))
	job.PreRun(paste, tx, s)
	output.Printf("Paste successful (took %.1f seconds)", time.Since(t).Seconds())
}

func (Paste) Allow(s cmd.Source) bool {
	return allower(s)
}
