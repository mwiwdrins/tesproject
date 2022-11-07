package main

import "fmt"

func returns(name string) string {

	return "Hello " + name
}

func getHallo(names string) string {
	if names == "" {
		return "Hello Sobatkuh"
	} else {
		return "Hello " + names
	}

}

func main() {

	result := returns("ORANG")
	fmt.Println(result)

	result2 := getHallo("")
	result3 := getHallo("Heker")

	fmt.Print(result2)
	fmt.Print(result3)

}
