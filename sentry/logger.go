package sentry

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/nasermirzaei89/problem"
)

func New() problem.Logger {
	return func(ctx context.Context, err error) string {
		res := sentry.CaptureException(err)
		if res != nil {
			return string(*res)
		}

		return ""
	}
}

func NewWithClient(client *sentry.Client) problem.Logger {
	return func(ctx context.Context, err error) string {
		res := client.CaptureException(err, nil, nil)
		if res != nil {
			return string(*res)
		}

		return ""
	}
}
