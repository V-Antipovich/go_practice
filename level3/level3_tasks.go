package main

import (
	"fmt"
)

func merge(a, b []int) []int {
	merged := []int{}
	idx1, idx2 := 0, 0
	for idx1 < len(a) && idx2 < len(b) {
		if a[idx1] < b[idx2] {
			merged = append(merged, a[idx1])
			idx1 += 1
		} else {
			merged = append(merged, b[idx2])
			idx2 += 1
		}
	}
	for ; idx1 < len(a); idx1++ {
		merged = append(merged, a[idx1])
	}
	for ; idx2 < len(b); idx2++ {
		merged = append(merged, b[idx2])
	}
	return merged
}

type Stack struct {
	arr []int
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Clear() {
	s.arr = nil
}

func (s *Stack) Push(n int) {
	s.arr = append(s.arr, n)
}

func (s *Stack) Back() int {
	return s.arr[len(s.arr)-1]
}

func (s *Stack) Pop() int {
	if len(s.arr) > 0 {
		last := s.Back()
		s.arr = s.arr[:len(s.arr)-1]
		return last
	}
	return int((^uint(0)) >> 1)
}

type Queue struct {
	s1, s2 Stack
}

func (q *Queue) Size() int {
	return q.s1.Size() + q.s2.Size()
}

func (q *Queue) Clear() {
	q.s1.Clear()
	q.s2.Clear()
}

func (q *Queue) Push(n int) {
	q.s1.Push(n)
}

func (q *Queue) move() {
	if q.s2.Size() == 0 {
		for q.s1.Size() > 0 {
			q.s2.Push(q.s1.Pop())
		}
	}
}

// Если очередь пуста, то вызов функций приведет к ошибке
func (q *Queue) Front() int {
	q.move()
	return q.s2.Back()
}

func (q *Queue) Pop() int {
	q.move()
	last := q.s2.Pop()
	return last
}

func main() {
	fmt.Println("27. Слияние отсортированных массивов")
	a := []int{5, 115}
	b := []int{2, 3, 4, 5, 6, 7, 7, 8, 9}
	c := merge(a, b)
	fmt.Println(c)

	fmt.Println("28. Хэш-таблица с коллизиями")

	fmt.Println("30. Очередь на основе двух стеков")
	q := Queue{
		s1: Stack{
			arr: []int{},
		},
		s2: Stack{
			arr: []int{},
		},
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	for q.Size() > 0 {
		fmt.Println(q.Pop(), q.Size())
	}
	fmt.Println("dsf")

}
