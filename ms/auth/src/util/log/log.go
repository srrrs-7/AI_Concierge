package log

import (
	"log"
	"net/http"
	"os"
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

func (l *Repository) Error(msg any) {
	log.Fatal(msg)
}

func NewLogging(next http.Handler) http.Handler {
	file, err := os.OpenFile("../log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("failed to create log file:", err)
	}
	defer file.Close()

	// ログの出力先をファイルに設定
	log.SetOutput(file)

	// ログのフォーマットを設定
	log.SetFlags(log.Ldate | log.Ltime)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("receive request: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
