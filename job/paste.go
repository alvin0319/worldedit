package job

import (
	"github.com/alvin0319/worldedit/session"
	"github.com/alvin0319/worldedit/structure"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

// Paste is a job that pastes a structure at a given position in the world.
type Paste struct {
	s structure.Structure
	// Pos is the current position of the player.
	Pos cube.Pos

	// TODO: respect player's rotation
}

func NewPaste(s structure.Structure, pos cube.Pos) *Paste {
	return &Paste{
		s:   s,
		Pos: pos,
	}
}

func (p *Paste) Run(tx *world.Tx, h *session.History, s session.Selection) int {
	// assume Pos1 is a current position
	tx.BuildStructure(p.Pos, p.s)
	return 0 // count of blocked changed are unknown due to external library
}
