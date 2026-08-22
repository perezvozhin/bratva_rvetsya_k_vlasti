package parser

import "JOB_FINDER/internals/domain"

// HHResponse представляет из себя ответ HH API.
type HHResponse struct {
	Items   []HHItem `json:"items"`
	Found   int      `json:"found"`
	Page    int      `json:"page"`
	Pages   int      `json:"pages"`
	PerPage int      `json:"per_page"`
}

// HHItem описывает одну вакансию из массива items.
type HHItem struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	URL      string   `json:"alternate_url"`
	Employer Employer `json:"employer"`
	Salary   *Salary  `json:"salary"`
	Snippet  Snippet  `json:"snippet"`
}

type Employer struct {
	Name string `json:"name"`
}

type Salary struct {
	From     int    `json:"from"`
	To       int    `json:"to"`
	Currency string `json:"currency"`
}

type Snippet struct {
	Requirement    string `json:"requirement"`
	Responsibility string `json:"responsibility"`
}

func hhToVacancy(hh HHResponse) []domain.Vacancy {
	vacancies := make([]domain.Vacancy, 0, len(hh.Items))

	for _, item := range hh.Items {
		var domainSalary *domain.Salary
		if item.Salary != nil {
			domainSalary = &domain.Salary{
				From:     item.Salary.From,
				To:       item.Salary.To,
				Currency: item.Salary.Currency,
			}
		}

		vacancies = append(vacancies, domain.Vacancy{
			ID:             item.ID,
			Name:           item.Name,
			URL:            item.URL,
			EmployerName:   item.Employer.Name,
			Salary:         domainSalary,
			Requirement:    item.Snippet.Requirement,
			Responsibility: item.Snippet.Responsibility,
			Found:          hh.Found,
			Page:           hh.Page,
			Pages:          hh.Pages,
			PerPage:        hh.PerPage,
		})
	}

	return vacancies
}
