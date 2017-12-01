package main

import (
	"flag"
	"fmt"
)

// 試合スケジュール
type MatchSchedule struct {
	Match [][]int
	Num   int
	Max   int
}

func CheckTargetValue(num int, target []int) []int {
	if target[1] > num {
		target[1] = target[1] % num
	} else if target[0] == target[1] {
		target[1]++
	} else if target[0] > target[1] {
		tmp := target[1]
		target[1] = target[0]
		target[0] = tmp
	} else {
		return target
	}
	return CheckTargetValue(num, target)
}

// 対戦順を生成
func createMatchSchdule(num int) *MatchSchedule {
	// 対戦順生成
	matches := &MatchSchedule{Max: num * (num - 1) / 2, Num: num, Match: [][]int{}}
	turn := 1
	current := 1
	for i := 0; i < matches.Max; i++ {
		target := CheckTargetValue(matches.Num, []int{current, current + turn})

		for j := 0; j < len(matches.Match); j++ {
			if matches.Match[j][0] == target[0] && matches.Match[j][1] == target[1] {
				break
			}
			target[1]++
			target = CheckTargetValue(matches.Num, target)
		}

		matches.Match = append(matches.Match, target)
		fmt.Println(matches.Match)
		current++
		if turn == 1 {
			current++
		}
		if current > matches.Num {
			current = 1
			turn++
		}
	}
	return matches
}

func main() {
	var num int
	// 引数からチーム数設定
	flag.IntVar(&num, "num", 0, "int flag")
	flag.Parse()

	// 対戦表作成
	matches := createMatchSchdule(num)
	fmt.Println(matches.Match)
}
