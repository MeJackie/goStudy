package main

import (
	"fmt"
	"github.com/jian/study/advancedGoProgramming/distributedSystem"
)

func main() {
	var indexes = []int{0,1,2,3,4,5,6}
	distributedSystem.Shuffle(indexes)
	fmt.Println(indexes)
}


