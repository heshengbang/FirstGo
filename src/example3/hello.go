package main

import "fmt"

func main() {
	//var a int = 21
	//var b int = 10
	//var c int
	//c = a + b
	//fmt.Printf("%d\n", c)
	//c = a - b
	//fmt.Printf("%d\n",c)
	//c = a * b
	//fmt.Printf("%d\n",c)
	//c = a / b
	//fmt.Printf("%d\n",c)
	//c = a % b
	//fmt.Printf("%d\n",c)
	//a++
	//fmt.Printf("%d\n",a)
	//a--
	//fmt.Printf("%d\n",a)
	println("----------------------------------------")
	var a int = 4
	var b int32
	var c float32
	var ptr *int
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)
	ptr = &a
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",*ptr)
	fmt.Println(*ptr)
}
