package consumption

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConsumption(t *testing.T) {
	id := "993c644a-bace-4e1e-9b42-2a9c18aadcfe"
	meterID := 2
	activeEnergy := 17234.731809999997
	reactiveEnergy := 10516.07749
	capacitiveReactive := 0.0
	solar := 0.6388935728544158
	date := time.Now()

	consumption := NewConsumption(id, meterID, activeEnergy, reactiveEnergy, capacitiveReactive, solar, date)

	assert.NotNil(t, consumption)
	assert.Equal(t, id, consumption.ID())
	assert.Equal(t, meterID, consumption.MeterID())
	assert.Equal(t, activeEnergy, consumption.ActiveEnergy())
	assert.Equal(t, reactiveEnergy, consumption.ReactiveEnergy())
	assert.Equal(t, capacitiveReactive, consumption.CapacitiveReactive())
	assert.Equal(t, solar, consumption.Solar())
	assert.Equal(t, date, consumption.Date())
}
