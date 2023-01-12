package main

import (
	"fmt"
	"github.com/kkrypt0nn/gomanuf"
)

func main() {
	manuf, _ := gomanuf.Search("C4:A8:1D:73:D7:8C")
	fmt.Println(manuf)
}
