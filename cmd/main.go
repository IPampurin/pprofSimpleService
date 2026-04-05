package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IPampurin/pprofSimpleService/internal/configuration"
	"github.com/IPampurin/pprofSimpleService/internal/server"
	"github.com/IPampurin/pprofSimpleService/internal/service"
)

func main() {

	// считываем конфигурацию
	cfg, err := configuration.Load()
	if err != nil {
		log.Fatalf("ошибка конфигурации: %v", err)
	}

	// заводим конрневой контекст
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// запускаем ожидание сигналов
	go signalHandler(ctx, cancel)

	// получаем экземпляр логики
	svc := service.NewService()

	// получаем сервер
	srv := server.NewServer(&cfg, svc)

	// запускаем сервер
	if err := srv.Run(ctx); err != nil {
		log.Printf("сервер остановлен: %v", err)
	}
}

// signalHandler слушает SIGINT/SIGTERM и отменяет контекст
func signalHandler(ctx context.Context, cancel context.CancelFunc) {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigChan)

	select {
	case <-ctx.Done():
		return
	case <-sigChan:
		cancel()
		return
	}
}
