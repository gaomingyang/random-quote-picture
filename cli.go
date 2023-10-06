package main

import (
	"flag"
	"fmt"
)

func main() {
	env := flag.String("env", "", "environment")
	flag.Parse()
	fmt.Println("env:", *env)
}
