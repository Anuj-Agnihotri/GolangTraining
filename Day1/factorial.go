package main

import "fmt"

func Factorial(value int){
	fact:=1
	for i:=value; i>0; i--{
		fact = fact*i
	}
	fmt.Printf("Factorial of %d is %d ",value,fact)
}

func main(){
	var num int
	fmt.Scan(&num)
	Factorial(num)
}