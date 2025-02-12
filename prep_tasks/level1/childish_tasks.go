package main

import (
	"fmt"
	"strings"
)

func helloWorld() {
	fmt.Println("Привет, мир!")
}
func add(a, b int) int {
	return a + b
}

func oddOrEven(n int) string {
	s := ""
	if n%2 == 0 {
		s = "четное"
	} else {
		s = "нечетное"
	}
	return fmt.Sprintf("Число %d %s", n, s)
}

func threeMax(a, b, c int) int {
	return max(a, b, c)
}

func factorial(n int) int {
	var fact int = 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	return fact
}

func isVowel(r string) bool {
	var checkString string = "aeiuoy"
	return len(r) == 1 && strings.Contains(checkString, r)
}

func getPrimes(n int) []int {
	var j int
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}
	var primes []int
	for i := 2; i*i <= n; i++ {
		if sieve[i] {
			j = i * i
			for ; j <= n; j += i {
				sieve[j] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func reverse(s string) string {
	rs := []rune(s)
	var rev []rune
	for i := len(rs) - 1; i >= 0; i-- {
		rev = append(rev, rs[i])
	}
	return string(rev)
}

func arrSum(arr []int) int {
	var ans int
	for _, v := range arr {
		ans += v
	}
	return ans
}

type Rectangle struct {
	length, height int
}

func (r Rectangle) getArea() int {
	return r.height * r.length
}

func main() {
	var a, b, c, n int = 100, 200, 500, 6
	fmt.Println("1. Привет, мир!")
	helloWorld()

	fmt.Println("2. Сложение чисел")
	fmt.Printf("%d+%d = %d\n", a, b, add(a, b))

	fmt.Println("3. Четное или нечетное")
	fmt.Println(oddOrEven(5))

	fmt.Println("4. Максимум из трех чисел")
	fmt.Println(threeMax(a, b, c))

	fmt.Println("5. Факториал числа")
	fmt.Printf("%d! = %d\n", n, factorial(n))

	fmt.Println("6. Проверка символа")
	var r string = "a"
	fmt.Printf("Является ли '%s' гласной (eng): %t\n", r, isVowel(r))

	fmt.Println("7. Простые числа")
	n = 101
	primes := getPrimes(n)
	fmt.Println("Простые до числа", n, ":", primes)

	fmt.Println("8. Строка и ее перевертыш")
	fmt.Println(reverse("8. Строка и ее перевертыш"))

	fmt.Println("9. Массив и его сумма")
	arr := []int{1, 6, 3, 4}
	fmt.Println(arrSum(arr))

	fmt.Println("10. Структуры и методы")
	rec := Rectangle{5, 5}
	fmt.Println(rec.getArea())
}
