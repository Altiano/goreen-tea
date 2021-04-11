package trace

import (
	"context"
)

type (
	Tracer interface {
		Start(ctx context.Context, spanName string) (context.Context, Span)
	}

	Span interface {
		End()
	}
)
