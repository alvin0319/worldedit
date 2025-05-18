package job

import (
	"github.com/alvin0319/worldedit/session"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

// Fill is a job that simply fills given area
type Fill struct {
	Simple
}

func NewFill(b *world.Block) *Fill {
	return &Fill{Simple{*b}}
}

func (f Fill) Run(tx *world.Tx, h *session.History, s session.Selection) int {
	var changed = 0
	x1, y1, z1 := s.Min()
	x2, y2, z2 := s.Max()

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			for z := z1; z <= z2; z++ {
				h.SetBlock(cube.Pos{int(x), int(y), int(z)}, f.Block, tx)
				changed++
			}
		}
	}
	return changed
}
