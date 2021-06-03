# Демо для курса по Docker 2
- Сборка на golang:alpine
- Запуск на scratch из собранного bin

Сборка следующей командой:
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```