package job

import (
	"github.com/alvin0319/worldedit/session"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"reflect"
)

// Replace searches From block from area and replaces them to To
type Replace struct {
	// From is a block to be replaced
	From world.Block
	// To is a block to replace
	To world.Block
}

func NewReplace(from world.Block, to world.Block) *Replace {
	return &Replace{
		From: from,
		To:   to,
	}
}

func (r Replace) Run(tx *world.Tx, h *session.History, s session.Selection) int {
	var changed = 0
	x1, y1, z1 := s.Min()
	x2, y2, z2 := s.Max()
	wantedName, wantedProp := r.From.EncodeBlock()
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			for z := z1; z <= z2; z++ {
				pos := cube.Pos{int(x), int(y), int(z)}
				name, prop := tx.Block(pos).EncodeBlock()
				if name == wantedName && reflect.DeepEqual(prop, wantedProp) {
					h.SetBlock(pos, r.To, tx)
					changed++
				}
			}
		}
	}
	return changed
}
