package main

import (
	"github.com/linkpoolio/bridges/bridge"
	"net/http"
	"os"
)

type AlphaVantage struct {}

func (av *AlphaVantage) Run(h *bridge.Helper) (interface{}, error) {
	r := make(map[string]interface{})
	return r, h.HTTPCallWithOpts(
		http.MethodGet,
		"https://www.alphavantage.co/query",
		&r,
		bridge.CallOpts{
			Auth: bridge.NewAuth(bridge.AuthParam, "apikey", os.Getenv("API_KEY")),
			QueryPassthrough: true,
		},
	)
}

func (av *AlphaVantage) Opts() *bridge.Opts {
	return &bridge.Opts{
		Name:   "AlphaVantage",
		Lambda: true,
	}
}

func main() {
	bridge.NewServer(&AlphaVantage{}).Start(8080)
}
