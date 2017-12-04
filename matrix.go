package main

import (
	"flag"
	"fmt"
	//	"reflect"
)

/*func IsContain(list [][]int, target []int) bool {
	for i := range list {
		reverse := []int{target[1], target[0]}
		if reflect.DeepEqual(list[i], target) || reflect.DeepEqual(list[i], reverse) {
			return true
		}
	}
	return false
}*/

func CreateMatchList(num int) [][]int {
	matchNum := num * (num - 1) / 2
	matchList := [][]int{}

	current := 1
	turn := 1
	for i := 0; i < matchNum; i++ {
		target := []int{current, current + turn}
		if target[1] > num {
			target[1] = target[1] % num
		}

		matchList = append(matchList, target)
		current++
		if current >= num || target[1] > num {
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
