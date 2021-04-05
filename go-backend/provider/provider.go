package provider

import "context"

// Provider is the the interface that should be implemented
// for each backend provider (i.e. mongodb, etcd, etc)
type Provider interface {
	Init(ctx context.Context) error // initializes the provider
}
