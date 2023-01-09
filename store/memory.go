package store

import (
	"github.com/anthonysyk/fizzbuzz/model"
	"sync/atomic"
)

var _ Store = MemoryStore{}

type MemoryStore map[string]*int64

func NewMemoryStore() MemoryStore {
	return make(map[string]*int64)
}

func (ms MemoryStore) AddQueryHit(key string) error {
	if ms[key] == nil {
		ms[key] = new(int64)
	}
	atomic.AddInt64(ms[key], 1)
	return nil
}

func (ms MemoryStore) GetMostUsedQuery() model.MostUsedQuery {
	topHits := model.MostUsedQuery{}
	for k, v := range ms {
		if *v > topHits.Hits {
			topHits = model.MostUsedQuery{Params: k, Hits: *v}
		}
	}
	return topHits
}
