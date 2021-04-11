package trace

import (
	"context"
	"log"
	"sync"
	"time"

	"gitlab.com/altiano/golang-boilerplate/src/shared"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	coretrace "go.opentelemetry.io/otel/trace"
)

type (
	otelTracer struct {
		config shared.Config
		tp     *sdktrace.TracerProvider
		tracer coretrace.Tracer
	}

	otelSpan struct {
		span coretrace.Span
	}
)

var (
	manager *otelTracer
	once    sync.Once
)

func NewOtelManager(config shared.Config) Tracer {
	once.Do(func() {
		if config.JeagerUrl == "" {
			manager = &otelTracer{}
			return
		}

		// Create exporter
		jaegerExp, err := jaeger.NewRawExporter(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.JeagerUrl)))
		if err != nil {
			log.Fatalf("failed to initialize jaeger export pipeline: %v", err)
		}

		// Create tracer provider
		tp := sdktrace.NewTracerProvider(
			sdktrace.WithSpanProcessor(
				sdktrace.NewBatchSpanProcessor(jaegerExp),
			),
			sdktrace.WithResource(resource.NewWithAttributes(
				semconv.ServiceNameKey.String(shared.ServiceName),
			)),
		)

		// Setting global tracer provider
		otel.SetTracerProvider(tp)

		// Create service's main tracer
		t := tp.Tracer(shared.ServiceName)

		//
		manager = &otelTracer{
			config: config,
			tp:     tp,
			tracer: t,
		}
	})

	//
	return manager
}

func Close() {
	if manager != nil && manager.tp != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := manager.tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}
}

/*
	otelTracer implementation
*/
func (m *otelTracer) Start(ctx context.Context, spanName string) (context.Context, Span) {
	if m.tracer == nil {
		return ctx, &otelSpan{}
	}

	ctx, span := m.tracer.Start(ctx, spanName)

	return ctx, &otelSpan{
		span: span,
	}
}

/*
	otelSpan implementation
*/
func (m *otelSpan) End() {
	if m.span == nil {
		return
	}

	m.span.End()
}
