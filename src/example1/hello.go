package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var value float64
	fmt.Println(value)
	var a int = 5;
	var b = 10;
	c := 20;
	fmt.Println(a+' '+b+c);

	var a1  = "fuck"
	var b1 string = "go"
	var c1 bool
	println(a1,b1,c1)

	var (
		vName1 string = "fuck"
		vName2 int64 = 256
	)
	println(vName1,vName2)
	println("___________________________________________________________")
	var vName3 string = "yes"
	var vName4 = "yes"
	println(&vName3, &vName4)
	var vName5 = vName4
	var vName6 = "yes "
	println(&vName5, &vName4, &vName6)
	println("___________________________________________________________")
	const identifier  = "Yes"
	println(identifier)
}
