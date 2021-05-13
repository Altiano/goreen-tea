package mocks

import (
	"testing"

	"github.com/golang/mock/gomock"
	coMocks "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/mocks"
	waiterMocks "gitlab.com/altiano/goreen-tea/src/domain/waiter/mocks"
)

type (
	RepoBase struct {
		Base
		Database
		Trace
	}

	Repo struct {
		CustomerOrder *coMocks.MockRepo
		Waiter        *waiterMocks.MockRepo
	}
)

func SetupRepoBase(t *testing.T) RepoBase {
	ctrl := gomock.NewController(t)

	return RepoBase{
		Base:     setupBase(),
		Database: setupDatabase(ctrl),
		Trace:    setupTrace(ctrl),
	}
}

func setupRepo(ctrl *gomock.Controller) Repo {
	return Repo{
		CustomerOrder: coMocks.NewMockRepo(ctrl),
		Waiter:        waiterMocks.NewMockRepo(ctrl),
	}
}
