package main

import (
	"geekGolang/week2"
)

func main() {
	rows, err := week2.QueryDbData(nil, "")
	if err != nil {
		println(err.Error())
		return
	}
	println(rows)
}
