package provider

import "context"

// EntityMarshaller converts a raw entity to an expected type
type EntityMarshaller func(key string, in interface{}) (out interface{}, err error)

// KV key-value provider interface type. Should implement
// the basic provider interface along with key-value specific methods
type KV interface {
	Init(ctx context.Context) error // initializes the provider
	Get(ctx context.Context, bucket, key string, value interface{}) (exists bool, err error)
	Set(ctx context.Context, bucket, key string, newValue, oldValue interface{}) (updated bool, err error)
	Del(ctx context.Context, bucket, key string, oldValue interface{}) (deleted bool, err error)
	Keys(ctx context.Context, bucket string) (keys []string, err error)
	Find(ctx context.Context, bucket, query string, marshaller EntityMarshaller) (rsp interface{}, err error)
	// Index(ctx context.Context, bucket string, fields []string) (err error)
	// Unique(ctx context.Context, bucket string, fields []string) (err error)
}
