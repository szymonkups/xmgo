package main

import (
	"log"

	"github.com/szymonkups/xmgo/xmparser"
)

func main() {
	_, err := xmparser.ParseFile("./xm_files/millenium.xm")

	if err != nil {
		log.Fatalln(err)
	}
}
