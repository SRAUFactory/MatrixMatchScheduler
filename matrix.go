package main

import (
	"flag"
	"fmt"
)

func createMatchCards(num int) [][]int {
	matchCards := [][]int{}
	for iy := 1; iy <= num; iy++ {
		for ix := iy + 1; ix <= num; ix++ {
			current := []int{iy, ix}
			matchCards = append(matchCards, current)
		}
	}

	return matchCards
}

func main() {
	var num int
	// 引数からチーム数設定
	flag.IntVar(&num, "num", 0, "int flag")
	flag.Parse()
	fmt.Printf("チーム数: %d\n", num)

	matchCards := createMatchCards(num)
	fmt.Println(matchCards)
}
