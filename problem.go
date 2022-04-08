package problem

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Problem struct {
	Type       string
	Title      string
	Status     int
	Detail     string
	Instance   string
	Extensions map[string]interface{}
}

func (p Problem) MarshalJSON() ([]byte, error) {
	c := make(map[string]interface{})

	c["type"] = "about:blank"
	if p.Type != "" {
		c["type"] = p.Type
	}

	c["status"] = http.StatusInternalServerError
	if p.Status != 0 {
		c["status"] = p.Status
	}

	c["title"] = http.StatusText(c["status"].(int))
	if p.Title != "" {
		c["title"] = p.Title
	}

	c["detail"] = p.Detail

	if p.Instance != "" {
		c["instance"] = p.Instance
	}

	for k, v := range p.Extensions {
		switch k {
		case "type", "status", "title", "detail", "instance":
			c["_"+k] = v
		default:
			c[k] = v
		}
	}

	return json.Marshal(c)
}

func (p Problem) StatusCode() int {
	if p.Status == 0 {
		return http.StatusInternalServerError
	}

	return p.Status
}

func (p Problem) Header() http.Header {
	res := make(http.Header)
	res.Set("Content-Type", "application/problem+json")

	return res
}

type Option func(e *Problem)

func WithType(typ string) Option {
	return func(e *Problem) {
		e.Type = typ
	}
}

func WithTitle(title string) Option {
	return func(e *Problem) {
		e.Title = title
	}
}

func WithDetail(detail string) Option {
	return func(e *Problem) {
		e.Detail = detail
	}
}

func WithStatus(status int) Option {
	return func(e *Problem) {
		e.Status = status
	}
}

func WithExtension(key string, val interface{}) Option {
	return func(e *Problem) {
		if e.Extensions == nil {
			e.Extensions = make(map[string]interface{})
		}

		e.Extensions[key] = val
	}
}

func CustomError(options ...Option) Problem {
	e := Problem{}

	for i := range options {
		options[i](&e)
	}

	return e
}

var logger Logger = NewVoidLogger()

func SetLogger(l Logger) {
	logger = l
}

func InternalServerError(err error, options ...Option) Problem {
	e := Problem{
		Status:     http.StatusInternalServerError,
		Title:      "Unexpected error occurred.",
		Extensions: map[string]interface{}{},
	}

	if logger != nil {
		id := logger(err)

		if id != "" {
			WithExtension("tracking_code", id)(&e)
			WithDetail(fmt.Sprintf("An unhandled error occurred, but we caught it. Please send us the tracking code: %s", id))(&e)
		} else {
			WithDetail("An unhandled error occurred")(&e)
		}
	}

	for i := range options {
		options[i](&e)
	}

	return e
}

func BadRequest(detail string, options ...Option) Problem {
	e := Problem{
		Status:     http.StatusBadRequest,
		Title:      "Invalid request inputs received.",
		Detail:     detail,
		Extensions: map[string]interface{}{},
	}

	for i := range options {
		options[i](&e)
	}

	return e
}

func Unauthorized(detail string, options ...Option) Problem {
	e := Problem{
		Status:     http.StatusUnauthorized,
		Title:      "Unauthorized request received.",
		Detail:     detail,
		Extensions: map[string]interface{}{},
	}

	for i := range options {
		options[i](&e)
	}

	return e
}

func Forbidden(detail string, options ...Option) Problem {
	e := Problem{
		Status:     http.StatusForbidden,
		Title:      "You are not allowed to do this request.",
		Detail:     detail,
		Extensions: map[string]interface{}{},
	}

	for i := range options {
		options[i](&e)
	}

	return e
}

func NotFound(detail string, options ...Option) Problem {
	e := Problem{
		Status:     http.StatusNotFound,
		Title:      "The resource not found.",
		Detail:     detail,
		Extensions: map[string]interface{}{},
	}

	for i := range options {
		options[i](&e)
	}

	return e
}

func Conflict(detail string, options ...Option) Problem {
	e := Problem{
		Status:     http.StatusConflict,
		Title:      "There is a conflict in your request.",
		Detail:     detail,
		Extensions: map[string]interface{}{},
	}

	for i := range options {
		options[i](&e)
	}

	return e
}
