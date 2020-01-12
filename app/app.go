package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/lucasstettner/api-boilerplate/config"
)

// Defines the structure of whole app
type App struct {
	Router *chi.Mux
	Config *config.Config
}

// Begins the server
func (a *App) Start(gracefully bool) {
	a.Config = config.New()

	defer a.Config.Logger.Sync() // flushes buffer, if any

	// Create new chi mux and setup middlware/routes
	a.Router = Routes(a.Config)

	// Print out all routes
	if err := chi.Walk(a.Router, walkFunc); err != nil {
		log.Panicf("Logging err: %v\n", err.Error())
	}

	h := &http.Server{
		Handler:      a.Router,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Concurrently start http server process
	go func() {
		log.Println("Starting Server")

		// Shut down db connection when concurrent go process is closed
		defer a.Config.DB.Close()

		if err := h.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	if gracefully {
		waitForShutdown(h)
	}
}

// Gracefully shut down server
func waitForShutdown(s *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive signal
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Panic("Error shutting down")
	}

	log.Println("Shutting down")
	os.Exit(0)
}

// Shows all routes on start
func walkFunc(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	// Walk and print out all routes
	log.Printf("%s %s\n", method, route)
	return nil
}
