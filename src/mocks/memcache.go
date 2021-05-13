package mocks

import (
	"github.com/golang/mock/gomock"
	memcacheMocks "gitlab.com/altiano/goreen-tea/src/frameworks/memcache/mocks"
)

type (
	Memcache struct {
		Manager *memcacheMocks.MockManager
	}
)

func setupMemcache(ctrl *gomock.Controller) Memcache {
	return Memcache{
		Manager: memcacheMocks.NewMockManager(ctrl),
	}
}
