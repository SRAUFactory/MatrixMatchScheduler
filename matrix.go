package main

import (
	"flag"
	"fmt"
)

func isExistList(matchList [][]int, target []int) bool {
	for i := 0; i < len(matchList); i++ {
		currentMatch := matchList[i]
		if currentMatch[0] == target[0] && currentMatch[1] == target[1] {
			return true
		}
		if currentMatch[0] == target[1] && currentMatch[1] == target[0] {
			return true
		}
	}
	return false
}

func createMatchList(num int) [][]int {
	matchNum := num * (num - 1) / 2
	turnMatchNum := num/2 + num%2
	fmt.Printf("CreateMatchList %d, %d, %d\n", num, matchNum, turnMatchNum)
	matchList := [][]int{}

	current := 1
	turn := 1
	for i := 1; i <= matchNum; i++ {
		fmt.Printf("start i:%d, current:%d, turn:%d\n", i, current, turn)
		target := []int{current, current + turn}
		if target[1] > num {
			target[1] = target[1] % num
		}

		limit := 0
		for isExistList(matchList, target) || limit >= num {
			target[1]++
			limit++
			if target[1] > num {
				target[1] = target[1] % num
			} else if target[0] == target[1] {
				target[1]++
			}
		}
		fmt.Println(target)

		matchList = append(matchList, target)
		current++
		if i%turnMatchNum == 0 {
			current = 1
			turn++
		} else if turn == 1 {
			current = target[1] + turn
		}
	}
	return matchList
}

func main() {
	var num int
	// 引数からチーム数設定
	flag.IntVar(&num, "num", 0, "int flag")
	flag.Parse()

	// 対戦表作成
	matches := createMatchList(num)
	fmt.Println(matches)
}
