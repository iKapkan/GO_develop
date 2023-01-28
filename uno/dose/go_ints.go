package dose

import (
	"fmt"
	"reflect"
)

func Ints() {
	var t uint8 = 16 // не отрицательные
	fmt.Println(reflect.TypeOf(t), t)

	var e float32 = 17.56
	fmt.Println(reflect.TypeOf(e), e)

	var w int16 = 17
	fmt.Println(reflect.TypeOf(w), w)
	var f = &w
	fmt.Print(reflect.TypeOf(f), f, " Указатель на ячейку памяти w\n")

}
