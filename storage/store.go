package storage

import (
	"sync"
	"time"
)

type MapStore struct {
	store    sync.Map
	lifetime time.Duration
	close    chan struct{}
}

type elem struct {
	value       interface{}
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

func (ms *MapStore) Read(key string) (value interface{}, ok bool) {
	loaded, exists := ms.store.Load(key)
	if !exists {
		return
	}

	elem := loaded.(elem)

	if elem.removeAfter > 0 && time.Now().UnixNano() > elem.removeAfter {
		return
	}

	return elem.value, true
}

func (ms *MapStore) Write(key string, value interface{}) {
	var removeAfter int64

	if ms.lifetime > 0 {
		removeAfter = time.Now().Add(ms.lifetime).UnixNano()
	}

	ms.store.Store(key, elem{
		value:       value,
		removeAfter: removeAfter,
	})
}

func (ms *MapStore) Delete(key string) {
	ms.store.Delete(key)
}

func (ms *MapStore) Close() {
	ms.close <- struct{}{}
	ms.store = sync.Map{}
}
