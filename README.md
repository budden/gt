# Попытка наладить gomacro  

## Что оно делает?

Генерирует с помощью квазицитирования несколько вариантов функции вычисления чисел Фибоначчи для разных типов аргументов.

## Как запустить

Как-то так:

* клонировать репозиторий
* go mod init
* go mod tidy
* go get ./...
* go generate ./...
* gu run main.go

## А что там внутри?

Файл [app/make_fibonacci.gomacro](app/make_fibonacci.gomacro) во время `go generate` превращается в [app/make_fibonacci.go](app/make_fibonacci.go). Это происходит благодаря наличию требования `go:generate` в [app/run.go](app/run.go). 

