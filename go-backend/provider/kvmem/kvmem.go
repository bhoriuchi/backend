package kvmem

import (
	"context"
	"log"

	"github.com/bhoriuchi/backend/go-backend/provider"
	"github.com/bhoriuchi/backend/go-backend/utils"
	"github.com/blevesearch/bleve/v2"
)

// Provider a memory based kv provider
type Provider struct {
	buckets map[string]*Bucket
}

// New creates a new kvmem provider
func New() *Provider {
	return &Provider{}
}

// creates a bucket if it does not exist and returns it
func (p *Provider) ensureBucket(bucket string) (*Bucket, error) {
	var err error

	if p.buckets == nil {
		p.buckets = map[string]*Bucket{}
	}

	b, ok := p.buckets[bucket]
	if !ok {
		if b, err = NewBucket(bucket); err != nil {
			return nil, err
		}

		p.buckets[bucket] = b
	}

	return b, nil
}

// Init initialize the provider
func (p *Provider) Init(ctx context.Context) error {
	return nil
}

// Get an entity
func (p *Provider) Get(ctx context.Context, bucket, key string, value interface{}) (bool, error) {
	b, err := p.ensureBucket(bucket)
	if err != nil {
		return false, err
	}

	entry, exists := b.Data[key]
	if value != nil && entry != nil {
		err = utils.MapStructure(entry, value)
	}

	return exists, err
}

// Set sets an entry
func (p *Provider) Set(ctx context.Context, bucket, key string, newValue, oldValue interface{}) (bool, error) {
	b, err := p.ensureBucket(bucket)
	if err != nil {
		return false, err
	}

	exists, err := p.Get(ctx, bucket, key, oldValue)
	if err != nil {
		return false, err
	}

	var entry interface{}
	if err := utils.MapStructure(newValue, &entry); err != nil {
		return false, err
	}

	b.Data[key] = entry
	if err := b.index.Index(key, entry); err != nil {
		log.Println("index error:", err)
	}

	return exists, nil
}

// Del deletes an entry
func (p *Provider) Del(ctx context.Context, bucket, key string, oldValue interface{}) (deleted bool, err error) {
	b, err := p.ensureBucket(bucket)
	if err != nil {
		return false, err
	}

	exists, err := p.Get(ctx, bucket, key, oldValue)
	if err != nil {
		return false, err
	}

	delete(b.Data, key)
	if err := b.index.Delete(key); err != nil {
		log.Println("index error:", err)
	}

	return exists, nil
}

// Keys list keys
func (p *Provider) Keys(ctx context.Context, bucket string) ([]string, error) {
	b, err := p.ensureBucket(bucket)
	if err != nil {
		return nil, err
	}

	keys := []string{}
	for key := range b.Data {
		keys = append(keys, key)
	}

	return keys, nil
}

// Find
func (p *Provider) Find(ctx context.Context, bucket, query string, marshaller provider.EntityMarshaller) (interface{}, error) {
	b, err := p.ensureBucket(bucket)
	if err != nil {
		return nil, err
	}

	q := bleve.NewMatchQuery(query)
	search := bleve.NewSearchRequest(q)
	results, err := b.index.Search(search)
	if err != nil {
		return nil, err
	}

	rsp := map[string]interface{}{}
	for _, result := range results.Hits {
		var entry interface{}
		found, err := p.Get(ctx, bucket, result.ID, &entry)
		if err != nil {
			return nil, err
		}
		if found {
			rsp[result.ID], err = marshaller(result.ID, entry)
			if err != nil {
				return nil, err
			}
		}
	}

	return rsp, nil
}
