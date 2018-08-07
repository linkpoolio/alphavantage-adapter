package web

import (
	"github.com/ant0ine/go-json-rest/rest"
	log "github.com/sirupsen/logrus"
	"github.com/linkpoolio/alpha-vantage-cl-ea/av"
	"gopkg.in/guregu/null.v3"
	"encoding/json"
	"io/ioutil"
)

type Input struct {
	RunResult
	Data map[string]string `json:"data"`
}
type Output struct {
	RunResult
	Data map[string]*json.RawMessage `json:"data"`
}

type RunResult struct {
	JobRunID     string       `json:"jobRunId"`
	Status       string      `json:"status"`
	ErrorMessage null.String `json:"error"`
	Pending      bool        `json:"pending"`
}

var client *av.Client

func Api() *rest.Api{
	api := rest.NewApi()
	api.Use(rest.DefaultCommonStack...)
	router, err := rest.MakeRouter(
		rest.Post("/query", GetResponse),
	)
	if err != nil {
		log.Fatal(err)
	}
	if Conf.APIKey == "" {
		log.Fatal("no api key set")
	}
	client = av.NewClient(Conf.APIKey)
	api.SetApp(router)
	log.Print("api started")
	return api
}

func GetResponse(w rest.ResponseWriter, r *rest.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)

	var i Input
	err := json.Unmarshal(bytes, &i)
	if err != nil {
		writeError(w, i, err)
		return
	}

	cr, err := client.HttpGet(i.Data)
	if err != nil {
		writeError(w, i, err)
		return
	}

	var o Output
	o.RunResult = i.RunResult
	o.Data = cr

	w.WriteJson(o)
}

func writeError(w rest.ResponseWriter, i Input, err error) {
	i.ErrorMessage = null.StringFrom(err.Error())
	i.Pending = false
	w.WriteJson(i)
}