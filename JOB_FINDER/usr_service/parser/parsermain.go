package parser

//import (
//	_ "github.com/mattn/go-sqlite3"
//)

//func ParserMain() {
//	logger := loggersystem.Init()
//
//	db, err := sql.Open("sqlite3", "../out/sqlite/vacancies.db")
//	if err != nil {
//		panic(err)
//	}
//
//	client := parser.NewParserClient()
//	repo := sqlite.NewVacancyRepo(db)
//
//	parserService := parser.NewParserService(client, repo)
//	parserService.RunParserProcess(logger)
//}
