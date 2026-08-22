package usr_service

import (
	"os"

	"go.uber.org/zap"
)

type Usr_service struct {
	logger *zap.SugaredLogger
	cv     *os.File
}

func Init(logger *zap.SugaredLogger, cv *os.File) *Usr_service {

	return &Usr_service{logger: logger, cv: cv}
}
