package go_out

import (
	"fmt"
	"os"
)

func Speacks() {
	var (
		name string
		age  uint8
	)

	fmt.Println("Enter your name: ")
	var a, b = fmt.Scan(&name) //func Scan(a ...any) (n int, err error)
	a += 0
	if b != nil {
		fmt.Printf("Ошибка: %v", b)
		os.Exit(0)
	}
	fmt.Println("Enter your age: ")
	var t, y = fmt.Scan(&age)
	t += 0
	if y != nil {
		fmt.Printf("Ошибка: %v", y)
		os.Exit(0)
	}
	fmt.Printf("Hello %s, you are %d years old\n", name, age)
}
