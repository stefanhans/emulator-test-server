package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Run starts the service.
func main() {

	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var (
	htmlStartTxt string = "<html><header><title>This is title</title></header><body><h1>"
	htmlEndTxt   string = "</h1></body></html>"

	jsonTxt string = fmt.Sprintf("{\n\t%q: %v,\n\t%q: %v,\n\t%q: %v\n}",
		"userId", 1, "id", 1, "title", "quidem molestiae enim")
)

type Album struct {
	Id         int    `json:"id"`
	Contagious bool   `json:"contagious"`
	Title      string `json:"title"`
}

func Index(w http.ResponseWriter, r *http.Request) {

	album := &Album{
		Id:         1,
		Contagious: true,
		Title:      "My first album",
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Body: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	//var registration Registration
	err = json.Unmarshal(bytes, &album)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot unmarshall JSON input: %s", err), http.StatusInternalServerError)
		return
	}

	// Marshal album
	var albumJson []byte
	albumJson, err = json.Marshal(album)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal dump: %s", err), http.StatusInternalServerError)
		return
	}

	// Response
	//_, err = fmt.Fprintf(w, "%s%s%s", htmlStartTxt, jsonTxt, htmlEndTxt)

	fmt.Printf("Response: %s\n\n", string(albumJson))
	_, err = fmt.Fprintf(w, string(albumJson))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
	return
}
