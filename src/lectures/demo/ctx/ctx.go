package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case <-time.After(msDelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation completed", msg)
				return
			case <-ctx.Done():
				out <- fmt.Sprintf("%v aborted", msg)
				return
			}
		}
	}()
	return out
}

func main() {
	ctx := context.Background()
	// cancelCtx can cancel anything using ctx
	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := sampleOperation(ctx, "webServer", 200)
	microService := sampleOperation(ctx, "microservice", 500)
	database := sampleOperation(ctx, "database", 900)

	MainLoop:
	for {
		select {
		case val := <-webServer:
			fmt.Println(val)
		case val := <-microService:
			fmt.Println(val)
			fmt.Println("Cancel context")
			cancelCtx()
			break MainLoop
		case val := <-database:
			fmt.Println(val)
		}
	}

	fmt.Println(<-database)
}
