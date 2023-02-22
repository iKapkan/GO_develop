package dose

import (
	"fmt"
	"reflect"
)

func Ints() {
	var t uint8 = 16 // не отрицательные
	fmt.Println(reflect.TypeOf(t), " == ", t)

	var e float32 = 17.56
	fmt.Println(reflect.TypeOf(e), " == ", e)

	var w int16 = 17
	fmt.Println(reflect.TypeOf(w), " == ", w)
	f := &w //*int16 &-обращается в ячейку памяти "w"
	*f = 25 // указатель перезаписал var w и теперь там не 17 а 25
	fmt.Print(reflect.TypeOf(*f), " == ", *f, " Указатель на ячейку памяти w == 17\n")

	var h = 77
	var _ = h // чтоб не жаловался что h не вызывается

}

func Trino() {
	fmt.Println("Центр принятия решения идти за пивом," +
		" принял решение идти за пивом.")
}
