package main

import (
	"fmt"
	"strings"
)

func RemoveDuplicates(a []int) []int {
	were := make(map[int]bool)
	b := []int{}
	for _, v := range a {
		if !were[v] {
			b = append(b, v)
		}
		were[v] = true
	}
	return b
}

func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func Fib(n int) {
	f0, f1, f2 := 1, 1, 0
	if n >= 1 {
		fmt.Println(f0)
	}
	if n >= 2 {
		fmt.Println(f1)
	}
	for i := 2; i < n; i++ {
		f2 = f0 + f1
		fmt.Println(f2)
		f0, f1 = f1, f2
	}
}

func CountOccurrencies(a []int, n int) int {
	cnt := 0
	for _, v := range a {
		if v == n {
			cnt += 1
		}
	}
	return cnt
}

func ArrayIntersection(a, b []int) []int {
	were := make(map[int]int)
	c := []int{}
	for _, v := range a {
		were[v] += 1
	}
	for _, v := range b {
		if were[v] > 0 {
			c = append(c, v)
			were[v] -= 1
		}
	}
	return c
}

func AreAnagrams(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[len(b)-1-i] {
			return false
		}
	}
	return true
}

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

const (
	mod = 1e9 + 7
	x   = 314
)

func Pow(a, n, m int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		b := Pow(a, n/2, m) % m
		return (b * b) % m
	} else {
		return (Pow(a, n-1, m) % m * a % m) % m
	}
}

func h(s *string) int {
	hsh := 0
	for i, c := range *s {
		hsh = (hsh%mod + ((int(c)%mod)*(Pow(x, i, mod)%mod))%mod) % mod
	}
	return hsh
}

type pair struct {
	s string
	n int
}

type HashMap struct {
	table [][]pair
}

func (hm *HashMap) Insert(key string, val int) {
	hval := h(&key)
	// hm.table[hval] = append(hm.table[hval], pair{s: key, n: val})
	for i, v := range hm.table[hval] {
		if v.s == key {
			hm.table[hval][i].n = val
			return
		}
	}
	hm.table[hval] = append(hm.table[hval], pair{s: key, n: val})
}

// 0 is defaul value for int
func (hm *HashMap) Get(key string) int {
	hval := h(&key)
	val := 0
	for _, v := range hm.table[hval] {
		if v.s == key {
			val = v.n
			break
		}
	}
	return val
}

func (hm *HashMap) Delete(key string) {
	hval := h(&key)
	for i, v := range hm.table[hval] {
		if v.s == key {
			hm.table[hval] = append(hm.table[hval][:i], hm.table[hval][i+1:]...)
			break
		}
	}
}

func make_hashmap(m int) HashMap {
	return HashMap{table: make([][]pair, m)}
}

func BinSearch(a []int, n int) int {
	L, R, m := 0, len(a), 0
	for L+1 < R {
		m = (L + R) / 2
		if a[m] > n {
			R = m
		} else {
			L = m
		}
	}
	if a[L] == n {
		return L
	} else {
		return -1
	}
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
	arr := []int{4, 6, 1, 8, 3, 2, 1, 4, 3, 4, 4, 10, 12, 35, -1}
	fmt.Println("21. Удаление дубликатов")
	fmt.Println(RemoveDuplicates(arr))

	fmt.Println("22. Сортировка пузырьком")
	fmt.Println(BubbleSort(arr))

	fmt.Println("23. Фибоначчиева последовательность")
	Fib(10)

	fmt.Println("24. Количество вхождений элемента в массив")
	fmt.Println(CountOccurrencies(arr, -2))

	fmt.Println("25. Пересечение двух массивов")
	arr1 := []int{4, 6, 1, 8, 3, 2, 4, 3, 4, 4, 10, 12, 35}
	fmt.Println(ArrayIntersection(arr, arr1))

	fmt.Println("26. Анаграмма")
	fmt.Println(AreAnagrams("asSd)@", "@)dPssa"))

	fmt.Println("27. Слияние отсортированных массивов")
	a := []int{5, 115}
	b := []int{2, 3, 4, 5, 6, 7, 7, 8, 9}
	c := merge(a, b)
	fmt.Println(c)

	fmt.Println("28. Хэш-таблица с коллизиями")
	// Разрешение цепочками
	ht := make_hashmap(mod) //
	ht.Insert("asb", 5)
	ht.Insert("adr", 4)
	ht.Insert("aeq", 3)
	fmt.Println(ht.Get("ade"))
	fmt.Println(ht.Get("adr"))
	ht.Delete("asb")
	fmt.Println(ht.Get("asb"))

	fmt.Println("29. Бинарный поиск")
	// -1 если нет, иначе индекс элемента
	arr2 := []int{1, 1, 2, 3, 5, 6, 6, 6, 7, 8, 11, 12, 13}
	fmt.Println(BinSearch(arr2, 14))

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
}
