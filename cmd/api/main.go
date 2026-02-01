package main

import (
	"context"
	"log"
	"net/http"
	"newTaskManagerApi/internal/adapter/repository"
	"newTaskManagerApi/internal/config"
	"newTaskManagerApi/internal/service"
	router "newTaskManagerApi/internal/transport/http"
	handler "newTaskManagerApi/internal/transport/http/task"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()
	db, err := gorm.Open(postgres.Open(cfg.DATABASE_URL), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connecting error: %s", err.Error())
	}

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	r := router.NewRouter(taskHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("server starting...")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed to start: %s", err.Error())
		}
	}()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	log.Println("shutdowning...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("error shutdowning: %s", err)
	}

	log.Println("server closed")
}
