package visitorCounter

import (
	"context"

	"gitlab.com/altiano/goreen-tea/src/frameworks/memcache"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
)

type (
	Domain interface {
		IncreaseVisits(ctx context.Context) error
	}

	domain struct {
		Tracer   trace.Tracer
		Memcache memcache.Memcacher
	}
)

func NewDomain(tracer trace.Tracer, memcache memcache.Memcacher) Domain {
	return domain{
		Tracer:   tracer,
		Memcache: memcache,
	}
}
