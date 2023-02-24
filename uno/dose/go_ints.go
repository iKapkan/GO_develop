package dose

import (
	"fmt"
	"reflect"
)

func Ints() {
	var w int16 = 17
	f := &w //*int16 &-обращается в ячейку памяти "w"
	*f = 25 // указатель перезаписал var w и теперь там не 17 а 25
	var h = 77
	var _ = h // чтоб не жаловался что h не вызывается

	fmt.Println(reflect.TypeOf(w), " == ", w)
	fmt.Printf("Тип W== %v, значение %v, был == 17\n", reflect.TypeOf(w), w)

	fmt.Printf("Тип F== %v, значение %v\n", reflect.TypeOf(*f), *f)
	fmt.Printf("Указатель на ячейку памяти W %v\n", &f)
}

func Trino() {
	fmt.Println("Центр принятия решения идти за пивом," +
		" принял решение идти за пивом.")
}
