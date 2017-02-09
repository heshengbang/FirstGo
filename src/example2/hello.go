package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a,b,c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为:%d",area)
	println()
	println(a, b, c)

	println("_________________________________")
	const (
		vName1 = "abc"
		vName2 = len(vName1)
		vName3 = unsafe.Sizeof(vName1)
	)
	println(vName1, vName2, vName3)
	println("_________________________________")
	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)
	println(a1, b1, c1)
	println("_________________________________")
	const (
		a2 = iota
		b2
		c2
	)
	println(a2, b2, c2)
	println("_________________________________")
	const (
		a3 = iota   //0
		b3          //1
		c3          //2
		d3 = "ha"   //独立值，iota += 1
		e3          //"ha"   iota += 1
		f3 = 100    //iota +=1
		g3          //100  iota +=1
		h3 = iota   //7,恢复计数
		i3          //8
	)
	fmt.Println(a3, b3, c3, d3, e3, f3, g3, h3, i3)
}
