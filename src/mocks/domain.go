package mocks

import (
	"testing"

	"github.com/golang/mock/gomock"
	assistanceCoordinatorMocks "gitlab.com/altiano/golang-boilerplate/src/domain/assistanceCoordinator/mocks"
	coMocks "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/mocks"
	visitorCounterMocks "gitlab.com/altiano/golang-boilerplate/src/domain/visitorCounter/mocks"
	waiterMocks "gitlab.com/altiano/golang-boilerplate/src/domain/waiter/mocks"
)

type (
	DomainBase struct {
		Base
		Database
		Trace
		Email
		Repo
		Memcache
	}

	Domain struct {
		Waiter                *waiterMocks.MockDomain
		CustomerOrder         *coMocks.MockDomain
		AssistanceCoordinator *assistanceCoordinatorMocks.MockDomain
		VisitorCounter        *visitorCounterMocks.MockDomain
	}
)

func SetupDomainBase(t *testing.T) DomainBase {
	ctrl := gomock.NewController(t)

	return DomainBase{
		Base:     setupBase(),
		Database: setupDatabase(ctrl),
		Trace:    setupTrace(ctrl),
		Email:    setupEmail(ctrl),
		Repo:     setupRepo(ctrl),
		Memcache: setupMemcache(ctrl),
	}
}

func setupDomain(ctrl *gomock.Controller) Domain {
	return Domain{
		Waiter:                waiterMocks.NewMockDomain(ctrl),
		CustomerOrder:         coMocks.NewMockDomain(ctrl),
		AssistanceCoordinator: assistanceCoordinatorMocks.NewMockDomain(ctrl),
		VisitorCounter:        visitorCounterMocks.NewMockDomain(ctrl),
	}
}
