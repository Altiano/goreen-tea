package app

import (
	"testing"

	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func setupApp(t *testing.T) (m struct {
	mocks.AppBase
	sut App
}) {
	m.AppBase = mocks.SetupAppBase(t)
	m.sut = NewApp(
		m.Trace.Tracer,
		m.Domain.CustomerOrder,
		m.Domain.AssistanceCoordinator,
		m.Domain.VisitorCounter)
	return m
}
