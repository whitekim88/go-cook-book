package main

import (
	"github.com/whitekim88/go-cook-book/internal/01io_filesystem/template"
)

func main() {
	if err := template.RunTemplate(); err != nil {
		panic(err)
	}
	if err := template.InitTemplate(); err != nil {
		panic(err)
	}
	if err := template.HTMLDifferences(); err != nil {
		panic(err)
	}

}
