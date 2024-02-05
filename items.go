package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	id       string
	name     string
	quantity int
}

func createItem(w http.ResponseWriter, request *http.Request) {
	var item Item

	err := json.NewDecoder(request.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Item is created successfully: %+v", item)

}

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var urlPath = r.URL.Path

		switch urlPath {

		case "/items":
			createItem(w, r)
		}
	})

	http.ListenAndServe(":1337", nil)
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Fprintln(i)
	}

}
