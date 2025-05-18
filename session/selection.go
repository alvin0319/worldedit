package session

import (
	"github.com/df-mc/dragonfly/server/block/cube"
)

type Selection struct {
	Pos1 cube.Pos
	Pos2 cube.Pos

	Pos1Set bool
	Pos2Set bool

	WorldName string
}

func (s *Selection) Ready() bool {
	return s.Pos1Set && s.Pos2Set
}

func (s *Selection) Min() (int, int, int) {
	return min(s.Pos1.X(), s.Pos2.X()), min(s.Pos1.Y(), s.Pos2.Y()), min(s.Pos1.Z(), s.Pos2.Z())
}

func (s *Selection) Max() (int, int, int) {
	return max(s.Pos1.X(), s.Pos2.X()), max(s.Pos1.Y(), s.Pos2.Y()), max(s.Pos1.Z(), s.Pos2.Z())
}
