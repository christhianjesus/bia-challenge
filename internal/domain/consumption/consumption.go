package consumption

import (
	"time"
)

func NewConsumption(id string, meterID int, activeEnergy, reactiveEnergy, capacitiveReactive, solar float64, date time.Time) *Consumption {
	return &Consumption{
		id,
		meterID,
		activeEnergy,
		reactiveEnergy,
		capacitiveReactive,
		solar,
		date,
	}
}

type Consumption struct {
	id                 string
	meterID            int
	activeEnergy       float64
	reactiveEnergy     float64
	capacitiveReactive float64
	solar              float64
	date               time.Time
}

func (c *Consumption) ID() string {
	return c.id
}

func (c *Consumption) MeterID() int {
	return c.meterID
}

func (c *Consumption) ActiveEnergy() float64 {
	return c.activeEnergy
}

func (c *Consumption) ReactiveEnergy() float64 {
	return c.reactiveEnergy
}

func (c *Consumption) CapacitiveReactive() float64 {
	return c.capacitiveReactive
}

func (c *Consumption) Solar() float64 {
	return c.solar
}

func (c *Consumption) Date() time.Time {
	return c.date
}
