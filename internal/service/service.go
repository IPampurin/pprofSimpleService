package service

import (
	"fmt"
)

type Service struct{}

func NewService() *Service {

	return &Service{}
}

// Sum реализует сложение
func (s *Service) Sum(a, b int64) int64 {

	return a + b
}

// Fib реализует рекурсивное вычисление чисел Фибоначчи (для n == 0 возвращает 0)
func (s *Service) Fib(n int) int64 {

	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return s.Fib(n-1) + s.Fib(n-2)
}

// Allocate выделяет память и сразу освобождает
func (s *Service) Allocate(size int) (int, error) {

	if size <= 0 {
		return 0, fmt.Errorf("размер должен быть положительным: %d", size)
	}
	// выделяем слайс, чтобы нагрузить GC
	_ = make([]byte, size)

	return size, nil
}
