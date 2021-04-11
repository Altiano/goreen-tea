package mocks

import (
	"testing"

	"github.com/golang/mock/gomock"
)

type (
	AppBase struct {
		Base
		Database
		Trace
		Memcache
		Domain
	}
)

func SetupAppBase(t *testing.T) AppBase {
	ctrl := gomock.NewController(t)

	return AppBase{
		Base:     setupBase(),
		Database: setupDatabase(ctrl),
		Trace:    setupTrace(ctrl),
		Memcache: setupMemcache(ctrl),
		Domain:   setupDomain(ctrl),
	}
}
