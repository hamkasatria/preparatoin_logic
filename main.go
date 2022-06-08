package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	// fmt.Println("Hello world")
	// alt case

	AltCase("Hello World")

	//mutiplecase
	Multiple3And5(int(20))

	//fiibonace
	NeFibbon([]int32{15, 1, 3})

	//rev
}

// func GetBiggestNUmber(data int32) {

// 	num := helperBigNUmber(data)
// 	fmt.Println(num)
// }

// func helperBigNUmber(data int32) int32 {
// 	string := strconv.Itoa(data)

// 	return nil(string)
// }

func MapArrayData(data []int) {
	sort.Slice(data, func(i, j int) bool {
		return data[j] < data[i]
	})
	fmt.Println(data)
}

func AltCase(data string) {
	arr := []rune(data)
	var newArr []string
	for _, v := range arr {
		if unicode.IsUpper(v) {
			newArr = append(newArr, strings.ToLower(string(v)))
		} else {
			newArr = append(newArr, strings.ToUpper(string(v)))
		}
	}
	fmt.Println(strings.Join(newArr, ""))
}

func Multiple3And5(num int) {
	value := 0
	for i := 0; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			value += i
		}
	}
	fmt.Println(value)

}

func NeFibbon(data []int32) {
	var sum int32
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	fmt.Println(Fib(sum, 1, 1))
}

func Fib(sum int32, n1 int32, n2 int32) int32 {
	count := n1 + n2
	var rangeUp, rangeDown int32
	if sum <= n2 && sum >= n1 {
		rangeUp = n2 - sum
		rangeDown = sum - n1
		if rangeUp > rangeDown {
			return rangeDown
		}
		return rangeUp
	}
	return Fib(sum, n2, count)
}

// func Rev(data string) string {
// 	arr := []rune(data)
// 	var newArr []string
// 	for _, v := range arr {
// 		if unicode.IsUpper(v) {
// 			newArr = append(newArr, strings.ToLower(string(v)))
// 		} else {
// 			newArr = append(newArr, strings.ToUpper(string(v)))
// 		}
// 	}
// 	fmt.Println(strings.Join(newArr, ""))
// }
