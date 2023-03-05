package main

import (
	"fmt"

	"github.com/whitekim88/go-cook-book/internal/04errors/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.UnWrap()
	fmt.Println()
	errwrap.StackTrace()
}
