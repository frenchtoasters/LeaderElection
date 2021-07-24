package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ElectionResults struct {
	Leader   string `json:"leader"`
	Election string `json:"election"`
}

type electionHandler struct{}

func leaderHandler(resp http.ResponseWriter, req *http.Request) {
	response, err := http.Get("http://localhost:4040/")
	if err != nil {
		fmt.Errorf("failed to connect to webservice: %v", err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("error reading body: %v", err)
	}

	var results ElectionResults
	err = json.Unmarshal(data, &results)
	if err != nil {
		fmt.Errorf("failed to unmarshal data: %v", err)
	}
	fmt.Printf("Current Leader of %s: %s", results.Election, results.Leader)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/leader", leaderHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	s.ListenAndServe()
}
