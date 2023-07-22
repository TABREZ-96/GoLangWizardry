package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	fmt.Println(concat("Lane,", "happy Birthday"))
	fmt.Println(concat("Elon ", "hope thats tesla"))
	fmt.Println(concat("Go ", " thats tesla"))

}
