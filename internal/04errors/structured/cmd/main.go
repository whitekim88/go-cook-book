package main

import (
	"fmt"

	"github.com/whitekim88/go-cook-book/internal/04errors/structured"
)

func main() {
	fmt.Println("Logrus:")
	structured.Logrus()
	fmt.Println()
	fmt.Println("Apex:")
	structured.Apex()
}
