package http

import (
	"context"
	"net/http"
	"webapi/pkg/app/errors"
	"webapi/pkg/interfaces/http/models"
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
	Txn     interface{}
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

func BadRequest(body models.ErrorResponse, headers http.Header) HttpResponse {
	body.StatusCode = 400
	return HttpResponse{
		StatusCode: 400,
		Body:       body,
		Headers:    headers,
	}
}

func Unauthorized(body models.ErrorResponse, headers http.Header) HttpResponse {
	body.StatusCode = 401
	return HttpResponse{
		StatusCode: 401,
		Body:       body,
		Headers:    headers,
	}
}

func Forbiden(body models.ErrorResponse, headers http.Header) HttpResponse {
	body.StatusCode = 403
	return HttpResponse{
		StatusCode: 403,
		Body:       body,
		Headers:    headers,
	}
}

func NotFound(body models.ErrorResponse, headers http.Header) HttpResponse {
	body.StatusCode = 404
	return HttpResponse{
		StatusCode: 404,
		Body:       body,
		Headers:    headers,
	}
}

func Conflict(body models.ErrorResponse, headers http.Header) HttpResponse {
	body.StatusCode = 409
	return HttpResponse{
		StatusCode: 409,
		Body:       body,
		Headers:    headers,
	}
}

func InternalServerError(body models.ErrorResponse, headers http.Header) HttpResponse {
	body.StatusCode = 500
	return HttpResponse{
		StatusCode: 500,
		Body:       body,
		Headers:    headers,
	}
}

func ErrorResponseMapper(err error, headers http.Header) HttpResponse {
	switch err.(type) {
	case errors.BadRequestError:
		return BadRequest(models.ErrorResponse{Message: err.Error()}, headers)
	case errors.NotFoundError:
		return NotFound(models.ErrorResponse{Message: err.Error()}, headers)
	case errors.ConflictError:
		return Conflict(models.ErrorResponse{Message: err.Error()}, headers)
	default:
		return InternalServerError(models.ErrorResponse{Message: err.Error()}, headers)
	}
}
