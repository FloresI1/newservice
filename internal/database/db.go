package database

import (
    "context"
    "log"

    "github.com/jackc/pgx/v4/pgxpool"
    
)

// InitDB инициализирует подключение к базе данных
func InitDB() (*pgxpool.Pool, error) {
    config, err := pgxpool.ParseConfig("postgresql://postgres:1@localhost:5432/news")
    if err != nil {
        log.Fatalf("Ошибка при разборе конфигурации: %v", err)
        return nil, err
    }

    config.MaxConns = 10 // устанавливаем максимальное количество соединений
    config.MinConns = 5  // устанавливаем минимальное количество соединений

    pool, err := pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        log.Fatalf("Ошибка при подключении к базе данных: %v", err)
        return nil, err
    }

    // Проверка соединения
    if err := pool.Ping(context.Background()); err != nil {
        log.Fatalf("Ошибка при проверке соединения с базой данных: %v", err)
        return nil, err
    }

    log.Println("Подключение к базе данных успешно")
    return pool, nil
}

// Другие функции для работы с базой данных, принимающие пул соединений...
