package main

func main() {
	n := 40
	memo := make([]int, n)
	println(Fib2(n, memo))
	println(Fib1(n))
	println(Fib3(n, make([]int, n + 1)))
}

// Fib1 returns Fibonacci sequence 1 1 2 3 5 8 13 ...
// O(2^n)
func Fib1(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib1(n-1) + Fib1(n-2)
}

// Fib2
// Memoization
// T(n) = calls x time it takes <= 2n + 1 (first time) . O(1) = O(n)
// 1...n
func Fib2(n int, memo []int) int {
	if n <= 1 {
		return 1
	}

	if memo[n-1] != 0 {
		return memo[n-1]
	}

	res := Fib1(n-1) + Fib1(n-2)

	memo[n-1] = res

	return res
}

func Fib3(n int, up []int) int {
	if n <= 1 {
		return 1
	}

	up[0] = 1
	up[1] = 1

	//i=2: 1 1 2
	//i=3: 1 1 2 3
	//i=4: 1 1 2 3 ...
	for i := 2; i <= n; i++ {
		up[i] = up[i-1] + up[i-2]
	}

	return up[n]
}

