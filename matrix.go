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

// 対戦順を出力
func printMatchSchdule(matches *MatchSchedule) {
	for iy := 0; iy < matches.Num; iy++ {
		for ix := 0; ix < matches.Num; ix++ {
			switch {
			case ix == iy:
				fmt.Print("-")
			default:
				fmt.Print(matches.Match[ix][iy])
			}
		}
		fmt.Println()
	}
	j := 1
	for i := 1; i <= matches.Max; i++ {
		for iy := 0; iy < matches.Num; iy++ {
			for ix := iy + 1; ix < matches.Num; ix++ {
				if matches.Match[ix][iy] == i {
					fmt.Printf("%d: %d - %d\n", j, iy+1, ix+1)
					j++
				}
			}
		}
	}
}

// 設定済みの値を取得する
func getSetValueList(matches *MatchSchedule, ix int, iy int) []int {
	var values []int = make([]int, ix+iy)
	for i := 0; i < ix; i++ {
		values[i] = matches.Match[i][iy]
	}
	for i := 0; i < iy; i++ {
		values[ix+i] = matches.Match[ix][i]
	}
	return values
}

// 抽出した値の中に含まれていない値の最小値を算出
func getMinmumValue(matches *MatchSchedule, values []int, min int) int {
	for j := min; j <= matches.Max; j++ {
		isFound := false
		for i := 0; i < len(values); i++ {
			if j == values[i] {
				isFound = true
				break
			}
		}
		if isFound == true {
			if j == matches.Max {
				j = 0
			}
			continue
		} else {
			min = j
			break
		}
	}
	return min
}

// 対戦順を生成
func createMatchSchdule(num int) *MatchSchedule {
	// 対戦表を作るチーム数は偶数になるように設定
	// 奇数の場合は+1する
	n := num/2 + num%2
	n2 := n * 2

	// 1チームあたりの最大試合数
	matches := &MatchSchedule{Max: 2*n - 1, Num: num}
	matches.Match = make([][]int, n2)
	for i := 0; i < n2; i++ {
		matches.Match[i] = make([]int, n2)
	}

	// 対戦順生成
	for iy := 0; iy < n2; iy++ {
		for ix := iy + 1; ix < n2; ix++ {
			if ix != iy {
				// 対象行/列で設定済みの値をまとめて抽出
				values := getSetValueList(matches, ix, iy)

				// 抽出した値の中に含まれていない値の最小値を算出
				min := ix
				matches.Match[ix][iy] = getMinmumValue(matches, values, min)
				matches.Match[iy][ix] = matches.Match[ix][iy]
			}
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
	// 対戦順出力
	printMatchSchdule(matches)
}
