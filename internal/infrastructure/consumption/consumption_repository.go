package consumption

import (
	"context"
	"database/sql"
	"time"

	"github.com/christhianjesus/bia-challenge/internal/domain/consumption"
	"github.com/lib/pq"
)

type postgreSQLConsumptionRepository struct {
	db *sql.DB
}

func NewPostgreSQLConsumptionRepository(db *sql.DB) consumption.ConsumptionRepository {
	return &postgreSQLConsumptionRepository{db: db}
}

func (r *postgreSQLConsumptionRepository) GetByMetersIDsAndDateRange(ctx context.Context, metersIDs []int, startDate, endDate time.Time) ([]*consumption.Consumption, error) {
	query := `
        SELECT id, meter_id, active_energy, reactive_energy, capacitive_reactive, solar, date FROM consumptions
        WHERE meter_id = ANY($1) AND date >= $2 AND date <= $3
    `
	rows, err := r.db.QueryContext(ctx, query, pq.Array(metersIDs), startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumptions []*consumption.Consumption
	for rows.Next() {
		var (
			id                 string
			meterID            int
			activeEnergy       float64
			reactiveEnergy     float64
			capacitiveReactive float64
			solar              float64
			date               time.Time
		)
		err := rows.Scan(&id, &meterID, &activeEnergy, &reactiveEnergy, &capacitiveReactive, &solar, &date)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, consumption.NewConsumption(
			id,
			meterID,
			activeEnergy,
			reactiveEnergy,
			capacitiveReactive,
			solar,
			date,
		))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return consumptions, nil
}
