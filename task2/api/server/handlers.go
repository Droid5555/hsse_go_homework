package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hsse_go_homework/task2/tools/decode_tools"
	"hsse_go_homework/task2/tools/version_tools"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const shutdownTimeout = 15 * time.Second

type Server struct {
	http.Server
}

func NewOnPort(port string) Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /version", GetVersion)
	mux.HandleFunc("POST /decode", Decode)
	mux.HandleFunc("GET /hard-op", HardOp)

	return Server{http.Server{
		Addr:    port,
		Handler: mux,
	}}
}

func (server *Server) Start() error {
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("Got error in Listen", err)
		return fmt.Errorf("failed to serve http server: %w", err)
	}
	fmt.Println("after listener")

	return nil
}

func (server *Server) Stop(ctx context.Context) error {
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

func Decode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input decode_tools.Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	decodedString, err := base64.StdEncoding.DecodeString(input.Request)
	if err != nil {
		http.Error(w, "400 Bad Request : (Decode Error)", http.StatusBadRequest)
		return
	}

	output := decode_tools.Output{
		Response: string(decodedString),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, "500 Internal Server Error : (Json Encode Problem)", http.StatusInternalServerError)
		return
	}
}

func HardOp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	responseTime := time.Duration(10_000+rand.Intn(10_000)) * time.Millisecond
	log.Println("RESPONSE TIME:", responseTime)
	time.Sleep(responseTime)

	if rand.Intn(2) == 0 {
		randErr := 500 + rand.Intn(10)
		_, err := fmt.Fprintln(w, randErr, http.StatusText(randErr))
		if err != nil {
			return
		}
	} else {
		_, err := fmt.Fprintln(w, http.StatusOK, "OK")
		if err != nil {
			return
		}
	}
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	version := fmt.Sprintf("v%d.%d.%d", version_tools.VERSION.Major, version_tools.VERSION.Minor, version_tools.VERSION.Patch)
	_, err := w.Write([]byte(version))
	if err != nil {
		http.Error(w, "500 Internal Server Error : (Response Write Problem)", http.StatusInternalServerError)
		return
	}
}
