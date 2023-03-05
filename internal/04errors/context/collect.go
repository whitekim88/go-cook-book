package context

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Initialize 함수는 설정을 위해 세 개의 함수를 호출한다.
// 그런 다음, 종료 전에 로그를 기록한다.
func Initialize() {
	// 기본 로그를 설정한다
	log.SetHandler(text.New(os.Stdout))
	// context를 초기화 한다.
	ctx := context.Background()
	// 로거를 생성하고 context에 연결한다.
	ctx, e := FromContext(ctx, log.Log)

	// 필드를 설정한다..
	ctx = WithField(ctx, "id", "123")
	e.Info("starting")
	gatherName(ctx)
	e.Info("after gatherName")
	gatherLocation(ctx)
	e.Info("after gatherLocation")
}

func gatherName(ctx context.Context) {
	ctx = WithField(ctx, "name", "Go CookBook")
}

func gatherLocation(ctx context.Context) {
	ctx = WithFields(ctx, log.Fields{"city": "Seoul", "country": "Korea"})
}
