package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	fmt.Println("Started server at :8000")
	_ = http.ListenAndServe(":8080", nil)
}
