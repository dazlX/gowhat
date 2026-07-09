package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *ResponseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.HandlerFunc, log *zap.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id := uuid.New()

		wrapper := &ResponseWriterWrapper{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		log.Info("Запрос на:",
			zap.String("ID: ", id.String()),
			zap.String("Метод: ", r.Method),
			zap.String("Путь: ", r.URL.Path))

		next(wrapper, r)

		log.Info("Ответ",
			zap.String("ID: ", id.String()),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("status", wrapper.statusCode),
		)
	}

}
