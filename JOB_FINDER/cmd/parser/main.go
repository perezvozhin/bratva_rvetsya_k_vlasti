package main

import (
	loggersystem "JOB_FINDER/internals/logger"
	"JOB_FINDER/usr_service/parser"
)

func main() {
	logger := loggersystem.Init()

	client := parser.NewParserClient()
	client.ParseHH(logger)
}
