package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// WrappedError 함수는 오류 래핑과 오류에 주석을 표시하는
// 방법을 보여준다.
func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

// ErrorTyped 구조체는 예제에서 확인할 오류다.
type ErrorTyped struct {
	error
}

// Wrap 함수는 오류를 래핑할 때 호출한다.
func Wrap() {
	e := errors.New("standard error")
	fmt.Println("Ragular Error - ", WrappedError(e))
	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))
	fmt.Println("Nil Error - ", WrappedError(nil))
}
