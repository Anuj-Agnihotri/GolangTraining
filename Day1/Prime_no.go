package main

import (
    "fmt"
)

func IsPrime(value int) bool {
    for i := 2; i <= value/2; i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func main() {
	var num int
	fmt.Scan(&num)
        if IsPrime(num) {
            fmt.Printf("%v is Prime\n", num)
		} else{
			fmt.Printf("%v is not Prime\n",num)
		}
}