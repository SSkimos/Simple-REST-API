package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//viper.SetConfigFile("./pkg/common/envs/.env")
	//viper.ReadInConfig()
	//
	//publisherAPI := viper.Get("PUBLISHER").(string)
	//retrievalAPI := viper.Get("RETRIEVAL").(string)

	// Создаем клиент с кастомным транспортом без поддержки перенаправлений
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true, // Отключаем keep-alive, чтобы избежать перенаправлений
		},
	}

	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodGet && r.Method != http.MethodPost {
			http.Error(w, "Метод запроса не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Формируем URL целевого API
		targetURL := ""
		if r.Method == http.MethodGet {
			targetURL += "http://retrieval:80"
			targetURL += r.URL.Path
			log.Printf("target: %w", targetURL)

			newReq, err := http.NewRequest(http.MethodGet, targetURL, strings.NewReader(string(body)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			newReq.Header = make(http.Header)
			copyHeaders(newReq.Header, r.Header)
			r = newReq
		}

		if r.Method == http.MethodPost {
			targetURL += "http://publisher:8080"
			targetURL += r.URL.Path

			// Создаем новый запрос с передачей тела
			newReq, err := http.NewRequest(http.MethodPost, targetURL, strings.NewReader(string(body)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			newReq.Header = make(http.Header)
			copyHeaders(newReq.Header, r.Header)
			r = newReq
		}

		// Выполняем запрос к целевому API
		targetResp, err := client.Do(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetResp.Body.Close()

		// Копируем заголовки ответа от целевого API
		copyHeaders(w.Header(), targetResp.Header)

		// Отправляем статус код и тело ответа клиенту
		w.WriteHeader(targetResp.StatusCode)
		body, err = ioutil.ReadAll(targetResp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(body)
	})

	// Запускаем сервер на порту 8000
	fmt.Println("Прокси-сервер запущен на порту 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("Ошибка при запуске прокси-сервера: %s\n", err)
	}
}

// Функция для копирования заголовков из одной структуры в другую
func copyHeaders(dest, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dest.Add(key, value)
		}
	}
}
