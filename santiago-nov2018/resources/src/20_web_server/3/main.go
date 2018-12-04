package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var results []string

// START OMIT
func GetHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		results = append(results, string(body))
		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// END OMIT
func main() {
	results = append(results, time.Now().Format(time.RFC3339))
	http.HandleFunc("/", GetHandler)
	http.HandleFunc("/post", PostHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
