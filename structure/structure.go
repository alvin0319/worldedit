package structure

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

// Structure represents a 3D structure in the world, defined by its width, height, and depth.
type Structure struct {
	world.Structure

	Width, Height, Depth int
	// Blocks is a map of positions to blocks that make up the structure.
	// cube.Pos is used as the key to represent the 3D coordinates of each block
	// that represents an offset of the block in the structure.
	Blocks map[cube.Pos]world.Block
}

func New(width, height, depth int) Structure {
	return Structure{
		Width:  width,
		Height: height,
		Depth:  depth,
		Blocks: make(map[cube.Pos]world.Block),
	}
}

func (s Structure) Dimensions() [3]int {
	return [3]int{s.Width, s.Height, s.Depth}
}

func (s Structure) At(x, y, z int, f func(x, y, z int) world.Block) (world.Block, world.Liquid) {
	if b, ok := s.Blocks[cube.Pos{x, y, z}]; ok {
		return b, nil
	}
	return nil, nil
}

func (s Structure) Set(x, y, z int, b world.Block) {
	s.Blocks[cube.Pos{x, y, z}] = b
}
