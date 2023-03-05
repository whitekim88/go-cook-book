package main

import "github.com/whitekim88/go-cook-book/internal/04errors/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
