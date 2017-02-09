package main

import "fmt"

func recursion() {
	println("-")
	recursion()
}

func Factorial(x int) (result int) {
	if x == 0{
		result = 1
	} else {
		result = x * Factorial(x - 1)
	}
	return ;
}

func fibonaci(n int) int {
	if n < 2 {
		return n
	}
	return fibonaci(n-2) + fibonaci(n-1)
}

func typeChange()  {
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为：%f\n",mean)
}

type Phone interface {
	call()
}
type NokiaPhone struct {
}
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I'm Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new (IPhone)
	phone.call()
}
