package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect(host, user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Успешное подключение к базе данных")

	// Создание таблиц, если они не существуют
	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	// SQL для создания таблиц
	createPostsTable := `
 CREATE TABLE IF NOT EXISTS posts (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  comments_enabled BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
 );`

	createCommentsTable := `
 CREATE TABLE IF NOT EXISTS comments (
  id SERIAL PRIMARY KEY,
  post_id INT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  text TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
 );`

	// Выполнение SQL-запросов
	if _, err := db.Exec(createPostsTable); err != nil {
		return fmt.Errorf("ошибка при создании таблицы posts: %v", err)
	}

	if _, err := db.Exec(createCommentsTable); err != nil {
		return fmt.Errorf("ошибка при создании таблицы comments: %v", err)
	}

	log.Println("Таблицы успешно созданы или уже существуют")
	return nil
}
