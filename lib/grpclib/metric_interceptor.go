package grpclib

import (
	"context"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

var (
	MetricUnaryClientInterceptor  grpc.UnaryClientInterceptor
	MetricStreamClientInterceptor grpc.StreamClientInterceptor

	MetricUnaryServerInterceptor  grpc.UnaryServerInterceptor
	MetricStreamServerInterceptor grpc.StreamServerInterceptor
)

func init() {
	clMetrics := grpcprom.NewClientMetrics(
		grpcprom.WithClientHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)

	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics, clMetrics)
	exemplarFromContext := func(ctx context.Context) prometheus.Labels {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return prometheus.Labels{"traceID": span.TraceID().String()}
		}
		return nil
	}

	MetricUnaryClientInterceptor = clMetrics.UnaryClientInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext))
	MetricStreamClientInterceptor = clMetrics.StreamClientInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext))
	MetricUnaryServerInterceptor = srvMetrics.UnaryServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext))
	MetricStreamServerInterceptor = srvMetrics.StreamServerInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext))
}
