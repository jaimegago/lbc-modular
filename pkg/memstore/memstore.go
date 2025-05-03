package memstore

import (
	fb "github.com/jaimegago/lbc-modular/pkg/fizzbuzz"
)

type MemStore struct {
	stats map[*fb.ReqData]int
}

func NewMemStore() *MemStore {
	return &MemStore{stats: map[*fb.ReqData]int{}}
}

func (m MemStore) Update(*fb.ReqData) {

}
