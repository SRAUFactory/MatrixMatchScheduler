package main

import (
	"flag"
	"fmt"
)

func CreateMatchList(num int) [][]int {
	matchNum := num * (num - 1) / 2
	turnMatchNum := num/2 + num%2
	fmt.Printf("CreateMatchList %d, %d, %d\n", num, matchNum, turnMatchNum)
	matchList := [][]int{}

	current := 1
	turn := 1
	for i := 1; i <= matchNum; i++ {
		fmt.Printf("start %d, %d, %d\n", i, current, turn)
		target := []int{current, current + turn}
		if target[1] > num {
			target[1] = target[1] % num
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
		fmt.Printf("end %d, %d, %d\n", i, current, turn)
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
