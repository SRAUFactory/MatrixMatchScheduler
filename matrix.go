package main

import (
	"flag"
	"fmt"
)

// 試合スケジュール
type matchSchedule struct {
	match [10][10]int
	num   int
	max   int
}

// 対戦順を出力
func printMatchSchdule(matches matchSchedule) {
	j := 1
	for i := 1; i <= matches.max; i++ {
		for iy := 0; iy < matches.num; iy++ {
			for ix := iy + 1; ix < matches.num; ix++ {
				if matches.match[ix][iy] == i {
					fmt.Printf("%d: %d - %d\n", j, iy+1, ix+1)
					j++
				}
			}
		}
	}
}

// 設定済みの値を取得する
func getValueList(matches matchSchedule, ix int, iy int) [18]int {
	var iv [18]int
	index := 0
	for i := 0; i < ix; i++ {
		iv[index] = matches.match[i][iy]
		index++
	}
	for i := 0; i < iy; i++ {
		iv[index] = matches.match[ix][i]
		index++
	}
	return iv
}

// 抽出した値の中に含まれていない値の最小値を算出
func getMinmumValue(matches matchSchedule, iv [18]int, min int) int {
	for j := min; j <= matches.max; j++ {
		isFound := false
		for i := 0; i < len(iv); i++ {
			if j == iv[i] {
				isFound = true
				break
			}
		}
		if isFound == true {
			if j == matches.max {
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
func createMatchSchdule(num int) matchSchedule {
	// 対戦表を作るチーム数は偶数になるように設定
	// 奇数の場合は+1する
	n := num/2 + num%2
	n2 := n * 2

	// 1チームあたりの最大試合数
	var matches matchSchedule
	matches.max = 2*n - 1
	matches.num = num

	// 対戦順生成
	for iy := 0; iy < n2; iy++ {
		for ix := iy + 1; ix < n2; ix++ {
			if ix != iy {
				// 対象行/列で設定済みの値をまとめて抽出
				iv := getValueList(matches, ix, iy)

				// 抽出した値の中に含まれていない値の最小値を算出
				matches.match[ix][iy] = getMinmumValue(matches, iv, ix)
				matches.match[iy][ix] = matches.match[ix][iy]
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
