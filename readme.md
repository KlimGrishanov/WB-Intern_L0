# WB Internship L0

Данный проект разработан в рамках прохождения стажировки WildBerries L0

### Структура проекта

WB_Intern_L0

```
cmd -> main.go // Entry Point to Server
configs // Service configs
entity // Required go struct entities
internal // handler, usecase, repo layers
pkg // Additional Service (Cache, NATS)
schema // PostgreSQL Schema
loadtest // Result of loadtest
server.go // Basic of Server
```

## Установка и настройка

### Установка и настройка БД

Схема БД находится в /schema, для быстрой развертки проекта, я использовал систему make. Настройку команд вы можете посмотреть в makefile.
PostgreSQL работает в рамках Docker-контейнера.

```cmd
> make db-create # Создание БД
> make migrations-up # Загрузка миграций
```

### Запуск сервера

Чтобы запустить сервис достаточно выполнить следующую команду.

```cmd
> make run-server
```

