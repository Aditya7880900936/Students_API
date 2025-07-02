package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aditya7880900936/Students_API/internal/config"
)

func main() {
	// fmt.Println("Welcome to Students Api")
	// load config
	cfg := config.MustLoad()

	// databse setup
	// router setup

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Students API"))
	})

	// setup server

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server is Running at the Port :", cfg.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("Failed to Start the Server", slog.String("error", err.Error()))
		}
	}()

	<-done

	slog.Info("Shutting Down the Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to Shutdown the Server", slog.String("error", err.Error()))
	}

	slog.Info("Server Shutdown Successfully")

}
