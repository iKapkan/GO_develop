package uno

import (
	"fmt"
	"os"
)

func Calculator() {
	var (
		CPU_sum   int64 = 100
		Testo_sum int64 = 200
		SSD_sum   int64 = 300
	)
	var variants int8
	var kollvo int64

	fmt.Printf("Change params: \n"+
		"1. CPU  от %vруб.\n"+
		"2. Testo_1 от %vруб.\n"+
		"3. SSD от %vруб.\n", CPU_sum, Testo_sum, SSD_sum)

	var cheng, err = fmt.Scan(&variants)
	cheng += 0
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(0)
	}
	if variants == 1 {
		fmt.Println("kollichestvo CPU: ")
		var cheng, err = fmt.Scan(&kollvo)
		cheng += 0
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(0)
		}
		fmt.Printf("Result: %v rub", kollvo*CPU_sum)
	} else if variants == 2 {
		fmt.Println("kollichestvo Testo_1: ")
		var cheng, err = fmt.Scan(&kollvo)
		cheng += 0
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(0)
		}
		fmt.Printf("Result: %v rub", kollvo*Testo_sum)
	} else if variants == 3 {
		fmt.Println("kollichestvo SSD: ")
		var cheng, err = fmt.Scan(&kollvo)
		cheng += 0
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(0)
		}
		fmt.Printf("Result: %v rub", kollvo*SSD_sum)
	} else {
		fmt.Println("Ну это габэлла, Сэр!")
	}

}
