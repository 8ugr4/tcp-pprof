package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

// go-wrk -d 200 http://localhost:8080
// go tool pprof -seconds=5 localhost:8080/debug/pprof/profile?seconds=5
// use top and web with pprof

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^(.+)@golang.org$")
	path := r.URL.Path[1:]
	match := re.FindAllStringSubmatch(path, -1)
	if match != nil {
		_, err := fmt.Fprintf(w, "Hello, %s!", match[0])
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(w, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
