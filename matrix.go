package main

import (
	"flag"
	"fmt"
)

func main() {
	var num, n, n2, ix, iy, i, j, index, min, max int
	flag.IntVar(&num, "num", 0, "int flag")
	flag.Parse()
	n = num/2 + num%2
	n2 = n * 2
	//fmt.Print("n : ")
	//fmt.Println(n)
	//fmt.Print("max : ")
	max = 2*n - 1
	//fmt.Println(max)

	var m [10][10]int
	var iv [18]int
	for iy = 0; iy < n2; iy++ {
		for ix = iy + 1; ix < n2; ix++ {
			if ix != iy {
				index = 0
				for i = 0; i < ix; i++ {
					iv[index] = m[i][iy]
					index++
				}
				for i = 0; i < iy; i++ {
					iv[index] = m[ix][i]
					index++
				}
				min = ix
				for j = min; j <= max; j++ {
					var isFound bool
					isFound = false
					for i = 0; i < index; i++ {
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
		for i = 0; i < n2; i++ {
			fmt.Print(m[i][iy])
		}
		fmt.Println()
	}

	j = 1
	for i = 1; i <= max; i++ {
		for iy = 0; iy < num; iy++ {
			for ix = iy + 1; ix < num; ix++ {
				if m[ix][iy] == i {
					fmt.Printf("%d: %d - %d\n", j, iy+1, ix+1)
					j++
				}
			}
		}
	}
}
