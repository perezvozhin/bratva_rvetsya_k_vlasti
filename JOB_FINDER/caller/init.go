package caller

import (
	"net/http"

	"go.uber.org/zap"
)

type Caller struct {
	client *http.Client
	logger *zap.SugaredLogger
	path   string
}

// создаем простой клиент
func Init(logs *zap.SugaredLogger, path string) *Caller {
	client := &http.Client{}

	return &Caller{client: client, logger: logs, path: path}
}
