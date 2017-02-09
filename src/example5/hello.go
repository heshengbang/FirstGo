package main

import (
	"fmt"
)

func max(num1, num2 int) int {
	var result int
	if(num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x,y string) (string,string) {
	return y,x
}

func getSequence() func() int {
	i := 0
	return func() int {
		i+=1
		return i
	}
}

type Circle struct {
	radius float64
}
func strP(){
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ",c1.getArea())
}
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func arrayP () {
	var balance = [5]float32{10.2,20.3,30.4,40.5}
	for i:=0; i<5; i++ {
		fmt.Print(balance[i])
		println(i)
	}
}

func pointP () {
	var a int = 20
	var ip *int

	ip = &a
	fmt.Printf("a变量的地址是：%x\n", &a)
	fmt.Printf("ip变量的存储地址：%x\n", ip)
	fmt.Printf("*ip 变量的值：%d\n", *ip)
}

func arrayp2()  {
	var n [10] int
	var i,j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j<10; j++ {
		fmt.Printf("%d - %d\n",j,n[j])
	}
}
func arrayp3() {
	var a = [5][2] int {{0,0}, {1,2}, {2,4},{3,6},{4,8}}
	var i,j int
	for i = 0; i < 5 ; i++ {
		for j = 0 ; j < 2 ; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func pointP2() {
	var ptr *int
	fmt.Printf("%d\n",*ptr)
}

func pointP3() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000
	ptr = &a
	pptr = &ptr

	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *pstr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

	fmt.Printf("指针变量 a = %x\n", &a)
	fmt.Printf("指针变量 *pstr = %x\n", ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %x\n", *pptr)
}

func swap2(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func pointP4() {
	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a 的值：%d\n", a )
	fmt.Printf("交换前 b 的值：%d\n", b )

	swap2(&a,&b);

	fmt.Printf("交换后 a 的值：%d\n", a )
	fmt.Printf("交换后 b 的值：%d\n", b )
}

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func structP() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	structP2(Book1)
	structP2(Book2)
}

func structP2(book Books) {
	fmt.Printf(book.title)
	fmt.Printf(book.author)
	fmt.Printf(book.subject)
	fmt.Println(book.book_id)
}

func sliceP() {
	numbers := []int{0,1,2,3,4,5,6,7,8,9}
	printSlice(numbers)

	fmt.Println("numbers == ",numbers)
	fmt.Println("numbers[1:4] == ",numbers[1:4])
	fmt.Println("numbers[:3] == ",numbers[:3])
	fmt.Println("numbers[4:] == ",numbers[4:])
	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2)

	number3 := numbers[2:5]
	printSlice(number3)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main() {
	sliceP()
}
