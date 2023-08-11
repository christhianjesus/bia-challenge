package period

import "errors"

type PeriodFactory struct{}

func (pf *PeriodFactory) CreatePeriodStrategy(periodType PeriodType) (PeriodStrategy, error) {
	switch periodType {
	case Monthly:
		return &MonthlyPeriodStrategy{}, nil
	case Weekly:
		return &WeeklyPeriodStrategy{}, nil
	case Daily:
		return &DailyPeriodStrategy{}, nil
	default:
		return nil, errors.New("unsupported period type")
	}
}
