package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Quote struct {
	Type  string `json:"type"`
	Value struct {
		ID   int    `json:"id"`
		Joke string `json:"joke"`
	} `json:"value"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/quote", inspQuote)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func inspQuote(writer http.ResponseWriter, request *http.Request) {
	response, err := http.Get("http://api.icndb.com/jokes/random")
	if err != nil {
		log.Fatal(err)
	} else {
		var quote Quote
		body, err := ioutil.ReadAll(io.LimitReader(response.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := response.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &quote); err != nil {
			writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
			if err := json.NewEncoder(writer).Encode(err); err != nil {
				panic(err)
			}
		}
		quoteToSend := quote
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(writer).Encode(quoteToSend); err != nil {
			panic(err)
		}
	}
}
