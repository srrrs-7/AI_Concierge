package log

import (
	"log"
	"net/http"
)

var (
	std  = NewRepository()
	Info = std.Info
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (l *Repository) Info(msg any) {
	log.Println(msg)
}

func NewLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("receive request: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
