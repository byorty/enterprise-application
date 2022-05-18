# Enterprise Application

Пример микросервисного приложения на GRPC и DDD.

## Зависимости

1. [protoc](https://grpc.io/docs/protoc-installation/)
2. [docker](https://www.docker.com)
3. [golang](https://go.dev)
4. [buf](http://buf.build)
5. [make](https://ru.wikipedia.org/wiki/Make)

## Установка

```shell
make install      # установка плагинов для сборки proto
make vendor       # скачивания зависимостей для проекта
make buf_update   # скачивание зависимостей для сборки proto
make buf_generate # генерация golang-файлов из proto-файлов
```

## Использование

```shell
make up   # запуск микросервисов
make down # остановка микросервисов
```