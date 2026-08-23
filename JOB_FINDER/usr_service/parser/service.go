package parser

import (
	"JOB_FINDER/caller"
	"JOB_FINDER/internals/repository/sqlite"
)

type ParserService struct {
	client *caller.Caller
	repo   *sqlite.VacancyRepo
}

func NewParserService(client *caller.Caller, repo *sqlite.VacancyRepo) *ParserService {
	return &ParserService{
		client: client,
		repo:   repo,
	}
}
