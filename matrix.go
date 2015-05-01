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
	fmt.Print("n : ")
	fmt.Println(n)
	fmt.Print("max : ")
	max = 2*n - 1
	fmt.Println(max)

	var m [10][10]int
	var iv [18]int
	for iy = 0; iy < n2; iy++ {
		for ix = 0; ix < n2; ix++ {
			if ix != iy {
				index = 0
				for i = 0; i < ix; i++ {
					if i != iy {
						iv[index] = m[i][iy]
						index++
					}
				}
				for i = 0; i < iy; i++ {
					if i != ix {
						iv[index] = m[ix][i]
						index++
					}
				}
				min = 1
				for j = 1; j <= max; j++ {
					var found bool
					found = false
					for i = 0; i < index; i++ {
						if j == iv[i] {
							found = true
							break
						}
					}
					if found == true {
						continue
					}
					min = j
					break
				}
				m[ix][iy] = min
			}
		}
		for i = 0; i < n2; i++ {
			fmt.Print(m[i][iy])
		}
		fmt.Println()
	}
}
