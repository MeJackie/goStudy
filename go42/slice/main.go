package main

import (
	"fmt"
	"os"
	"errors"
)

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res := make([]byte, 3)
	copy(res, raw[:3])
	return res
}

func div(a, b int)  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到异常： %s \n", r)
		}
	}()

	if b < 0 {
		panic("除数需要大于0")
	}

	fmt.Println("余数为：", a/b)
}

func main() {
	div(1,0)
	div(1,-1)
	os.Exit(1)
	f := 0
	fmt.Println(fmt.Errorf("square root of negative number %g", f))
	fmt.Println(errors.New("math - square root of negative number"))
	os.Exit(1)

	data := get()
	fmt.Println(len(data), cap(data), &data[0])
	os.Exit(1)

	sli := make([]int, 5, 10)
	fmt.Printf("切片sli长度和容量：%d, %d \n", len(sli), cap(sli))
	fmt.Println(sli)
	newsli := sli[:cap(sli)]
	fmt.Println(newsli)

	var x = []int{2, 3, -1, 22}
	fmt.Printf("切片的长度和容量： %d, %d \n", len(x), cap(x))


}
