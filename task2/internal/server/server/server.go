package server

import (
	"context"
	"errors"
	"fmt"
	"hsse_go_homework/task2/internal/server/decode"
	"hsse_go_homework/task2/internal/server/hardop"
	"hsse_go_homework/task2/internal/server/version"
	"log"
	"net/http"
	"time"
)

const shutdownTimeout = 15 * time.Second

func NewOnPort(port string) http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /version", version.Handler)
	mux.HandleFunc("POST /decode", decode.Handler)
	mux.HandleFunc("GET /hard-op", hardop.Handler)

	return http.Server{
		Addr:    port,
		Handler: mux,
	}
}

func Start(server *http.Server) error {
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("Got error in Listen", err)
		return fmt.Errorf("failed to serve http server: %w", err)
	}
	fmt.Println("after listener")

	return nil
}

func Stop(ctx context.Context, server *http.Server) error {
	<-ctx.Done()
	log.Println("Context is done")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
		return err
	}
	log.Println("Server shutdown")

	return nil
}
