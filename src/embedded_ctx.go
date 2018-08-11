package main

import (
	"context"
	"fmt"
	"time"
)

type Server struct {

	/*
	 * Use context to control the
	 * child goroutines
	 */
	Ctx  context.Context
	Cncl context.CancelFunc
}

func (srvr *Server) Initialize() (err error) {
	srvr.Ctx, srvr.Cncl = context.WithTimeout(context.Background(),
		time.Hour)
	return
}

func (srvr *Server) Perform() error {
	for {

		fmt.Println("Let's do something now")

		select {
		case <-srvr.Ctx.Done():
			return srvr.Ctx.Err()
		case <-time.After(time.Second):
			// wait for 1 second
		}
	}
	return nil
}
