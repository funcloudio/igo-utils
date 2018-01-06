// Package req has helpers for parsing a http Request
package req

import (
	"encoding/json"
	"net/http"
)

type Args map[string]string
type Path map[string]string

type Request struct {
	http.Request
	Args `json:"args"`
	Path `json:"path"`
}

// Decode Unmarshal a Request json
func Decode(s string) (r Request, err error) {
	err = json.Unmarshal([]byte(s), &r)
	return
}

// Get return a value in a Path map by a key
func (p Path) Get(key string) string {
	if v, ok := p[key]; ok {
		return v
	}
	return ""
}

// Get return a value in a Args map by a key
func (ar Args) Get(key string) string {
	if v, ok := ar[key]; ok {
		return v
	}
	return ""
}
