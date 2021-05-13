package memcache

import (
	"gitlab.com/altiano/goreen-tea/src/shared"
)

type (
	dummy struct {
		config      shared.Config
		incrementor map[string]int
	}
)

func NewDummy(config shared.Config) Memcacher {
	return &dummy{
		config:      config,
		incrementor: map[string]int{},
	}
}

func (m *dummy) Increment(prefix string) (int, error) {
	m.incrementor[prefix]++
	return m.incrementor[prefix], nil
}
