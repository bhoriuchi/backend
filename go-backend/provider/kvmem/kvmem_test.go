package kvmem

import (
	"context"
	"log"
	"testing"

	"github.com/bhoriuchi/backend/go-backend/utils"
)

const bucket = "test_index"

func TestOperations(t *testing.T) {
	testData := map[string]map[string]string{
		"foo": {
			"test": "foo",
		},
		"bar": {
			"test": "bar",
		},
	}

	p := New()
	p.Init(context.Background())

	// add a record
	for k, v := range testData {
		updated, err := p.Set(context.Background(), bucket, k, v, nil)

		if err != nil {
			t.Error(err)
			return
		} else if updated {
			t.Errorf("should not return updated")
			return
		}
	}

	// fget a record
	var foo interface{}
	found, err := p.Get(context.Background(), bucket, "foo", &foo)
	if err != nil {
		t.Error(err)
		return
	} else if !found {
		t.Errorf("record foo not found")
		return
	}

	res, err := p.Find(context.Background(), bucket, "test=bar", func(key string, in interface{}) (interface{}, error) {
		m := map[string]string{}
		if err := utils.MapStructure(in, &m); err != nil {
			return nil, err
		}
		return m, nil
	})

	if err != nil {
		t.Error(err)
		return
	}

	log.Println(res)
}
