package kvmem

import "github.com/blevesearch/bleve/v2"

// Bucket a bucket implementation
type Bucket struct {
	Name  string
	Data  map[string]interface{}
	index bleve.Index
}

// NewBucket creates a new bucket
func NewBucket(name string) (*Bucket, error) {
	var err error

	b := &Bucket{
		Name: name,
		Data: map[string]interface{}{},
	}

	mapping := bleve.NewIndexMapping()
	b.index, err = bleve.NewMemOnly(mapping)
	if err != nil {
		return nil, err
	}

	return b, nil
}
