package storage

import (
	"sync"
	"sync/atomic"
	"time"
)

type MapStore struct {
	store    sync.Map
	lifetime time.Duration
	size     int64
	close    chan struct{}
}

type elem struct {
	value       interface{}
	index       int
	removeAfter int64
}

func InitMapStore(storeDuration time.Duration) *MapStore {
	sMap := &MapStore{
		lifetime: storeDuration,
		close:    make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				now := time.Now().UnixNano()

				sMap.store.Range(func(key, value interface{}) bool {
					elem := value.(elem)
					if elem.removeAfter > 0 && now > elem.removeAfter {
						sMap.store.Delete(key)
						tmp := atomic.AddInt64(&sMap.size, -1)
						sMap.size = tmp
					}
					return true
				})

			case <-sMap.close:
				return
			}
		}
	}()

	return sMap
}

func (ms *MapStore) Read(key string) (i int, value interface{}) {
	loaded, exists := ms.store.Load(key)
	if !exists {
		return
	}

	elem := loaded.(elem)

	if elem.removeAfter > 0 && time.Now().UnixNano() > elem.removeAfter {
		tmp := atomic.AddInt64(&ms.size, -1)
		ms.size = tmp
		return
	}

	return elem.index, elem.value
}

func (ms *MapStore) Size() int {
	return int(ms.size)
}

func (ms *MapStore) Write(key string, value interface{}) {
	var removeAfter int64

	if ms.lifetime > 0 {
		removeAfter = time.Now().Add(ms.lifetime).UnixNano()
	}

	ms.store.Store(key, elem{
		value:       value,
		removeAfter: removeAfter,
		index:       int(ms.size),
	})
	tmp := atomic.AddInt64(&ms.size, 1)
	ms.size = tmp
}

func (ms *MapStore) Close() {
	ms.size = 0
	ms.close <- struct{}{}
	ms.store = sync.Map{}
}
