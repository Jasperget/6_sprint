package main

import (
	"log"
	"os"

	"6_sprint/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)

	// Создаем сервер
	srv := server.NewServer(logger)

	// Запускаем сервер
	logger.Println("Сервер запущен на порту 8080")
	if err := srv.HTTP.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
