package main

import (
	"fmt"

	"github.com/Alpensin/go-obsmonster/api/rest"
)

func main() {
	fmt.Println("start")
	rest.Run()
	fmt.Println("finish")
}
