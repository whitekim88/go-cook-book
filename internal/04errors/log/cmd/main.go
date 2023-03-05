package main

import (
	"fmt"

	"github.com/whitekim88/go-cook-book/internal/04errors/log"
)

func main() {
	fmt.Println("basic logging and modification of logger")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
