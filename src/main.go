package main

import (
	"fmt"
	"time"
)

func runEmbeddedCtx() {

	var (
		srvr Server
		err  error
	)

	if err = srvr.Initialize(); err != nil {
		fmt.Println("Failed to Initialize Server")
	}

	go func(srvrObj *Server) {
		time.Sleep(2 * time.Second)
		srvr.Cncl()
	}(&srvr)

	fmt.Println(srvr.Perform())
}

func runArgsCtx() {
	Args()
}

func main() {

	ctx_mode := 2

	if ctx_mode == 1 {
		fmt.Println("Running Embedded Context")
		runEmbeddedCtx()
	} else {
		fmt.Println("Running Args based Context")
		runArgsCtx()
	}
}
