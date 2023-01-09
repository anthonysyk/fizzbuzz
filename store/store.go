package store

import (
	"github.com/anthonysyk/fizzbuzz/model"
)

type Store interface {
	GetMostUsedQuery() model.MostUsedQuery
	AddQueryHit(key string) error
}
