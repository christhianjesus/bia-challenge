package strategies

import (
	"errors"

	"github.com/christhianjesus/bia-challenge/internal/application/period"
)

type PeriodStrategyFactory struct{}

func (pf *PeriodStrategyFactory) CreatePeriodStrategy(kind period.PeriodKind) (period.PeriodStrategy, error) {
	switch kind {
	case period.Monthly:
		return &MonthlyPeriodStrategy{}, nil
	case period.Weekly:
		return &WeeklyPeriodStrategy{}, nil
	case period.Daily:
		return &DailyPeriodStrategy{}, nil
	default:
		return nil, errors.New("unsupported period type")
	}
}
