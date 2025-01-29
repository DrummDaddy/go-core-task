package main

import (
	"sync"
)

// StringIntMap структура для хранения пар строка-целое число
type StringIntMap struct {
	mu    sync.RWMutex
	items map[string]int
}

// NewStringIntMap конструктор для StringIntMap
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		items: make(map[string]int),
	}
}

// Add - метод для добавления новой пары ключ-значение
func (s *StringIntMap) Add(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = value

}

// Remove - метод для удаления по ключу
func (s *StringIntMap) Remove(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, key)
}

// Copy метод для создания копии карты
func (s *StringIntMap) Copy() map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	copyMap := make(map[string]int)
	for key, value := range s.items {
		copyMap[key] = value
	}
	return copyMap

}

// Exists - метод для проверки ключа
func (s *StringIntMap) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.items[key]
	return exists
}

// Get - метод для получения значения по ключу

func (s *StringIntMap) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.items[key]
	return value, exists
}
