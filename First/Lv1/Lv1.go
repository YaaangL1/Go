package Lv1

import "fmt"

func InvertString() {
	var input string
	fmt.Scan(&input)
	//中文字符在UTF8中占用三个字符，英文和数字一个，所以需要转换成rune（int32）
	var (
		inputRune = []rune(input)
		output    string
	)
	for i := 0; i < len(inputRune)/2; i++ {
		if inputRune[i] != inputRune[len(inputRune)-i-1] {
			fmt.Printf("false")
			return
		}
	}
	for i := 0; i < len(inputRune)/2; i++ {
		output += string(inputRune[i])
	}
	fmt.Println(output)
}
