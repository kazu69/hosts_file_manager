package hfm

import (
	"os"
)

// Check hosts file writable
func (h *Hosts) IsWritable() bool {
	if h.path == "" {
		panic("Error: path is Empty")
	}

	_, err := os.OpenFile(h.path, os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return false
	}
	return true
}
