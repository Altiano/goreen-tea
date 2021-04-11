package assistanceCoordinator

import (
	"testing"

	"gitlab.com/altiano/golang-boilerplate/src/mocks"
)

func setupDomain(t *testing.T) (m struct {
	mocks.AppBase
	sut Domain
}) {
	m.AppBase = mocks.SetupAppBase(t)
	m.sut = NewDomain(m.Trace.Tracer, m.Domain.Waiter, m.Domain.CustomerOrder)
	return m
}
