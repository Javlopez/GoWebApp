package main

import (
	"net/http"
	"io/ioutil"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	path := "public/" + request.URL.Path
	data, err := ioutil.ReadFile(string(path))

	if err ==  nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
