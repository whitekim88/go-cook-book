package main

import (
	"fmt"

	"github.com/whitekim88/go-cook-book/internal/04errors/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
