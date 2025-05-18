package session

import (
	"errors"
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

type BlockPosHash uint64

// History is a struct that stores the history of blocks in the world
type History struct {
	// Blocks is a map of block positions to blocks
	Blocks map[cube.Pos]world.Block

	lastHash BlockPosHash

	checkDuplicate bool

	WorldName string
}

func NewHistory(worldName string, checkDuplicate bool) *History {
	return &History{
		checkDuplicate: checkDuplicate,
		WorldName:      worldName,
		Blocks:         make(map[cube.Pos]world.Block),
	}
}

// AddBlock adds a block to the history
func (h *History) AddBlock(pos cube.Pos, block world.Block) error {
	if h.checkDuplicate {
		if _, ok := h.Blocks[pos]; ok {
			return errors.New(fmt.Sprintf("duplicate block for pos %v", pos))
		}
	}
	h.Blocks[pos] = block
	return nil
}

// RemoveBlock removes the block at the given position
func (h *History) RemoveBlock(pos cube.Pos) {
	if _, ok := h.Blocks[pos]; ok {
		delete(h.Blocks, pos)
	}
}

// SetBlock records the block at the given position and sets it in the world
func (h *History) SetBlock(pos cube.Pos, b world.Block, tx *world.Tx) {
	_ = h.AddBlock(pos, tx.Block(pos))
	tx.SetBlock(pos, b, nil)
}

// Block returns the block at the given position
func (h *History) Block(pos cube.Pos) world.Block {
	block, ok := h.Blocks[pos]
	if !ok {
		return nil
	}
	return block
}

func (h *History) LastHash() BlockPosHash {
	return h.lastHash
}

// Revert reverts the last change made to the world
func (h *History) Revert(tx *world.Tx) (int, error) {
	if len(h.Blocks) == 0 {
		return -1, errors.New("no blocks to revert")
	}
	for pos, block := range h.Blocks {
		tx.SetBlock(pos, block, nil)
	}
	changed := len(h.Blocks)
	h.Blocks = make(map[cube.Pos]world.Block)
	return changed, nil
}
