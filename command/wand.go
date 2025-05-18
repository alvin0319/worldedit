package command

import (
	"github.com/alvin0319/worldedit/item"
	"github.com/df-mc/dragonfly/server/cmd"
	item2 "github.com/df-mc/dragonfly/server/item"
	player2 "github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Wand struct {
}

func (Wand) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	if _, ok := source.(*player2.Player); !ok {
		output.Error("This command can only be run by a player.")
		return
	}
	player := source.(*player2.Player)
	_, _ = player.Inventory().AddItem(item2.NewStack(&item.Wand{}, 1))
	output.Print("Gave you a wand.")
}

func (Wand) Allow(s cmd.Source) bool {
	return playerOnly(s)
}
