package main

import (
	"flag"
	"fmt"
)

func main() {
	var num int
	// 引数からチーム数設定
	flag.IntVar(&num, "num", 0, "int flag")
	flag.Parse()

	fmt.Printf("チーム数: %d\n", num)
}
