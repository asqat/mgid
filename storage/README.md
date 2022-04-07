
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
BenchmarkMapStore_Write-20       3867932               307.6 ns/op            56 B/op          3 allocs/op
BenchmarkMapStore_Read-20       294730365                3.984 ns/op           0 B/op          0 allocs/op
```