# Инициализируем библиотеки
```bash
go mod tidy
```

# Запуск сервера
```bash
go run main.go
```

Сервис запутится на 5600 порту

# Запуск юнит тестов
```bash
go test -v ./pkg/server
```

# Кодогенерация из proto-файла
```bash
go generate
```