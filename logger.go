package problem

type Logger func(err error) (trackingCode string)

func NewVoidLogger() Logger {
	return func(err error) string {
		return ""
	}
}
