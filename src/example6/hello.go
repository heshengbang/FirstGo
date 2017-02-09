package main

import "fmt"

func rangeP()  {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:",sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Printf("%d - %c\n",i, c)
		fmt.Println(i,c)
	}
}

func mapP() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	printAll(countryCapitalMap)
	capital, ok := countryCapitalMap["United States"]
	if ok {
		fmt.Println("Capital of United States is ", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}
	delete(countryCapitalMap,"France");
	printAll(countryCapitalMap)
}
func printAll(country map[string]string){
	for k,v := range country  {
		fmt.Println("Capital of ",k," is ",v)
	}
}

func main() {
	mapP()
}
