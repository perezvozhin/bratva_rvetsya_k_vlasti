package caller

import (
	"net"
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
		//no timeout, because video uploads and analysis can take more time
		Transport: &http.Transport{
			DialContext:           (&net.Dialer{Timeout: 10 * time.Second}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 60 * time.Second,
		},
	}

	return &Caller{client: client, logger: logs}
}

func (c *Caller) Client() *http.Client {
	return c.client
}

func (c *Caller) Logger() *zap.SugaredLogger {
	return c.logger
}
