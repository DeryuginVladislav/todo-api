# TODO REST API

API для управления задачами (TODO-лист) на Go + Fiber + PostgreSQL.

## Запуск локально

1. Создать файл `.env` в корне с переменными:
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=
DB_NAME=todo_db
PORT=8080
2. Создать базу в PostgreSQL:
psql -h $DB_HOST -U $DB_USER -c "CREATE DATABASE $DB_NAME;
3. Запустить миграцию:
psql -h $DB_HOST -U $DB_USER -d $DB_NAME -f migrations/create_tasks_table.up.sql
4. Установить зависимости и запустить:
go mod tidy
go run main.go
