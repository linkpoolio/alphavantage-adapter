package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/linkpoolio/alpha-vantage-cl-ea/web"
)

func main() {
	web.InitialiseConfig()

	log.Print("chainlink alpha vantage adaptor")
	log.Printf("starting to serve on port %d", web.Conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", web.Conf.Port), web.Api().MakeHandler()))
}
