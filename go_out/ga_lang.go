package go_out

import "fmt"

func Speacks() {
	var (
		name string
		age  uint8
	)
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)
	fmt.Printf("Hello %s, you are %d years old\n", name, age)
}
