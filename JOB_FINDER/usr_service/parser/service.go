package parser

import "JOB_FINDER/internals/repository/sqlite"

type ParserService struct {
	client *ParserClient
	repo   *sqlite.VacancyRepo
}

func NewParserService(client *ParserClient, repo *sqlite.VacancyRepo) *ParserService {
	return &ParserService{
		client: client,
		repo:   repo,
	}
}
