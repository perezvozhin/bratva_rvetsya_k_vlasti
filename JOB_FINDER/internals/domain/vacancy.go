package domain

import "context"

type Vacancy struct {
	ID             string
	Source         string
	Name           string
	URL            string
	EmployerName   string
	Salary         *Salary
	Experience     string
	Requirement    string
	Responsibility string
	Remote         bool
}

type Salary struct {
	From     int
	To       int
	Currency string
}

type SearchQuery struct {
	Text       string
	Area       string
	Experience string
	OnlyRemote bool
	Page       int
	PerPage    int
}

type Provider interface {
	Name() string
	Search(ctx context.Context, q SearchQuery) ([]Vacancy, error)
}

type VacancyRepository interface {
	SaveVacancy(v Vacancy) error
	SaveVacancies(list []Vacancy) error
	ListVacancies(limit, offset int) ([]Vacancy, error)
	CountVacancies() (int, error)
	SearchVacancies(text string, limit, offset int) ([]Vacancy, error)
	CountSearch(text string) (int, error)
	GetVacancy(id string) (Vacancy, error)
}
