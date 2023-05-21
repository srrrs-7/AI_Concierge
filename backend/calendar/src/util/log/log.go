package log

import (
	"log"
	"net/http"
)

var (
	std  = NewLogger()
	Info = std.Info
)

type Logger interface {
	Info(msg any)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Info(msg any) {
	log.Println(msg)
}

func NewLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("receive request: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
