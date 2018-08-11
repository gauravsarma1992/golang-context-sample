package main

import (
	"context"
	"fmt"
	"time"
)

func Args() {
	ctx, cncl := context.WithTimeout(context.Background(),
		time.Hour)

	go func(ctx context.Context) {
		GrandparentFunc(ctx)
	}(ctx)

	time.Sleep(4 * time.Second)
	cncl()
}

func GrandparentFunc(ctx context.Context) {

	go ParentFunc(ctx)

	for {
		fmt.Println("In Grandparent Func")

		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(time.Second):
			// wait for 1 second
		}
	}
}

func ParentFunc(ctx context.Context) {
	go ChildFunc(ctx)
	for {
		fmt.Println("In Parent Func")

		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(time.Second):
			// wait for 1 second
		}
	}
}

func ChildFunc(ctx context.Context) {
	for {
		fmt.Println("In Child Func")

		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(time.Second):
			// wait for 1 second
		}
	}
}
