package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDailyPriceCollector(t *testing.T) {
	domain := "www.twse.com.tw"
	c := NewTaiwanDailyPriceCollector(domain)

	assert.Contains(t, c.AllowedDomains, domain)
}

func TestTaiwanDailyPriceCollector_GetMonthlyPrices(t *testing.T) {
	domain := "www.twse.com.tw"
	c := NewTaiwanDailyPriceCollector(domain)
	data, err := c.GetMonthlyPrices("2330", 2022, 11)

	assert.NotEmpty(t, data)
	assert.NoError(t, err)
}
