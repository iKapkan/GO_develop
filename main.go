package main

import (
	rut "main/go_out" // rut - alias
	ht "main/uno"
	fut "main/uno/dose"
)

func main() {
	fut.Ints()
	rut.Speacks() //name, age
	rut.TwoSum([]int{2, 7, 11, 15}, int(9))
	ht.WeatherHttp()
}
