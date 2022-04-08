package output

import (
	"fmt"
	"log"

	"github.com/nasermirzaei89/problem"
)

func New() problem.Logger {
	return func(err error) string {
		log.Println(fmt.Sprintf("%+v", err))

		return ""
	}
}

func NewWithLogger(l *log.Logger) problem.Logger {
	return func(err error) string {
		l.Println(fmt.Sprintf("%+v", err))

		return ""
	}
}
