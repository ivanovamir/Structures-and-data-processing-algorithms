package stack

import "fmt"

type stack struct {
	item []Bucket
	top  *Bucket
	cap  *int
}

type Bucket struct {
	val int
	nx  *int
}

type Stack interface {
	IsEmpty() bool
	IsFull() bool
	Push(item int) error
	Pop() int
	Statistic() *stack
}

func NewStack(cap int) Stack {
	return &stack{
		cap:  &cap,
		item: make([]Bucket, 0, cap),
	}
}

func (s *stack) IsFull() bool {
	// Если указ на вместимость не нулевой, то возвращаем результат сравнения длины стека и знаечния вместимости стека
	if !(s.cap == nil) {
		return len(s.item) == *s.cap
	}
	return true
}

func (s *stack) IsEmpty() bool {
	// Если указ на вместимость не нулевой, то возвращаем результат сравнения ссылки топ элемента стека и нулевого указателя
	if !(s.cap == nil) {
		return s.top == nil
	}
	return false
}

func (s *stack) Push(item int) error {
	// Проверка на заполненность стека
	if s.IsFull() {
		return fmt.Errorf("stack is full")
	}

	// Создаем нвоый бакет в памяти
	buck := Bucket{
		val: item,
		nx:  nil,
	}

	// Проверяем является ли указ топ элемент стека нулевым
	// Если нет - то указа на след элемент нового бакета равен указ на топ элемент
	if s.top != nil {
		buck.nx = &s.top.val
	}

	// Новый топ элемент - бакет нового элемента
	s.top = &buck

	// Добавляем бакет в стек
	s.item = append(s.item, buck)

	return nil
}

func (s *stack) Pop() int {

	// Проверка на пустоту стека
	if s.IsEmpty() {
		return 0
	}

	s.top = &s.item[len(s.item)-1]  // берем последний элемент слайса, теперь он - наш топ
	s.item = s.item[:len(s.item)-1] // берем срез с 0 до последнего элемента слайса, теперь это - наш стек
	return s.top.val
}

func (s *stack) Statistic() *stack {
	return s
}
