package pages

import (
	"github.com/gorilla/sessions"
)

// Store is
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	Store = sessions.NewCookieStore(key)
)

var (
	getDeviceID string
)

type requestJSON struct {
	RequestPage     string      `json:"request_page"`
	SessionName     string      `json:"session_name"`
	SessionUsername interface{} `json:"session_username"`
	DeviceID        string      `json:"device_id"`
	Request         string      `json:"request"`
}
