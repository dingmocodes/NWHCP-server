package handlers

import (
	"NWHCP/NWHCP-server/gateway/models/users"
	"NWHCP/NWHCP-server/gateway/sessions"
)

// Handler blah
type Handler struct {
	SessionKey   string         `json:"key"`
	SessionStore sessions.Store `json:"sessions"`
	UserStore    users.Store    `json:"users"`
}
