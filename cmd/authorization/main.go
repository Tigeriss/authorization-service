package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func heading(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("X-Powered-By", "Tigeriss' Engine")
		f(writer, request)
	}
}

func private(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("X-Token")
		if token != "MagicKey" {
			writer.WriteHeader(http.StatusForbidden)
			log.Println("wrong X-Token")
			return
		}
		f(writer, request)
	}
}

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

	http.HandleFunc("/api/lowercase", private(heading(lowercaseHandle)))
	http.HandleFunc("/api/uppercase", heading(uppercaseHandler))

	log.Fatal(http.ListenAndServe(":9090", nil))
}
