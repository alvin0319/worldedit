package worldedit

import (
	item2 "github.com/alvin0319/worldedit/item"
	"github.com/alvin0319/worldedit/session"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
)

type Handler struct {
	player.NopHandler
	s *session.Session
}

func NewHandler(s *session.Session) Handler {
	return Handler{
		s: s,
	}
}

func (h Handler) HandleStartBreak(ctx *player.Context, pos cube.Pos) {
	p := ctx.Val()
	s, _ := p.HeldItems()
	if s.Empty() {
		return
	}
	if _, ok := s.Item().(item2.Wand); !ok {
		return
	}
	ctx.Cancel()
	h.s.SelectPos1(p, pos)
}

func (h Handler) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, _ cube.Face, _ mgl64.Vec3) {
	p := ctx.Val()
	s, _ := p.HeldItems()
	if s.Empty() {
		return
	}
	if _, ok := s.Item().(item2.Wand); !ok {
		return
	}
	ctx.Cancel()
	h.s.SelectPos2(p, pos)
}

func (h Handler) HandleQuit(p *player.Player) {
	session.RemoveSession(p)
}
