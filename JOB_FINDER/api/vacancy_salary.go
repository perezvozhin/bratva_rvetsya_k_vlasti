package api

import (
	"fmt"
	"strings"

	"JOB_FINDER/internals/domain"
)

func formatSalary(s *domain.Salary) string {
	if s == nil || (s.From == 0 && s.To == 0) {
		return "Не указана"
	}

	cur := currencySign(s.Currency)

	switch {
	case s.From > 0 && s.To > 0:
		return fmt.Sprintf("%s – %s %s", thousands(s.From), thousands(s.To), cur)
	case s.From > 0:
		return fmt.Sprintf("от %s %s", thousands(s.From), cur)
	default:
		return fmt.Sprintf("до %s %s", thousands(s.To), cur)
	}
}

func currencySign(code string) string {
	switch strings.ToUpper(code) {
	case "RUR", "RUB":
		return "₽"
	case "USD":
		return "$"
	case "EUR":
		return "€"
	default:
		return code
	}
}

func thousands(n int) string {
	s := fmt.Sprintf("%d", n)

	var b strings.Builder
	for i, r := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteRune(' ')
		}
		b.WriteRune(r)
	}
	return b.String()
}
