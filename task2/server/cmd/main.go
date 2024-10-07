package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const shutdownTimeout = 15 * time.Second

type DecodeInput struct {
	InputString string `json:"inputString"`
}

type DecodeOutput struct {
	OutputString string `json:"outputString"`
}

const (
	VersionMajor = 1
	VersionMinor = 21
	VersionPatch = 1
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	version := fmt.Sprintf("v%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
	_, err := w.Write([]byte(version))
	if err != nil {
		http.Error(w, "500 Internal Server Error : (Response Write Problem)", http.StatusInternalServerError)
		return
	}
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input DecodeInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	decodedString, err := base64.StdEncoding.DecodeString(input.InputString)
	if err != nil {
		http.Error(w, "400 Bad Request : (Decode Error)", http.StatusBadRequest)
		return
	}

	output := DecodeOutput{
		OutputString: string(decodedString),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, "500 Internal Server Error : (Json Encode Problem)", http.StatusInternalServerError)
		return
	}
}

func hardOpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	time.Sleep(time.Duration(10+rand.Intn(10_000)) * time.Millisecond)

	if rand.Intn(2) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		http.Error(w, "200 OK", http.StatusOK)
	}
}

func NewServerOnPort(port string) http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", versionHandler)
	mux.HandleFunc("/decode", decodeHandler)
	mux.HandleFunc("/hard-op", hardOpHandler)

	return http.Server{
		Addr:    port,
		Handler: mux,
	}
}

func StartServer(server *http.Server) error {
	log.Println("Starting server")
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("error in listen: %s\n", err)
		return fmt.Errorf("failed to serve http server: %w", err)
	}

	fmt.Println("after listener")
	return nil
}

func main() {
	httpServer := NewServerOnPort(":8080")

	// Took code below from seminar 1.6
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("err in listen: %s\n", err)
			return fmt.Errorf("failed to serve http server: %w", err)
		}
		fmt.Println("after listener")

		return nil
	})

	group.Go(func() error {
		fmt.Println("before ctx done")
		<-ctx.Done()
		fmt.Println("after ctx done")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
			return err
		}
		log.Println("Server shutdown")

		return nil
	})

	// Group errors handling
	err := group.Wait()
	if err != nil {
		log.Println("Got non-nil error from group\n", err)
		return
	}

}

/*
TODO:
	- clear main func
	- put version in json
	- add comments
	- divide code into packages
*/
