package domain

type Vacancy struct {
	Items                       []Item
	Found, Page, Pages, PerPage int
}

type Item struct {
	ID             string
	Name           string
	URL            string
	EmployerName   string
	Salary         *Salary
	Requirement    string
	Responsibility string
}

type Salary struct {
	From     int
	To       int
	Currency string
}

func NewVacancy(
	items []Item,
	found, page, pages, perPage int,
) Vacancy {
	return Vacancy{
		Items:   items,
		Found:   found,
		Page:    page,
		Pages:   pages,
		PerPage: perPage,
	}
}
