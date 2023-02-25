package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("RUNNING WITH ARG: ", strings.Join(os.Args[1:], " "))
}
