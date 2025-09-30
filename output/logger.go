package output

import (
	"context"
	"log"

	"github.com/nasermirzaei89/problem"
)

func New() problem.Logger {
	return func(_ context.Context, err error) string {
		log.Printf("%+v\n", err)

		return ""
	}
}

func NewWithLogger(l *log.Logger) problem.Logger {
	return func(_ context.Context, err error) string {
		l.Printf("%+v\n", err)

		return ""
	}
}
