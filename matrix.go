package main

import (
	"flag"
	"fmt"
)

func createMatchSchdule(num int) {
	// 対戦表を作るチーム数は偶数になるように設定
	// 奇数の場合は+1する
	n := num/2 + num%2
	n2 := n * 2

	// 1チームあたりの最大試合数
	max := 2*n - 1

	// 試合順
	var m [10][10]int
	var iv [18]int

	// 対戦順生成
	for iy := 0; iy < n2; iy++ {
		for ix := iy + 1; ix < n2; ix++ {
			if ix != iy {
				index := 0
				for i := 0; i < ix; i++ {
					iv[index] = m[i][iy]
					index++
				}
				for i := 0; i < iy; i++ {
					iv[index] = m[ix][i]
					index++
				}
				min := ix
				for j := min; j <= max; j++ {
					isFound := false
					for i := 0; i < index; i++ {
						if j == iv[i] {
							isFound = true
							break
						}
					}
					if isFound == true {
						if j == max {
							j = 0
						}
						continue
					} else {
						min = j
						break
					}
				}
				m[ix][iy] = min
				m[iy][ix] = m[ix][iy]
			}
		}
	}

	// 対戦順出力
	j := 1
	for i := 1; i <= max; i++ {
		for iy := 0; iy < num; iy++ {
			for ix := iy + 1; ix < num; ix++ {
				if m[ix][iy] == i {
					fmt.Printf("%d: %d - %d\n", j, iy+1, ix+1)
					j++
				}
			}
		}
	}
}

func main() {
	var num int
	// 引数からチーム数設定
	flag.IntVar(&num, "num", 0, "int flag")
	flag.Parse()

	createMatchSchdule(num)
}
