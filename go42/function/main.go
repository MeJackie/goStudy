package main

import(
	"os"
	"time"
	"fmt"
)

type funcType func(time.Time)

var G int = 7

func main() {
	y := func() int {
		fmt.Printf("G: %d, G的地址 %p \n", G, &G)
		G += 1
		return G
	}
	fmt.Println(y(), y)
	fmt.Println(y(), y)
	fmt.Println(y(), y)

	z := func() int {
		G += 1
		return G
	}

	fmt.Println(z, &z)
	fmt.Println(z, &z)
	fmt.Println(z, &z)





	//var a = complex(2, 2)
	const b = complex(1.0, -1.4)
	fmt.Println(b)
	os.Exit(1)


	f := func(t time.Time) time.Time {return t}
	fmt.Println(f(time.Now()))

	var timer funcType = CurrentTime
	timer(time.Now())

	funcType(CurrentTime)(time.Now())
}

func CurrentTime(start time.Time)  {
	fmt.Println(start)
}
