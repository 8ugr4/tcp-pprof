package main

import (
	"bufio"
	"log"
	_ "net/http/pprof"
	"os"
	"strconv"
	"tcp-server-project/profiler"
)

// go-wrk -d 20 http://localhost:8080
// go tool pprof -seconds=5 localhost:8080/debug/pprof/profile?seconds=5
// use top and web with pprof

func main() {
	reader := bufio.NewReader(os.Stdin)

	log.Printf("choose profiling: http or local")
	option, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("invalid input: %v", err)
	}
	option = option[:len(option)-2] // removing crlf

	log.Printf("choose how long the duration will be (min:30)")
	durationString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("invalid input: %v", err)
	}
	durationString = durationString[:len(durationString)-2]
	duration, err := strconv.Atoi(durationString)

	log.Printf("Chosen profiling: %q, Duration: %d", option, duration)
	switch option {
	case "local":
		log.Printf("creating cpuprofile now...\n")
		profiler.Profile(10)
	case "http":
		log.Printf("open another terminal and use the following commands.\n\n")
		log.Printf("go-wrk -d 20 http://localhost:8080")
		log.Printf("go tool pprof -seconds=5 localhost:8080/debug/pprof/profile?seconds=5\n")
		profiler.InitHTTPServer(duration)
	}

}
