package main

import (
	"flag"
	"fmt"
)

func isSameTeamInParrentMatch(parent []int, current []int, num int) bool {
	return false
}

func getNextIndex(index int, length int) int {
	index++
	if index >= length {
		return 1
	}
	return index
}

func sortMatchCards(target [][]int, num int) [][]int {
	sorted := [][]int{}
	sorted = append(sorted, target[0])
	length := len(target)
	index := getNextIndex(0, length)

	for len(sorted) < length {
		for isSameTeamInParrentMatch(sorted[len(sorted)-1], target[index], num) {
			index = getNextIndex(index, length)
		}
		sorted = append(sorted, target[index])
		index = getNextIndex(index, length)
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
