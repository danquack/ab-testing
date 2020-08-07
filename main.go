package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ArgsHandler struct {
	filePath string
}

func HeadersHandler(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func (ah *ArgsHandler) Handler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(ah.filePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	fmt.Fprintf(w, string(data))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	response := os.Args[1:]
	slashHandler := &ArgsHandler{filePath: response[0]}
	http.HandleFunc("/", slashHandler.Handler)
	http.HandleFunc("/headers", HeadersHandler)

	http.ListenAndServe(":9000", logRequest(http.DefaultServeMux))
}
