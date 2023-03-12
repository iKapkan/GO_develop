package dose

import (
	"fmt"
)

func Ints() {
	var First uint8 = 17
	fmt.Printf("Значение First %v\n", First)
	Second := &First //*int16 &-обращается в ячейку памяти "First"
	*Second = 25     // указатель перезаписал var w и теперь там не 17 а 25
	var h = 77
	var _ = h                                                    // чтоб не жаловался что h не вызывается
	fmt.Printf("Значение First %v, перезаписан Second\n", First) //reflect.TypeOf(w)
	fmt.Printf("Значение Second %v\n", *Second)                  //reflect.TypeOf(*f)
	fmt.Printf("Указатель на ячейку памяти W %v\n", &Second)
}
