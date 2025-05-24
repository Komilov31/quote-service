# Мини-сервис «Цитатник» на Go

Простой REST API-сервис для хранения и управления цитатами. Позволяет добавлять, получать, фильтровать, получать случайную цитату и удалять цитаты по ID.

---

## Функциональность

- **Добавление новой цитаты** — `POST /quotes`  
- **Получение всех цитат** — `GET /quotes`  
- **Получение случайной цитаты** — `GET /quotes/random`  
- **Фильтрация цитат по автору** — `GET /quotes?author=AuthorName`  
- **Удаление цитаты по ID** — `DELETE /quotes/{id}`  

---

## Технические детали

- Язык: Go  
- Хранение данных: в памяти (без базы данных)  
- Используемые библиотеки: стандартная библиотека Go и `gorilla/mux` для маршрутизации  
- Проект содержит README с инструкцией запуска  

---

## Запуск проекта

1. Клонируйте репозиторий:

```
git clone https://github.com/Komilov31/quote-service.git
```


2. Перейдите в директорию проекта:

```
cd quote-service
```


3. Запустите сервис:
```
go run cmd/main.go
```

Сервис будет доступен по адресу: `http://localhost:8080`

---

## Примеры запросов (curl)

- Добавить новую цитату:
```
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

- Получить все цитаты:
```
curl http://localhost:8080/quotes
```

- Получить случайную цитату:
```
curl http://localhost:8080/quotes/random
```

- Получить цитаты по автору:
```
curl http://localhost:8080/quotes?author=Confucius
```

- Удалить цитату по ID:
```
curl -X DELETE http://localhost:8080/quotes/1
```

---

## Структура проекта

- `cmd/main.go` — точка входа, запуск HTTP-сервера  
- `internal/` — основная логика приложения (обработка запросов, хранение цитат)   
- `README.md` — инструкция по запуску и использованию  

---

## Тестирование

В проекте написаны тесты для проверки хэндлеров. Запуск тестов:
```
go test internal/handlers/handlers_test.go
```
