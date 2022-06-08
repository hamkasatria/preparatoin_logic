package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// fmt.Println("Hello world")
	//BigNumber
	GetBiggestNUmber(592)

	//MapData
	MapArrayData([]int{15, 1, 3})

	// alt case
	AltCase("Hello World")

	//mutiplecase
	Multiple3And5(int(20))

	//fiibonace
	NeFibbon([]int32{15, 1, 3})

	//rev
	Rev("Hallo World")

	//midalf
	MidAlf("Q", "Z")
}

func GetBiggestNUmber(data int) {
	num, _ := strconv.Atoi(*helperBigNUmber(data))
	fmt.Println(num)
}

func helperBigNUmber(data int) *string {
	stt := strconv.Itoa(data)
	arr := []rune(stt)
	var arrInt []int
	var arrStr []string
	if len(arr) < 3 {
		return nil
	}

	for _, v := range arr {
		num, _ := strconv.Atoi(string(v))
		if num <= 0 {
			return nil
		}
		arrInt = append(arrInt, num)
	}

	sort.Slice(arrInt, func(i, j int) bool {
		return int(arrInt[j]) < int(arrInt[i])
	})

	for _, v := range arrInt {
		arrStr = append(arrStr, strconv.Itoa(v))
	}

	all := strings.Join(arrStr, "")
	return &all
}

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

func Rev(data string) {
	var arrStr []string
	arrSt := strings.Fields(data)
	for _, v := range arrSt {
		var tempSt []string
		arr := []rune(v)
		for i, w := range arr {
			if unicode.IsUpper(w) {
				tempSt = append(tempSt, strings.ToUpper(string(arr[len(arr)-i-1])))
			} else {
				tempSt = append(tempSt, strings.ToLower(string(arr[len(arr)-i-1])))
			}
		}
		arrStr = append(arrStr, strings.Join(tempSt, ""))
	}
	fmt.Println(strings.Join(arrStr, " "))
}

func MidAlf(fist string, last string) {
	const Alf = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	arrAlf := []rune(Alf)
	var firstQue, lastQue int
	for i, v := range arrAlf {
		if string(v) == fist {
			firstQue = i
		}
		if string(v) == last {
			lastQue = i
		}
	}
	sumQue := firstQue + lastQue
	avgQue := sumQue / 2

	if sumQue%2 != 0 {
		fmt.Print(string(arrAlf[avgQue]), string(arrAlf[avgQue+1]))
	} else {
		fmt.Print(string(arrAlf[avgQue]))
	}

}
