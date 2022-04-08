package output

import (
	"fmt"
	"github.com/nasermirzaei89/problem"
	"log"
)

func New() problem.Logger {
	return func(err error) string {
		log.Println(fmt.Sprintf("%+v", err))

		return ""
	}
}
