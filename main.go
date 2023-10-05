package main

import (
	"binary-tree/handlers"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int

func initConfig() {
	// Define the port flag
	flag.IntVar(&port, "port", 8080, "HTTP server port")
	// Parse the flags
	flag.Parse()
}

func main() {
	initConfig()
	mux := http.NewServeMux()
	// add  calculation endpoint
	mux.HandleFunc("/calculate", handlers.CalculationHandler)
	address := fmt.Sprintf(":%d", port)
	log.Printf("Server started, listening on port %d\n", port)
	panic(http.ListenAndServe(address, RecoveryMiddleware(mux)))
}

// RecoveryMiddleware is an HTTP middleware that recovers gracefully
// from panics in the HTTP handlers down the chain. Instead of crashing the server,
// it logs the error and responds with a generic "Internal server error" message.
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)

				// Send a generic server error message to the client
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
