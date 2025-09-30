package slog

import (
	"context"
	"log/slog"
	"runtime/debug"

	"github.com/nasermirzaei89/problem"
)

func New() problem.Logger {
	return func(ctx context.Context, err error) string {
		slog.ErrorContext(ctx, "unexpected error", "error", err, "stack", string(debug.Stack()))

		return ""
	}
}

func NewWithLogger(l *slog.Logger) problem.Logger {
	return func(ctx context.Context, err error) string {
		l.ErrorContext(ctx, "unexpected error", "error", err, "stack", string(debug.Stack()))

		return ""
	}
}
