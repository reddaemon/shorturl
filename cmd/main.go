package main

// 1. post большую - сокращенную
// 2. get водим полученую сокращенную, редиректимся на большую (http codes)
// 3. 1 запрос в секунду
// сокращенную ссылка неодноразовая

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"shorturl/internal/db"
	"shorturl/internal/handlers"
	"shorturl/internal/repository"
	"shorturl/internal/router"
	"shorturl/internal/service"
	"shorturl/internal/shorturl"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	database := db.InitDB()
	defer func() {
		_ = database.Close()
	}()

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(database, "./migrations"); err != nil {
		panic(err)
	}

	repoTool := repository.NewRepoTool(database)

	svc := service.NewService(repoTool)
	shortener := &shorturl.Url{}
	h := handlers.NewHandler(svc, shortener)
	r := router.RegisterRouter(h)

	httpServer := &http.Server{
		Addr:        ":8080",
		Handler:     r,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	// Run server
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main goroutine
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

	gracefulCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := httpServer.Shutdown(gracefulCtx); err != nil {
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
