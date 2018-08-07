package web

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"fmt"
	"bytes"
	"os"
)

type MetaData struct { // Weird keys in this API...
	Information string `json:"1. Information"`
	Symbol string `json:"2. Symbol"`
}

func init() {
	InitialiseConfig()
	Conf.APIKey = os.Getenv("API_KEY")
}

func TestTimeSeriesMonthlyAdjusted(t *testing.T) {
	i := map[string]string {
		"function": "TIME_SERIES_MONTHLY_ADJUSTED",
		"symbol": "MSFT",
	}
	r := getResponse(i)

	var m MetaData
	err := json.Unmarshal(*r.Data["Meta Data"], &m)
	assert.NoError(t, err)

	assert.Equal(t, m.Symbol, "MSFT")
	assert.Equal(t, m.Information, "Monthly Adjusted Prices and Volumes")
}

func getResponse(input map[string]string) *Output {
	s := httptest.NewServer(Api().MakeHandler())
	defer s.Close()

	var runResult Input
	runResult.JobRunID = "1234"
	runResult.Data = input
	b, err := json.Marshal(&runResult)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(fmt.Sprintf("%s/query", s.URL), "application/json", bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	priceBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	o := Output{}
	err = json.Unmarshal(priceBody, &o)
	if err != nil {
		log.Fatal(err)
	}

	return &o
}