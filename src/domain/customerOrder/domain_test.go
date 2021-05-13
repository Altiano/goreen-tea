package customerOrder

import (
	"testing"

	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func setupDomain(t *testing.T) (m struct {
	mocks.DomainBase
	sut Domain
}) {
	m.DomainBase = mocks.SetupDomainBase(t)
	m.sut = NewDomain(m.Trace.Tracer, m.Email.Manager, m.Repo.CustomerOrder)
	return m
}
