package main

import (
	loggersystem "JOB_FINDER/internals/logger"
	"JOB_FINDER/internals/repository/sqlite"
	"JOB_FINDER/usr_service/parser"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger := loggersystem.Init()

	db, err := sql.Open("sqlite3", "../out/sqlite/vacancies.db")
	if err != nil {
		panic(err)
	}

	client := parser.NewParserClient()
	repo := sqlite.NewVacancyRepo(db)

	parserService := parser.NewParserService(client, repo)
	parserService.RunParserProcess(logger)
}
