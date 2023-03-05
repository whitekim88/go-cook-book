package errwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// UnWrap 함수는 오류의 래핑을 해제하고
// 오류의 타입 어설션을 수행한다.
func UnWrap() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "Wrapped")

	// 다양한 오류 타입을 처리 할 수 있다.
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred", err)
	default:
		fmt.Println("an unkonwn error occurred")
	}
}

// StackTrace 함수는 오류에 대한 모든 스택을 출력한다.
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")
	fmt.Printf("%+v\n", err)
}
