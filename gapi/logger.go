package gapi

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"
)

type ResponseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (rr *ResponseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}

func (rr *ResponseRecorder) Write(body []byte) (int, error) {
	rr.body = body
	return rr.ResponseWriter.Write(body)
}

func GrpcLogger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	startTime := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(startTime)

	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	logger := log.Info()
	if err != nil {
		logger = log.Error().Err(err)
	}

	logger.
		Str("protocol", "grpc").
		Str("method", info.FullMethod).
		Int("status_code", int(statusCode)).
		Dur("duration", duration).
		Msg("received request")

	return result, err
}

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		w := &ResponseRecorder{ResponseWriter: wr, statusCode: http.StatusOK}
		handler.ServeHTTP(w, r)
		duration := time.Since(startTime)

		logger := log.Info()
		if w.statusCode != http.StatusOK {
			logger = log.Error().Bytes("body", w.body)
		}

		logger.
			Str("protocol", "http").
			Str("method", r.Method).
			Str("url", r.URL.RawPath).
			Int("status_code", int(w.statusCode)).
			Dur("duration", duration).
			Msg("received request")
	})
}
