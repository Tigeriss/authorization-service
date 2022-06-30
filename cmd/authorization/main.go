package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func lowercaseHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusForbidden)
		log.Println("Should ba a POST Method")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	receivedBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Can not read body of request")
		return
	}
	textToChange := string(receivedBytes)
	result := strings.ToLower(textToChange)
	log.Println(result)

	w.Write([]byte(result))
	return
}

func uppercaseHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusForbidden)
		log.Println("Should ba a POST Method")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	receivedBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Can not read body of request")
		return
	}
	textToChange := string(receivedBytes)
	result := strings.ToUpper(textToChange)
	log.Println(result)

	w.Write([]byte(result))
	return
}

func main() {

	http.HandleFunc("/api/lowercase", lowercaseHandle)
	http.HandleFunc("/api/uppercase", uppercaseHandler)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
