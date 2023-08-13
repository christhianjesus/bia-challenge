package period

type PeriodStrategyFactory interface {
	CreatePeriodStrategy(kind PeriodKind) (PeriodStrategy, error)
}
