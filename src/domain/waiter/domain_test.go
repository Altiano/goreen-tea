package waiter

import (
	"testing"

	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func setupDomain(t *testing.T) (m struct {
	mocks.DomainBase
	sut Domain
}) {
	m.DomainBase = mocks.SetupDomainBase(t)
	m.sut = NewDomain(m.Trace.Tracer, m.Repo.Waiter)
	return m
}
