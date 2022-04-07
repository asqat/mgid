package storage

import (
	"testing"
	"time"
)

func getStorage(d time.Duration) *MapStore {
	return InitMapStore(d)
}

func TestMapStore_Write(t *testing.T) {
	storage := getStorage(1 * time.Second) // Инициализируем время хранения 1 секунда

	storage.Write("foo1", "bar1") // Пишем первое значение
	storage.Write("foo2", "bar2") // Пишем второе значение

	if val, ok := storage.Read("foo1"); !ok || val == nil { // Проверяем первое значение
		t.Errorf("cannot found the wrote data with foo1 key")
	}

	if val, ok := storage.Read("foo2"); !ok || val == nil { // Проверяем второе значение
		t.Errorf("cannot found the wrote data with foo2 key")
	}

	time.Sleep(2 * time.Second) // Ждем пока срок хранения данных не истечет

	if val, ok := storage.Read("foo1"); ok || val != nil { // Проверяем первое значение. Убедимся, что его больше нет.
		t.Errorf("expired data found: foo1")
	}

	if val, ok := storage.Read("foo2"); ok || val != nil { // Проверяем второе значение. Убедимся, что его больше нет.
		t.Errorf("expired data found: foo1")
	}
}

func BenchmarkMapStore_Write(bench *testing.B) {
	store := InitMapStore(3 * time.Second)
	defer store.Close()

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Write("Foo", "Bar")
		}
	})
}

func BenchmarkMapStore_Read(bench *testing.B) {
	store := InitMapStore(3 * time.Second)
	defer store.Close()

	store.Write("Foo", "Bar")

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Read("Foo")
		}
	})
}
