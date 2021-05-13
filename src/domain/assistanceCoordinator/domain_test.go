package assistanceCoordinator

import (
	"testing"

	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func setupDomain(t *testing.T) (m struct {
	mocks.AppBase
	sut Domain
}) {
	m.AppBase = mocks.SetupAppBase(t)
	m.sut = NewDomain(m.Trace.Tracer, m.Domain.Waiter, m.Domain.CustomerOrder)
	return m
}
