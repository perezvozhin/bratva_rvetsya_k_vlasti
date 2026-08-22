package domain

type Vacancy struct {
	ID                          string
	Name                        string
	URL                         string
	EmployerName                string
	Salary                      *Salary
	Requirement                 string
	Responsibility              string
	Found, Page, Pages, PerPage int
}

type Salary struct {
	From     int
	To       int
	Currency string
}
