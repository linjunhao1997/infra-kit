package grpclib

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

var (
	LoggingUnaryClientInterceptor  grpc.UnaryClientInterceptor
	LoggingUnaryStreamInterceptor  grpc.StreamClientInterceptor
	LoggingUnaryServerInterceptor  grpc.UnaryServerInterceptor
	LoggingStreamServerInterceptor grpc.StreamServerInterceptor
)

func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func init() {
	logTraceID := func(ctx context.Context) logging.Fields {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return logging.Fields{"traceID", span.TraceID().String()}
		}
		return nil
	}

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		logging.WithFieldsFromContext(logTraceID),
	}

	LoggingUnaryClientInterceptor = logging.UnaryClientInterceptor(interceptorLogger(slog.Default()), opts...)
	LoggingUnaryStreamInterceptor = logging.StreamClientInterceptor(interceptorLogger(slog.Default()), opts...)
	LoggingUnaryServerInterceptor = logging.UnaryServerInterceptor(interceptorLogger(slog.Default()), opts...)
	LoggingStreamServerInterceptor = logging.StreamServerInterceptor(interceptorLogger(slog.Default()), opts...)

}
