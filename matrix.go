package main

import (
	"flag"
	"fmt"
)

func IsExistList(matchList [][]int, target []int) bool {
	for i := 0; i < len(matchList)-1; i++ {
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

func CreateMatchList(num int) [][]int {
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

		for IsExistList(matchList, target) {
			target[1]++
			if target[1] > num {
				target[1] = target[1] % num
			} else if target[0] == target[1] {
				break
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
	matches := CreateMatchList(num)
	fmt.Println(matches)
}
