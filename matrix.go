package main

import (
	"flag"
	"fmt"
)

func isSameTeamInParrentMatch(parent []int, current []int, num int) bool {
	return false
}

func sortMatchCards(target [][]int, num int) [][]int {
	sorted := [][]int{}
	for i := 0; i < len(target); i++ {
		j := i
		if j > 0 {
			for isSameTeamInParrentMatch(target[j-1], target[j], num) {
				j++
				if j >= len(target) {
					j = 1
				}
			}
		}
		sorted = append(sorted, target[j])
	}

	return sorted
}

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

	matchCards = sortMatchCards(matchCards, num)
	fmt.Println(matchCards)
}
