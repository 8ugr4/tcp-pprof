package profiler

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func InitHTTPServer() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(30 * time.Second)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
