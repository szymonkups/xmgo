package main

import (
	"fmt"
	"log"

	"github.com/szymonkups/xmgo/xmparser"
)

func main() {
	xm, err := xmparser.ParseFile("./xm_files/millenium.xm")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(xm)
}
