package app

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

type RequestScope interface {
	Now() time.Time
}

type requestScope struct {
	now       time.Time
	requestID string
}

func newRequestScope(now time.Time, logger *zerolog.Logger, request *http.Request) RequestScope {
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		logger.Debug().
			Str("RequestID", requestID).
			Msg("Request")
	}
	return &requestScope{
		now:       now,
		requestID: requestID,
	}
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}
