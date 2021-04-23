package main

import (
	"log"
)

func main() {
	//rows, err := QueryDbData(nil, "")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//println(rows)

	data, err := Download("https://ftp.gnu.org/gnu/binutils/binutils-2.7.tar.gz")
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	println("%s", data)
}
