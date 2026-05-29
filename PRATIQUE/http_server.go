package main

import (
	"fmt"
	"net/http"
)

func greating(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello evryone this is my first go http server \n")
}

func main() {

	http.HandleFunc("/hello", greating)

	http.ListenAndServe(":8690", nil)

}
