package av

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Client struct {
	APIKey   string
	BaseURL  string
	Endpoint string
}

func NewClient(apiKey string) *Client {
	c := &Client{}
	c.BaseURL = "https://www.alphavantage.co"
	c.Endpoint = "/query"
	c.APIKey = apiKey
	return c
}

func (c *Client) HttpGet(params map[string]string) (map[string]*json.RawMessage, error) {
	var om map[string]*json.RawMessage
	hc := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, c.Endpoint), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("apikey", c.APIKey)
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &om)
	if err != nil {
		return nil, err
	}
	return om, err
}
