package Lv2

import (
	"fmt"
)

func Calculator() {
	//定义变量
	var (
		num1, num2 float64
		cmd        string
		result     []float64
	)
	for {
		fmt.Scan(&num1, &cmd, &num2)
		switch cmd {
		case "+":
			result = append(result, num1+num2)
		case "-":
			result = append(result, num1-num2)
		case "*":
			result = append(result, num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("It is no allowed.")
			} else {
				result = append(result, num1/num2)
			}
		default:
			fmt.Println("No support")
		}
		fmt.Println(result)
	}

}
