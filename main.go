package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ryu-seunghyun/Todo-List/config/database"
	"github.com/Ryu-seunghyun/Todo-List/controllers"
	"github.com/Ryu-seunghyun/Todo-List/repository"
	"github.com/Ryu-seunghyun/Todo-List/server"
	"github.com/Ryu-seunghyun/Todo-List/service"
)

func main() {
	config := database.Database{}
	db := database.GetConnection(config)
	database.AutoMigrate()

	repos := repository.NewRepositories(db)

	service := service.NewServices(*repos)
	handler := controllers.NewHandler(service)
	svr := server.NewServer(handler.NewRouter())

	// 싱글톤 패턴과 연관.  프로그램에서 단 하나의 인스턴스만 생성하여 전역적으로 호출(사용).

	go func() {
		if err := svr.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	if err := svr.Stop(ctx); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}
}
