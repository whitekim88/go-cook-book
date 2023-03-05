package basicerrors

import (
	"errors"
	"fmt"
)

// ErrorValue는 확인을 위한 패키지 수준 오류를 만드는 방법이다.
// 즉, if, err == ErrorValue
var ErrorValue = errors.New("this is a typed error")

// TypedError는 err.(type) == ErrorValue를 수행할 수 있는
// 오류 타입을 만드는 방법이다.
type TypedError struct {
	error
}

// BasicErrors 함수는 오류를 생성하는 몇 가지 방법을 보여준다.
func BasicErrors() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", err)
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorValue
	fmt.Println("package level error: ", err)

	err = TypedError{errors.New("this is a typed error")}
	fmt.Println("typed error: ", err)
}
