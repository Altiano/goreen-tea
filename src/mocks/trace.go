package mocks

import (
	"context"

	"github.com/golang/mock/gomock"
	traceMocks "gitlab.com/altiano/goreen-tea/src/frameworks/trace/mocks"
)

type (
	Trace struct {
		Tracer *traceMocks.MockTracer
		Span   *traceMocks.MockSpan
	}
)

func setupTrace(ctrl *gomock.Controller) Trace {
	return Trace{
		Tracer: traceMocks.NewMockTracer(ctrl),
		Span:   traceMocks.NewMockSpan(ctrl),
	}
}

func (t Trace) ExpectSpan(ctx context.Context, name string) {
	t.Tracer.EXPECT().Start(ctx, name).Return(ctx, t.Span)
	t.Span.EXPECT().End().Return()
}
