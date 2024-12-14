package main

import (
	"backend_web_service/internal/config"
	"backend_web_service/internal/database"
	"backend_web_service/internal/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// Загрузка переменных из .env
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить .env файл, используются переменные окружения")
	}

	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db, err := database.Connect(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Инициализация резолверов и сервера
	resolver := &graph.Resolver{DB: db}
	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Настройка маршрутов
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", server)

	// Запуск сервера
	log.Printf("Сервер запущен на http://localhost:%s/", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
