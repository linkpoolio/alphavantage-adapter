package web

import (
	"flag"
)

type Config struct {
	Port 		   int
	APIKey         string
}

var Conf *Config

func InitialiseConfig() {
	Conf = &Config{}
	flag.IntVar(&Conf.Port, "p", 8080, "Port number to serve")
	flag.StringVar(&Conf.APIKey, "apiKey", "", "The API Key for Alpha Vantage")
	flag.Parse()
}