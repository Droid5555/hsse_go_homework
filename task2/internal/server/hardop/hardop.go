package hardop

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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
