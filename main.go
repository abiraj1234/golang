package main

func prime(n int) {

	if n < 1 {
		println("Enter the positive number")

	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			println("not prime number")
			return
		}

	}
	println("prime number")
}

func main() {
	prime(2)
	prime(11)
}
