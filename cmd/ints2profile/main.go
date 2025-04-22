package main

import (
	"log"

	ip "github.com/takanoriyanagitani/go-asn1-ints2profile"
)

func main() {
	var e error = ip.StdinToIntegersToStatsToDerToStdout()
	if nil != e {
		log.Printf("%v\n", e)
	}
}
