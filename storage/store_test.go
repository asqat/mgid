package storage

import (
	"math/rand"
	"testing"
	"time"
)

func genAlphaNum(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func genRandomData(length int) []string {
	bucket := make([]string, length)
	for i := range bucket {
		bucket[i] = genAlphaNum(17)
	}
	return bucket
}

func getStorage(d time.Duration) *MapStore {
	return InitMapStore(d)
}

func TestMapStore_Write(t *testing.T) {
	storage := getStorage(2 * time.Second) // Инициализируем время хранения 1 секунда

	storage.Write("foo1", "bar1") // Пишем первое значение
	storage.Write("foo2", "bar2") // Пишем второе значение

	if i, value := storage.Read("foo1"); i != 0 || value != "bar1" {
		t.Errorf("cannot found the wrote data with foo1 key")
	}
	if i, value := storage.Read("foo2"); i != 1 || value != "bar2" {
		t.Errorf("cannot found the wrote data with foo1 key")
	}

	if size := storage.Size(); size != 2 {
		t.Errorf("Size() must be 2, not %d", size)
	}

	time.Sleep(3 * time.Second) // Ждем пока срок хранения данных не истечет

	if i, value := storage.Read("foo1"); i == 0 && value == "bar1" {
		t.Errorf("expired data found: foo1")
	}

	if i, value := storage.Read("foo2"); i == 1 && value == "bar2" {
		t.Errorf("expired data found: foo2")
	}

	if size := storage.Size(); size != 0 {
		t.Errorf("Size() must be 0, not %d", size)
	}
}

func TestMapStore_Inc(t *testing.T) {
	store := InitMapStore(3 * time.Second)
	defer store.Close()

	testData := map[string]string{
		"foo1": "bar1",
		"foo2": "bar2",
		"foo3": "bar3",
		"foo4": "bar4",
		"foo5": "bar5",
		"foo6": "bar6",
	}

	for k, v := range testData {
		store.Write(k, v)
	}

	tests := []struct {
		name      string
		store     *MapStore
		wantValue interface{}
	}{
		{
			name:      "IncTest",
			store:     store,
			wantValue: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := 0
			for _, _ = range testData {
				iter++
				value := store.Inc()
				if value == nil {
					t.Errorf("cannot found value on %d iteration", iter)
				}
			}
		})
	}
}

func BenchmarkMapStore_Write_18(bench *testing.B) {
	store := InitMapStore(3 * time.Second)
	defer store.Close()

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Write(genAlphaNum(18), struct{}{})
		}
	})
}

func BenchmarkMapStore_Write_1000(bench *testing.B) {
	store := InitMapStore(3 * time.Second)
	defer store.Close()

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Write(genAlphaNum(1000), struct{}{})
		}
	})
}

func BenchmarkMapStore_Read_100(bench *testing.B) {
	generatedData := genRandomData(100)

	store := InitMapStore(3 * time.Second)
	defer store.Close()

	for _, d := range generatedData {
		store.Write(d, struct{}{})
	}

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, k := range generatedData {
				store.Read(k)
			}
		}
	})
}

func BenchmarkMapStore_Read_10000(bench *testing.B) {
	generatedData := genRandomData(10000)

	store := InitMapStore(3 * time.Second)
	defer store.Close()

	for _, d := range generatedData {
		store.Write(d, struct{}{})
	}

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, k := range generatedData {
				store.Read(k)
			}
		}
	})
}

func BenchmarkMapStore_Read_1000000(bench *testing.B) {
	generatedData := genRandomData(1000000)

	store := InitMapStore(3 * time.Second)
	defer store.Close()

	for _, d := range generatedData {
		store.Write(d, struct{}{})
	}

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, k := range generatedData {
				store.Read(k)
			}
		}
	})
}

func BenchmarkMapStore_Inc_100(bench *testing.B) {
	generatedData := genRandomData(100)

	store := InitMapStore(3 * time.Second)
	defer store.Close()

	for _, d := range generatedData {
		store.Write(d, struct{}{})
	}

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Inc()
		}
	})
}

func BenchmarkMapStore_Inc_10000(bench *testing.B) {
	generatedData := genRandomData(10000)

	store := InitMapStore(3 * time.Second)
	defer store.Close()

	for _, d := range generatedData {
		store.Write(d, struct{}{})
	}

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Inc()
		}
	})
}

func BenchmarkMapStore_Inc_1000000(bench *testing.B) {
	generatedData := genRandomData(1000000)

	store := InitMapStore(3 * time.Second)
	defer store.Close()

	for _, d := range generatedData {
		store.Write(d, struct{}{})
	}

	bench.ReportAllocs()
	bench.ResetTimer()

	bench.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			store.Inc()
		}
	})
}
