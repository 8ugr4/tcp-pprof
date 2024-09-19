package main

import (
	"bufio"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"syscall"
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
	option = option[:len(option)-1] // removing crlf

	log.Printf("choose how long the duration will be (min:30)")
	durationString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("invalid input: %v", err)
	}
	durationString = durationString[:len(durationString)-1]
	duration, err := strconv.Atoi(durationString)
	fmt.Println(duration)

	log.Printf("Chosen profiling: %q, Duration: %d", option, duration)
	switch option {
	case "local":
		log.Printf("creating cpuprofile now...\n")
		profiler.Profile(duration)
	case "http":
		log.Printf("open another terminal and use the following commands.\n\n")
		log.Printf("go-wrk -d 20 http://localhost:8080")
		log.Printf("go tool pprof -seconds=5 localhost:8080/debug/pprof/profile?seconds=5\n")
		profiler.InitHTTPServer(duration)
	}

}

// reachShell : opens the shell/bash
func reachShell() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
	}
	// send shells stdin/stdout/stderr to some server if wanted
	// if some conn is available, replace 'os.st...' 's with conn
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}

// learn WINAPI and c library to be able to calling winapi from go // go syscall
// can build dynamic libs (dll's with go)
// to inject shell code into memory etc.
// openProcess, virtualAllocEx,writeProcessMemory,createRemoteThread etc

// go payloads include debugging symbols by default
// $ go build (if not built)
// $ gdb .\<app.name>
// (gdb) list
// or
// (gdb) l
// to see the source code of any built go binaries.
// (gdb) file:

// but, go compiler can omit debug symbols and strip symbol table
// so no more source code extraction from binary (I guess so?)
// var names are converted to addresses
// it makes binaries smaller. here's how:
// go build -ldflags="-s -w" <app-name>.go

/* as follows:

$ project> gdb .\tcp-server-project.exe

	GNU gdb (GDB) 10.2
	Copyright (C) 2021 Free Software Foundation, Inc.

	Reading symbols from .\tcp-server-project.exe...
	(No debugging symbols found in .\tcp-server-project.exe)
	(gdb) l
	No symbol table is loaded.  Use the "file" command.
	(gdb) file
	No executable file now.
	No symbol file now.

*/

// can also hide console window, this also suppresses the window on ext. processes
// go build -ldflags -H=windowsgui <app-name>.go
// or

func runPowershell(cmd string) ([]byte, error) {
	//c := exec.Command("/bin/sh", "-c", cmd)
	c := exec.Command("powershell.exe", cmd)
	c.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:                 true, // this one for example
		CmdLine:                    "",
		CreationFlags:              0,
		Token:                      0,
		ProcessAttributes:          nil,
		ThreadAttributes:           nil,
		NoInheritHandles:           false,
		AdditionalInheritedHandles: nil,
		ParentProcess:              0,
	}
	output, err := c.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to combine output: %v", err)
	}
	return output, nil
}
