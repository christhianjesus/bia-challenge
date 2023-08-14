package frequency

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPeriod(t *testing.T) {
	t1 := time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	period := basePeriod{t1, t2}

	assert.NotNil(t, period)
	assert.Equal(t, t1, period.StartDate())
	assert.Equal(t, t2, period.EndDate())
}
