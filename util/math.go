package util

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"math"
)

func DistanceSquared(pos1, pos2 cube.Pos) float64 {
	dx := float64(pos1.X() - pos2.X())
	dy := float64(pos1.Y() - pos2.Y())
	dz := float64(pos1.Z() - pos2.Z())
	return (dx * dx) + (dy * dy) + (dz * dz)
}

func Distance(pos1, pos2 cube.Pos) float64 {
	return math.Sqrt(DistanceSquared(pos1, pos2))
}
