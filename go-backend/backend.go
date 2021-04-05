package backend

import "github.com/bhoriuchi/backend/go-backend/provider"

// Backend a backend
type Backend struct {
	provider provider.Provider
}

// NewBackend creates a new backend
func NewBackend(p provider.Provider) (*Backend, error) {
	b := &Backend{provider: p}
	return b, nil
}
