package memstore

import (
	fb "github.com/jaimegago/lbc-modular/pkg/fizzbuzz"
)

type MemStore struct {
	stats []fb.ReqData
}

func New() *MemStore {
	return &MemStore{stats: []fb.ReqData{}}
}

func (m *MemStore) UpdateStats(stats []fb.ReqData) error {
	m.stats = stats
	return nil
}

func (m MemStore) GetStats() ([]fb.ReqData, error) {
	return m.stats, nil
}
