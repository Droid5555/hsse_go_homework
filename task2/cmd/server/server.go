package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"hsse_go_homework/task2/internal/server/server"
	"hsse_go_homework/task2/tools/version_tools"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	httpServer := server.NewOnPort(":8080")

	err := version_tools.LoadFromJson("task2/tools/version.json")
	if err != nil {
		log.Fatalf("Error loading version from JSON: %v", err)
	}

	//Took code below from seminar 1.6
	// Create shutdown context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create a group for this context
	group, ctx := errgroup.WithContext(ctx)

	// Start server handling
	group.Go(func() error {
		return server.Start(&httpServer)
	})

	// Shutdown server handling
	group.Go(func() error {
		return server.Stop(ctx, &httpServer)
	})

	// Group errors handling
	err = group.Wait()
	if err != nil {
		log.Println("Got non-nil error from groups\n", err)
		return
	}

}
