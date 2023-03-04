package main

import (
	"flag"
	"fmt"

	pflag "github.com/whitekim88/go-cook-book/internal/02flag/flag"
)

func main() {
	// 설정을 초기화한다.
	c := pflag.Config{}
	c.Setup()

	// 다음 코드는 일반적으로 main에서 호출한다.
	flag.Parse()
	fmt.Println(c.GetMessage())
}
