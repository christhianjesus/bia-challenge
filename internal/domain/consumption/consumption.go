package consumption

import (
	"time"
)

type Consumption interface {
	ID() string
	MeterID() int
	ActiveEnergy() float64
	ReactiveEnergy() float64
	CapacitiveReactive() float64
	Solar() float64
	Date() time.Time
}

func NewConsumption(
	id string,
	meterID int,
	activeEnergy, reactiveEnergy, capacitiveReactive, solar float64,
	date time.Time,
) Consumption {
	return &consumption{
		id,
		meterID,
		activeEnergy,
		reactiveEnergy,
		capacitiveReactive,
		solar,
		date,
	}
}

type consumption struct {
	id                 string
	meterID            int
	activeEnergy       float64
	reactiveEnergy     float64
	capacitiveReactive float64
	solar              float64
	date               time.Time
}

func (c *consumption) ID() string {
	return c.id
}

func (c *consumption) MeterID() int {
	return c.meterID
}

func (c *consumption) ActiveEnergy() float64 {
	return c.activeEnergy
}

func (c *consumption) ReactiveEnergy() float64 {
	return c.reactiveEnergy
}

func (c *consumption) CapacitiveReactive() float64 {
	return c.capacitiveReactive
}

func (c *consumption) Solar() float64 {
	return c.solar
}

func (c *consumption) Date() time.Time {
	return c.date
}
