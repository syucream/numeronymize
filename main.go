package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/syucream/numeronymize/src/numeronymize"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	numeronymized := numeronymize.Numeronymize(string(in))
	fmt.Println(numeronymized)
}
