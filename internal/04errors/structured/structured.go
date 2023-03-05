package structured

import "github.com/sirupsen/logrus"

// Hook 구조체는 logrus의 hook 인터페이스를
// 구현한다.
type Hook struct {
	id string
}

// Fire 함수는 로그를 기록할 때마다 호출한다.
func (h *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = h.id
	return nil
}

// Levels 함수는 hook이 처리할 로그 레벨을 반환한다.
func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Logrus 함수는 몇가지 기본적인 logrus 기능을 보여준다.
func Logrus() {
	// json 포맷으로 로그를 기록한다.
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{id: "1234"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Somethig happend", "Just Now"}

	x := logrus.WithFields(fields)
	x.Warn("warning")
	x.Error("error!")
}
