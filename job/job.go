package job

import (
	"github.com/alvin0319/worldedit/session"
	"github.com/df-mc/dragonfly/server/world"
)

// Job is an interface that all operation should implement
type Job interface {
	// Run runs a job
	Run(tx *world.Tx, h *session.History, s session.Selection) int
}

func PreRun(j Job, tx *world.Tx, s *session.Session) int {
	h := session.NewHistory(tx.World().Name(), true)
	changed := j.Run(tx, h, *s.Selection())
	s.CommitHistory(h)
	return changed
}

// Simple is a struct that represents simple operation that needs only one block
type Simple struct {
	Block world.Block
}
