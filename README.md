# taskManagerApi

Простой pet-проект для управления задачами (todo-list) на Go с использованием PostgreSQL и gorm.

## Быстрый старт

### 1. Клонируй репозиторий и перейди в папку проекта

```
git clone https://github.com/montiquell/taskManagerApi.git
cd taskManagerApi
```

### 2. Создай файл `.env`

В корне проекта создай файл `.env` и пропиши строку подключения к базе:

```
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

### 3. Установи зависимости

```
go mod download
```

### 4. Применить миграции

Для работы с миграциями используется [golang-migrate](https://github.com/golang-migrate/migrate). Установи его, если не установлен:

- Windows: скачай бинарник с [релизов](https://github.com/golang-migrate/migrate/releases)
- Либо через пакетный менеджер (winget, choco, brew и т.д.)

Применить миграции:

```
make migrations-up
```

Откатить миграции:

```
make migrations-down
```

Миграции лежат в `internal/migrations/`.

### 5. Запусти сервер

```
go run ./cmd/api
```

Сервер стартует на `localhost:8080`.

## API

- `POST   /api/todos` — создать задачу
- `GET    /api/todos` — получить все задачи
- `PATCH  /api/todos/{id}` — обновить статус задачи
- `DELETE /api/todos/{id}` — удалить задачу

## Структура проекта

- `cmd/api/main.go` — точка входа
- `internal/adapter/repository/` — работа с БД
- `internal/domain/` — бизнес-модели и ошибки
- `internal/service/` — бизнес-логика
- `internal/transport/http/` — HTTP-роутинг и хендлеры
- `internal/migrations/` — миграции для БД
- `internal/config/` — конфиг

## Зависимости

- go-chi/chi — роутер
- gorm — ORM
- golang-migrate — миграции
- godotenv — переменные окружения

## Примечания

- Для работы нужен PostgreSQL.
- Все переменные окружения берутся из `.env`.
- Makefile содержит команды для миграций.

---

Если что-то не работает — смотри логи, они довольно подробные.
