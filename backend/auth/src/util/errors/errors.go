package errors

import "time"

type httpErr struct {
	message string
	code    int
	url     string
}

var (
	BAD_REQUEST           = httpErr{message: "bad request", code: 400}
	UNAUTHORIZED          = httpErr{message: "unauthorized", code: 401}
	BAD_GATEWAY           = httpErr{message: "bad gateway", code: 402}
	FORBIDDEN             = httpErr{message: "forbidden", code: 403}
	NOT_FOUND             = httpErr{message: "not found", code: 404}
	REQUEST_TIMEOUT       = httpErr{message: "request timeout", code: 408}
	CONFLICT              = httpErr{message: "conflict", code: 409}
	TOO_MANY_REQUESTS     = httpErr{message: "too many requests", code: 429}
	INTERNAL_SERVER_ERROR = httpErr{message: "internal server error", code: 500}
	SERVICE_UNAVAILABLE   = httpErr{message: "service unavailable", code: 503}
	GATEWAY_TIMEOUT       = httpErr{message: "gateway timeout", code: 504}
)

const (
	TIMEOUT_TIME = 5 * time.Second
)

func HttpError(errCode string, url string) httpErr {
	switch errCode {
	case "BAD_REQUEST":
		BAD_REQUEST.url = url
		return BAD_REQUEST
	case "UNAUTHORIZED":
		UNAUTHORIZED.url = url
		return UNAUTHORIZED
	case "BAD_GATEWAY":
		BAD_GATEWAY.url = url
		return BAD_GATEWAY
	case "FORBIDDEN":
		FORBIDDEN.url = url
		return FORBIDDEN
	case "NOT_FOUND":
		NOT_FOUND.url = url
		return NOT_FOUND
	case "REQUEST_TIMEOUT":
		REQUEST_TIMEOUT.url = url
		return REQUEST_TIMEOUT
	case "CONFLICT":
		CONFLICT.url = url
		return CONFLICT
	case "TOO_MANY_REQUESTS":
		TOO_MANY_REQUESTS.url = url
		return TOO_MANY_REQUESTS
	case "INTERNAL_SERVER_ERROR":
		INTERNAL_SERVER_ERROR.url = url
		return INTERNAL_SERVER_ERROR
	case "SERVICE_UNAVAILABLE":
		SERVICE_UNAVAILABLE.url = url
		return SERVICE_UNAVAILABLE
	case "GATEWAY_TIMEOUT":
		GATEWAY_TIMEOUT.url = url
		return GATEWAY_TIMEOUT
	default:
		return httpErr{message: "unknown error", code: 000, url: url}
	}
}
