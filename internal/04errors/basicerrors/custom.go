package basicerrors

import "fmt"

// CustomError는 Error() 인터페이스를 구현할 구조체다.
type CustomError struct {
	Result string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("there was an error; %s was the result", c.Result)
}

// SomeFunc 함수는 오류를 반환한다.
func SomeFunc() error {
	return CustomError{Result: "this"}
}
