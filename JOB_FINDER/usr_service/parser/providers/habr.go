package providers

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"JOB_FINDER/caller"
	"JOB_FINDER/internals/domain"
)

const (
	habrBaseURL     = "https://career.habr.com/api/frontend/vacancies"
	maxHabrPages    = 24
	habrConcurrency = 6
)

type Habr struct {
	client *caller.Caller
}

func NewHabr(client *caller.Caller) *Habr {
	return &Habr{client: client}
}

func (h *Habr) Name() string { return "habr" }

type habrResponse struct {
	List []habrItem `json:"list"`
	Meta struct {
		TotalResults int `json:"totalResults"`
		PerPage      int `json:"perPage"`
		CurrentPage  int `json:"currentPage"`
		TotalPages   int `json:"totalPages"`
	} `json:"meta"`
}

type habrItem struct {
	ID         int64  `json:"id"`
	Href       string `json:"href"`
	Title      string `json:"title"`
	RemoteWork bool   `json:"remoteWork"`
	Qualifi    string `json:"qualification"`
	Company    struct {
		Title string `json:"title"`
	} `json:"company"`
	Salary *struct {
		From     int    `json:"from"`
		To       int    `json:"to"`
		Currency string `json:"currency"`
	} `json:"salary"`
	Skills []struct {
		Title string `json:"title"`
	} `json:"skills"`
	Locations []struct {
		Title string `json:"title"`
	} `json:"locations"`
}

func (h *Habr) fetchPage(ctx context.Context, q domain.SearchQuery, page int) (habrResponse, error) {
	params := url.Values{
		"q":    {q.Text},
		"page": {strconv.Itoa(page)},
		"type": {"all"},
	}
	if q.OnlyRemote {
		params.Set("remote", "true")
	}

	var resp habrResponse
	if err := h.client.GetJSON(ctx, habrBaseURL+"?"+params.Encode(), nil, &resp); err != nil {
		return habrResponse{}, fmt.Errorf("habr search: %w", err)
	}
	return resp, nil
}

func (h *Habr) Search(ctx context.Context, q domain.SearchQuery) ([]domain.Vacancy, error) {
	first := q.Page
	if first < 1 {
		first = 1
	}

	head, err := h.fetchPage(ctx, q, first)
	if err != nil {
		return nil, err
	}

	out := make([]domain.Vacancy, 0, len(head.List))
	for _, item := range head.List {
		out = append(out, h.toVacancy(item))
	}

	last := head.Meta.TotalPages
	if limit := first + maxHabrPages - 1; last > limit {
		last = limit
	}
	if last <= first {
		return out, nil
	}

	rest := make([][]domain.Vacancy, last-first)
	sem := make(chan struct{}, habrConcurrency)
	var wg sync.WaitGroup

	for page := first + 1; page <= last; page++ {
		wg.Add(1)

		go func(page int) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			resp, err := h.fetchPage(ctx, q, page)
			if err != nil {
				return
			}

			part := make([]domain.Vacancy, 0, len(resp.List))
			for _, item := range resp.List {
				part = append(part, h.toVacancy(item))
			}
			rest[page-first-1] = part
		}(page)
	}

	wg.Wait()

	for _, part := range rest {
		out = append(out, part...)
	}

	return out, nil
}

func (h *Habr) toVacancy(item habrItem) domain.Vacancy {
	var salary *domain.Salary
	if item.Salary != nil && (item.Salary.From > 0 || item.Salary.To > 0) {
		salary = &domain.Salary{
			From:     item.Salary.From,
			To:       item.Salary.To,
			Currency: strings.ToUpper(item.Salary.Currency),
		}
	}

	skills := make([]string, 0, len(item.Skills))
	for _, s := range item.Skills {
		skills = append(skills, s.Title)
	}

	return domain.Vacancy{
		ID:           "habr:" + strconv.FormatInt(item.ID, 10),
		Source:       "habr",
		Name:         item.Title,
		URL:          "https://career.habr.com" + item.Href,
		EmployerName: item.Company.Title,
		Salary:       salary,
		Experience:   item.Qualifi,
		Requirement:  strings.Join(skills, ", "),
		Remote:       item.RemoteWork,
	}
}
