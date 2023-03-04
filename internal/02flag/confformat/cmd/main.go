package main

import "github.com/whitekim88/go-cook-book/internal/02flag/confformat"

func main() {
	if err := confformat.MarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := confformat.OtherJSONExample(); err != nil {
		panic(err)
	}
}
