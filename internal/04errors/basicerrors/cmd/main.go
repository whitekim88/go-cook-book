package main

import (
	"fmt"

	"github.com/whitekim88/go-cook-book/internal/04errors/basicerrors"
)

func main() {
	basicerrors.BasicErrors()
	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
