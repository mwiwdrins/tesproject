package main

import "fmt"
  
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func main() {

	factor := factorialLoop(5)

	fmt.Println(factor)

}
