package safemap

import (
	"fmt"
	"sync"
)

type SafeMap[K comparable, V any] struct{
	mu sync.RWMutex
	data map[K]V
}

func New[K comparable, V any]() *SafeMap[K, V]{
	return &SafeMap[K, V]{
		data: map[K]V{},
	}
}

func (sm *SafeMap[K, V]) Insert(key K, value V)(error){
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.data[key] = value
	return nil
}

func (sm *SafeMap[K, V]) Get(key K) (V, error){
	sm.mu.Lock()
	defer sm.mu.Unlock()

	value, ok := sm.data[key]
	if !ok {
		return value, fmt.Errorf("cannot found any key of name %v",key)
	}

	return value, nil
}

func (sm *SafeMap[K, V]) Update(key K, value V)(error){
	sm.mu.Lock()
	defer sm.mu.Unlock()


	_, ok := sm.data[key]
	if !ok {
		return fmt.Errorf("key %v not found for updation",key)
	}

	sm.data[key] = value
	return nil
}

func (sm *SafeMap[K, V]) Remove(key K)(error){
	sm.mu.Lock()
	defer sm.mu.Unlock()


	_, ok := sm.data[key]
	if !ok {
		return fmt.Errorf("key %v not found for deletion",key)
	}

	delete(sm.data, key)
	return nil
}

func (sm *SafeMap[K, V]) HasKey(key K)(bool){
	sm.mu.Lock()
	defer sm.mu.Unlock()


	_, ok := sm.data[key]

	return ok
}

func Hel() {
	fmt.Println("sdcwc")
}
