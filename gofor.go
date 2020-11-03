package main

import "fmt"

func main() {
	
	i := 1

	for num := 1; num <= 10; num++ {
		if (num%2 == 0) {
			fmt.Println ("Even ~> ", num)
		} else {
			fmt.Println ("Odd ~> ", num)
		}
	}
	if (99%2 == 0) {
		fmt.Println("99 is even")
	} else { 
		fmt.Println("99 is odd")
	}


	for i < 10 {
		fmt.Println(i)
	i++
	}

	for j := 9; j <= 100; j = j + 10 {
		fmt.Println(j)
	}

	for x := 0; x < 10; x++ {
		if x%2 == 0 {
			fmt.Println("~> ", x)
		}
	}
}

