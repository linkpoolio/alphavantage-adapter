package av

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestGetSupportedExchanges(t *testing.T) {
	c := NewClient(os.Getenv("API_KEY"))
	r, err := c.HttpGet(map[string]string {
		"function": "TIME_SERIES_MONTHLY_ADJUSTED",
		"symbol": "MSFT",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, r["Meta Data"])
}