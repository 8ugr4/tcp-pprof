package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"tcp-server-project/profiler"
)

// go-wrk -d 20 http://localhost:8080
// go tool pprof -seconds=5 localhost:8080/debug/pprof/profile?seconds=5
// use top and web with pprof

func main() {
	var option string
	// options are "local", "http"
	if _, err := fmt.Scanf("%s", &option); err != nil {
		log.Fatal(err)
	}
	if option == "local" {
		profiler.Profile()
	}
	if option == "http" {
		profiler.InitHTTPServer()
		// after initializing server, will wait 30 seconds
		// you have to run these yourself.
		// > go-wrk -d 20 http://localhost:8080
		// > go tool pprof -seconds=5 localhost:8080/debug/pprof/profile?seconds=5
	}
}
