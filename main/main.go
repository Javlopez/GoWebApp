package main

import (
	"net/http"
	//"io/ioutil"
	//"strings"
	//"os"
	//"bufio"
)

func main(){

	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}
