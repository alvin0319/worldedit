package session

import (
	"github.com/alvin0319/worldedit/structure"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"sync"
	"time"
)

type Session struct {
	mu sync.Mutex

	histories []*History

	s *Selection

	str      structure.Structure
	strReady bool
}

func NewSession() *Session {
	return &Session{}
}

// CommitHistory stores a history of the operations performed by the player
func (s *Session) CommitHistory(h *History) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.histories = append(s.histories, h)
}

// GetHistories returns the history of the operations performed by the player
func (s *Session) GetHistories() []*History {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.histories
}

func (s *Session) SelectPos1(p *player.Player, pos cube.Pos) {
	s.mu.Lock()

	if s.s == nil {
		s.s = &Selection{
			WorldName: p.Tx().World().Name(),
		}
	}
	s.s.Pos1 = pos
	s.s.Pos1Set = true
	p.Message(text.Colourf("<green>Pos 1 selected at %s</green>", pos.String()))
	s.mu.Unlock()
	s.notifyReady(p)
}

func (s *Session) SelectPos2(p *player.Player, pos cube.Pos) {
	s.mu.Lock()

	if s.s == nil {
		s.s = &Selection{
			WorldName: p.Tx().World().Name(),
		}
	}
	s.s.Pos2 = pos
	s.s.Pos2Set = true
	p.Message(text.Colourf("<green>Pos 2 selected at %s</green>", pos.String()))
	s.mu.Unlock()
	s.notifyReady(p)
}

func (s *Session) SelectionReady() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.s == nil {
		return false
	}
	return s.s.Ready()
}

func (s *Session) Selection() *Selection {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.s == nil {
		return nil
	}
	return s.s
}

func (s *Session) notifyReady(p *player.Player) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.s != nil && s.s.Ready() {
		x1, y1, z1 := s.s.Min()
		x2, y2, z2 := s.s.Max()
		nBlocks := (x2 - x1 + 1) * (y2 - y1 + 1) * (z2 - z1 + 1)
		p.Message(text.Colourf("<green>%d blocks selected</green>", nBlocks))
	}
}

func (s *Session) SetStructure(str structure.Structure) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.str = str
	s.strReady = true
}

func (s *Session) Structure() structure.Structure {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.strReady {
		panic("Structure not ready")
	}
	return s.str
}

func (s *Session) StructureReady() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.strReady
}

func (s *Session) Undo(p *player.Player, tx *world.Tx) {
	s.mu.Lock()
	defer s.mu.Unlock()

	h := s.histories[len(s.histories)-1]
	if h == nil {
		p.Message(text.Colourf("<red>No history to undo</red>"))
		return
	}
	w := tx.World()
	if w.Name() != h.WorldName {
		p.Message(text.Colourf("<red>You are not on the same world</red>"))
		return
	}
	t := time.Now()
	changed, err := h.Revert(tx)
	if err != nil {
		p.Message(text.Colourf("<red>Error reverting history: %v</red>", err))
		return
	}
	p.Message(text.Colourf("<green>%d blocks reverted (took %.1f seconds)</green>", changed, time.Since(t).Seconds()))
	s.histories = s.histories[:len(s.histories)-1]
}

func (s *Session) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.histories = nil
}
