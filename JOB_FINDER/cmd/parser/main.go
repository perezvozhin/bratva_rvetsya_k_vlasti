package main

import (
	loggersystem "github.com/perezvozhin/bratva_rvetsya_k_vlasti/JOB_FINDER/internals/logger"
	"github.com/perezvozhin/bratva_rvetsya_k_vlasti/JOB_FINDER/usr_service/parser"
)

func main() {
	logger := loggersystem.Init()

	client := parser.NewParserClient()
	client.ParseHH(logger)
}
