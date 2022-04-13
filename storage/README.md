
# Получение 
```bash
go get -v github.com/asqat/mgid/storage
```

# Использование
```go
storage := getStorage(1 * time.Second) // Инициализируем время хранения 1 секунда

storage.Write("foo1", "bar1") // Пишем первое значение
storage.Write("foo2", "bar2") // Пишем второе значение

if val, ok := storage.Read("foo1"); !ok || val == nil { // Проверяем первое значение
t.Errorf("cannot found the wrote data with foo1 key")
}

if val, ok := storage.Read("foo2"); !ok || val == nil { // Проверяем второе значение
t.Errorf("cannot found the wrote data with foo2 key")
}
```

# Запуск теста
```bash
go test -v
```

# Запуск бенчмарка

```bash
go test -bench=BenchmarkMapStore
```

### Результат на Intel(R) Core(TM) i9-10900 CPU @ 2.80GHz:
```
# Длина ключа 18 символов
BenchmarkMapStore_Write_18-20             551745              2401 ns/op             304 B/op          7 allocs/op

# Длина ключа 1000 символов
BenchmarkMapStore_Write_1000-20            15544             83077 ns/op            5333 B/op          7 allocs/op

# Каждая итерация содержит 100 чтения
BenchmarkMapStore_Read_100-20            2456925               460.7 ns/op             0 B/op          0 allocs/op

# Каждая итерация содержит 10000 чтения
BenchmarkMapStore_Read_10000-20            24722             48765 ns/op               0 B/op          0 allocs/op

# Каждая итерация содержит 1000000 чтения
BenchmarkMapStore_Read_1000000-20            102          11613822 ns/op             100 B/op          0 allocs/op
```