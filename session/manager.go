package session

import "github.com/df-mc/dragonfly/server/player"

var sessions = make(map[string]*Session)

// GetSession returns the session for the given player
func GetSession(p *player.Player) *Session {
	return sessions[p.XUID()]
}

// CreateSession creates a new session for the given player
func CreateSession(p *player.Player) *Session {
	s := NewSession()
	sessions[p.XUID()] = s
	return s
}

// RemoveSession removes the session for the given player
func RemoveSession(p *player.Player) {
	if s, ok := sessions[p.XUID()]; ok {
		s.Close()
		delete(sessions, p.XUID())
	}
}
