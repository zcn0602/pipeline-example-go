package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "Hello World! Abraham was here  2020/05/24  15:14:00.60"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
