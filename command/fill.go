package command

import (
	"github.com/alvin0319/worldedit/job"
	session2 "github.com/alvin0319/worldedit/session"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"time"
)

type Fill struct {
}

func (f Fill) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	if _, ok := source.(*player.Player); !ok {
		output.Error("This command can only be run by a player.")
		return
	}
	p := source.(*player.Player)
	session := session2.GetSession(p)
	if session == nil {
		output.Error("Session not found.")
		return
	}
	if !session.SelectionReady() {
		output.Error("You have to select a region first.")
		return
	}
	if session.Selection().WorldName != tx.World().Name() {
		output.Error("You can only fill blocks in the world you selected.")
		return
	}

	heldItem, _ := p.HeldItems()
	if _, ok := heldItem.Item().(world.Block); !ok {
		output.Error("You must hold a block in your hand to fill with.")
		return
	}

	b := heldItem.Item().(world.Block)

	fill := job.NewFill(&b)
	t := time.Now()
	changed := job.PreRun(fill, tx, session)
	tx.PlaySound(p.Position(), sound.BlockPlace{Block: b})
	output.Printf("Filled %d blocks (took %.1f seconds)", changed, time.Since(t).Seconds())
}

func (Fill) Allow(s cmd.Source) bool {
	return allower(s)
}
