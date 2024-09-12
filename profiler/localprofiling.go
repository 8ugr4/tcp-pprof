package profiler

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

// run the function and call the profile function
// go tool pprof .\cpuprofile.perf

func Profile(d int) {
	duration := time.Duration(d)
	f := createCPUProfiler()
	err := pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		pprof.StopCPUProfile()
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(duration * time.Second)
}

func createCPUProfiler() *os.File {
	f, err := os.Create("./cpuprofile.perf")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
