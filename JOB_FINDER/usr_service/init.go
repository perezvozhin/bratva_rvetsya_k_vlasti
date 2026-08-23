package usr_service

import (
	"JOB_FINDER/usr_service/parser"
	"os"

	"go.uber.org/zap"
)

type Usr_service struct {
	logger        *zap.SugaredLogger
	cv            *os.File
	ParserService *parser.ParserService
}

func Init(logger *zap.SugaredLogger, cv *os.File, ParserService *parser.ParserService) *Usr_service {

	return &Usr_service{logger: logger, cv: cv, ParserService: ParserService}
}
