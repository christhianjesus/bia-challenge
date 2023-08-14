package period

import (
	"errors"

	"github.com/christhianjesus/bia-challenge/internal/application/period/strategies"
	"github.com/christhianjesus/bia-challenge/internal/domain/period"
)

func CreatePeriodStrategy(kind period.KindPeriod) (PeriodStrategy, error) {
	switch kind {
	case period.Monthly:
		return &strategies.MonthlyPeriodStrategy{}, nil
	case period.Weekly:
		return &strategies.WeeklyPeriodStrategy{}, nil
	case period.Daily:
		return &strategies.DailyPeriodStrategy{}, nil
	default:
		return nil, errors.New("unsupported period type")
	}
}
