package server

import (
	"log"
	"net/http"
	"time"

	"6_sprint/internal/handlers"
)

// Server структура сервера с логгером и HTTP-сервером.
type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

// NewServer создает и настраивает HTTP-сервер.
func NewServer(logger *log.Logger) *Server {
	// Создаем HTTP-роутер
	router := http.NewServeMux()

	// Регистрируем хендлеры
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	// Настраиваем HTTP-сервер
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Возвращаем экземпляр сервера
	return &Server{
		Logger: logger,
		HTTP:   httpServer,
	}
}
