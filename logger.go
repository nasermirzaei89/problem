package problem

import "context"

type Logger func(ctx context.Context, err error) (trackingCode string)

func NewVoidLogger() Logger {
	return func(_ context.Context, _ error) string {
		return ""
	}
}
