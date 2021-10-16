package http

import (
	"context"
	"net/http"
)

type HttpMethod string

type HttpResponse struct {
	StatusCode int
	Body       interface{}
	Headers    http.Header
}

type HttpRequest struct {
	Body    []byte
	Headers http.Header
	Params  map[string]string
	Auth    interface{}
	Ctx     context.Context
}

type ErrorMessage struct {
	StatusCode int    `json:"statusCode`
	Message    string `json:"message"`
}

func Ok(body interface{}, headers http.Header) HttpResponse {
	return HttpResponse{
		StatusCode: 200,
		Body:       body,
		Headers:    headers,
	}
}

func Created(body interface{}, headers http.Header) HttpResponse {
	return HttpResponse{
		StatusCode: 201,
		Body:       body,
		Headers:    headers,
	}
}

func NoContent(headers http.Header) HttpResponse {
	return HttpResponse{
		StatusCode: 204,
		Headers:    headers,
	}
}

func BadRequest(msg string, headers http.Header) HttpResponse {
	return HttpResponse{
		StatusCode: 400,
		Body: ErrorMessage{
			StatusCode: 400,
			Message:    msg,
		},
		Headers: headers,
	}
}

func Unauthorized(msg string, headers http.Header) HttpResponse {
	return HttpResponse{
		StatusCode: 401,
		Body: ErrorMessage{
			StatusCode: 401,
			Message:    msg,
		},
		Headers: headers,
	}
}

func Forbiden(msg string, headers http.Header) HttpResponse {
	return HttpResponse{
		StatusCode: 403,
		Body: ErrorMessage{
			StatusCode: 403,
			Message:    msg,
		},
		Headers: headers,
	}
}

func NotFound(msg string, headers http.Header) HttpResponse {
	return HttpResponse{
		StatusCode: 404,
		Body: ErrorMessage{
			StatusCode: 404,
			Message:    msg,
		},
		Headers: headers,
	}
}

func Conflict(msg string, headers http.Header) HttpResponse {
	return HttpResponse{
		StatusCode: 409,
		Body: ErrorMessage{
			StatusCode: 409,
			Message:    msg,
		},
		Headers: headers,
	}
}

func InternalServerError(msg string, headers http.Header) HttpResponse {
	return HttpResponse{
		StatusCode: 500,
		Body: ErrorMessage{
			StatusCode: 500,
			Message:    msg,
		},
		Headers: headers,
	}
}
