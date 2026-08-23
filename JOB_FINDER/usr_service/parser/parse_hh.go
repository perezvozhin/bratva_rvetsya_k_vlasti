package parser

import (
	"JOB_FINDER/internals/domain"
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
)

const (
	userAgent = "HH-User-Agent"
)

func (p *ParserService) ParseHH(log *zap.SugaredLogger) ([]domain.Vacancy, error) {
	// FIXME: переписать логику транспорта hh.ru -- попробовать завебскрапить???

	// // FIXME: ынести в параметр функции
	// baseURL := "https://api.hh.ru/vacancies"
	// u, err := url.Parse(baseURL)
	// if err != nil {
	// 	log.Errorw("failed to parse URL", zap.Error(err))
	// 	return
	// }

	// // FIXME: Вынести в отдельную структуру
	// // FIXME: Нужны также отдельно заранее подготовленные кейсы (или пользователь сам будет вводить, а не иметь заготовки?)
	// // Будут скорее всего использовать только след параметры
	// // TODO salary, currency, only_with_salary
	// params := url.Values{
	// 	"text":       {`NAME:(Golang OR "Go разработчик" OR "Go backend")`},
	// 	"experience": {"noExperience"},
	// 	"area":       {"1"},
	// 	"page":       {"0"},
	// 	"per_page":   {"10"},
	// }

	// u.RawQuery = params.Encode()

	// req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	// if err != nil {
	// 	log.Errorw("failed to send request", zap.Error(err))
	// 	return
	// }

	// req.Header.Set(userAgent, "JobFinder/1.0 (jobsfind3r@gmail.com)")

	// resp, err := c.Do(req)
	// if err != nil {
	// 	log.Errorw("failed to get response", zap.Error(err))
	// 	return
	// }
	// defer resp.Body.Close()

	// // Для теста
	// fmt.Println("статус:", resp.StatusCode)

	// bodyBytes, _ := io.ReadAll(resp.Body)
	// fmt.Println("JSON:", string(bodyBytes))

	// // FIXME: Пока просто предположим, что прилетает 200 OK, хендлинг ошибок потом.

	// // var hhResponse HHResponse
	// // if err := json.NewDecoder(resp.Body).Decode(&hhResponse); err != nil {
	// // 	log.Errorw("failed to decode JSON", zap.Error(err))
	// // }

	bytes, err := os.ReadFile("./usr_service/parser/mock_vacancies.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var hhResponse HHResponse
	if err := json.Unmarshal(bytes, &hhResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	vacancies := hhToVacancy(hhResponse)
	return vacancies, nil
}
