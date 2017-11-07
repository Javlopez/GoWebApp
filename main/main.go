package main

import (
	"net/http"
	//"io/ioutil"
	"strings"
	"os"
	"bufio"
)

func main(){
	http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	path := "public/" + request.URL.Path
	f, err := os.Open(path)
	//data, err := ioutil.ReadFile(string(path))

	if err ==  nil {

		bufferedReader := bufio.NewReader(f)
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "text/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(w)
		//w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
