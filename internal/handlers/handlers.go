package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"6_sprint/internal/service"
)

// IndexHandler обрабатывает запросы к корневому эндпоинту / и возвращает HTML-страницу.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Получаем полный путь к файлу index.html
	fullPath, err := filepath.Abs("./index.html")
	if err != nil {
		http.Error(w, "Ошибка определения пути к HTML", http.StatusInternalServerError)
		return
	}

	// Отправляем статический HTML-файл клиенту
	http.ServeFile(w, r, fullPath)
}

// UploadHandler обрабатывает загрузку файла через форму.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Ограничиваем размер загружаемого файла
	r.ParseMultipartForm(10 << 20) // 10 MB

	// Получаем файл из формы
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// Конвертируем данные с помощью функции из пакета service
	result, err := service.ConvertTextToMorseOrViceVersa(string(data))
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаём локальный файл для результата, делаю в windows
	outputFileName := fmt.Sprintf("result_%s%s", time.Now().UTC().Format("15-04-05_02.01.2006"), filepath.Ext(handler.Filename))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Записываем результат в файл
	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат конвертации
	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, "Ошибка отправки результата клиенту", http.StatusInternalServerError)
		return
	}
}
