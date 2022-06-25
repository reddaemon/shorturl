package main

// 1. post большую - сокращенную
// 2. get водим полученую сокращенную, редиректимся на большую (http codes)
// 3. 1 запрос в секунду
// сокращенную ссылка неодноразовая

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"shorturl/internal/db"
	"shorturl/internal/handlers"
	"shorturl/internal/repository"
	"shorturl/internal/router"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	db := db.InitDB()
	defer db.Close()
	repo := repository.NewRepo(db)
	err := repo.CreateBucket("Urls")
	if err != nil {
		log.Printf("cannot create bucket: %v", err)
	}
	h := handlers.NewHandler(repo)
	r := router.RegisterRouter(h)

	httpServer := &http.Server{
		Addr:        ":8080",
		Handler:     r,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	// Run server
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")
	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("gracefully stopped\n")
	}

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)

}
