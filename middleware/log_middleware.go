package middleware

import (
	"log"
	"net/http"
	"os"
)

type LogRequest struct{}

func (lg *LogRequest) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New(os.Stdout, "logger: ", 0)
		logger.Printf(
			"%s %s",
			r.Method, r.URL,
		)
		next(w, r)
	})
}

func (lg *LogRequest) Apply(next http.Handler) http.HandlerFunc {
	return lg.ApplyFn(next.ServeHTTP)
}
