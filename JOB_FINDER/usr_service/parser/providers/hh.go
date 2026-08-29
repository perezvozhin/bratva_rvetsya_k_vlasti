package providers

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"sync"

	"JOB_FINDER/caller"
	"JOB_FINDER/internals/domain"
)

const (
	hhBaseURL     = "https://api.hh.ru/vacancies"
	maxHHPages    = 10
	hhConcurrency = 4
)

var ErrNoUserAgent = errors.New("hh: не задан HH_USER_AGENT, зарегистрируй приложение на dev.hh.ru")

type HH struct {
	client    *caller.Caller
	userAgent string
}

func NewHH(client *caller.Caller, userAgent string) *HH {
	return &HH{client: client, userAgent: userAgent}
}

func (h *HH) Name() string { return "hh" }

type hhResponse struct {
	Items   []hhItem `json:"items"`
	Found   int      `json:"found"`
	Page    int      `json:"page"`
	Pages   int      `json:"pages"`
	PerPage int      `json:"per_page"`
}

type hhItem struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"alternate_url"`
	Employer struct {
		Name string `json:"name"`
	} `json:"employer"`
	Salary *struct {
		From     int    `json:"from"`
		To       int    `json:"to"`
		Currency string `json:"currency"`
	} `json:"salary"`
	Experience struct {
		Name string `json:"name"`
	} `json:"experience"`
	Schedule struct {
		ID string `json:"id"`
	} `json:"schedule"`
	Snippet struct {
		Requirement    string `json:"requirement"`
		Responsibility string `json:"responsibility"`
	} `json:"snippet"`
}

func (h *HH) fetchPage(ctx context.Context, q domain.SearchQuery, page, perPage int) (hhResponse, error) {
	params := url.Values{
		"text":     {q.Text},
		"page":     {strconv.Itoa(page)},
		"per_page": {strconv.Itoa(perPage)},
	}
	if q.Area != "" {
		params.Set("area", q.Area)
	}
	if q.Experience != "" {
		params.Set("experience", q.Experience)
	}
	if q.OnlyRemote {
		params.Set("schedule", "remote")
	}

	headers := map[string]string{"User-Agent": h.userAgent}

	var resp hhResponse
	if err := h.client.GetJSON(ctx, hhBaseURL+"?"+params.Encode(), headers, &resp); err != nil {
		return hhResponse{}, fmt.Errorf("hh search: %w", err)
	}
	return resp, nil
}

func (h *HH) Search(ctx context.Context, q domain.SearchQuery) ([]domain.Vacancy, error) {
	if h.userAgent == "" {
		return nil, ErrNoUserAgent
	}

	perPage := q.PerPage
	if perPage <= 0 {
		perPage = 50
	}

	first := q.Page
	if first < 0 {
		first = 0
	}

	head, err := h.fetchPage(ctx, q, first, perPage)
	if err != nil {
		return nil, err
	}

	out := make([]domain.Vacancy, 0, len(head.Items))
	for _, item := range head.Items {
		out = append(out, h.toVacancy(item))
	}

	last := head.Pages - 1
	if limit := first + maxHHPages - 1; last > limit {
		last = limit
	}
	if last <= first {
		return out, nil
	}

	rest := make([][]domain.Vacancy, last-first)
	sem := make(chan struct{}, hhConcurrency)
	var wg sync.WaitGroup

	for page := first + 1; page <= last; page++ {
		wg.Add(1)

		go func(page int) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			resp, err := h.fetchPage(ctx, q, page, perPage)
			if err != nil {
				return
			}

			part := make([]domain.Vacancy, 0, len(resp.Items))
			for _, item := range resp.Items {
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

func (h *HH) toVacancy(item hhItem) domain.Vacancy {
	var salary *domain.Salary
	if item.Salary != nil && (item.Salary.From > 0 || item.Salary.To > 0) {
		salary = &domain.Salary{
			From:     item.Salary.From,
			To:       item.Salary.To,
			Currency: item.Salary.Currency,
		}
	}

	return domain.Vacancy{
		ID:             "hh:" + item.ID,
		Source:         "hh",
		Name:           item.Name,
		URL:            item.URL,
		EmployerName:   item.Employer.Name,
		Salary:         salary,
		Experience:     item.Experience.Name,
		Requirement:    item.Snippet.Requirement,
		Responsibility: item.Snippet.Responsibility,
		Remote:         item.Schedule.ID == "remote",
	}
}
