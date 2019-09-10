package main

import (
	"fmt"
	"log"
	"os"

	"github.com/szymonkups/xmgo/xmparser"
)

func main() {
	f, err := os.Open("./xm_files/millenium.xm")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	header, err := xmparser.ParseHeader(f)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(header)
	// id, err := xmparser.ParseIDText(f)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Printf("Id parsed: [%s]\n", id)

	// ver, err := xmparser.ParseVersionNumber(f)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Printf("Version parsed: [%d]\n", ver)
}
