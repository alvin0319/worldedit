package command

import (
	"github.com/alvin0319/worldedit/session"
	"github.com/alvin0319/worldedit/structure"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Copy struct {
}

func (Copy) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
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
	if !s.SelectionReady() {
		output.Error("You must select a region before copying.")
		return
	}
	minX, minY, minZ := s.Selection().Min()
	maxX, maxY, maxZ := s.Selection().Max()
	width, height, depth := maxX-minX+1, maxY-minY+1, maxZ-minZ+1
	str := structure.New(width, height, depth)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				offsetX := x - minX
				offsetY := y - minY
				offsetZ := z - minZ
				block := tx.Block(cube.Pos{x, y, z})
				str.Set(offsetX, offsetY, offsetZ, block)
			}
		}
	}
	s.SetStructure(str)
	output.Printf("Copied structure of size %dx%dx%d.", width, height, depth)
}

func (Copy) Allow(s cmd.Source) bool {
	return playerOnly(s)
}
