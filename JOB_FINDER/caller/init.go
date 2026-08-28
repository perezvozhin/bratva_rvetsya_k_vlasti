package caller

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Caller struct {
	client *http.Client
	logger *zap.SugaredLogger
}

func NewCaller(logs *zap.SugaredLogger) *Caller {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	return &Caller{client: client, logger: logs}
}

func (c *Caller) Client() *http.Client {
	return c.client
}

func (c *Caller) Logger() *zap.SugaredLogger {
	return c.logger
}
