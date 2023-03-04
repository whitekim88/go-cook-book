package confformat

import (
	"bytes"
	"fmt"

	"github.com/go-yaml/yaml"
)

// YAMLData 구조체는 YAML 구조체 태그를 갖는
// 일반적인 데이터 구조체다
type YAMLData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

// ToYaml 함수는 YAMLData 구조체를
// YAML 포맷의 btyes.Buffer로 덤프(반환)한다.
func (t *YAMLData) ToYaml() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(d)
	return b, nil
}

// Decode 함수는 YAMLData 구조체로 디코딩 한다.
func (t *YAMLData) Decode(data []byte) error {
	fmt.Println(string(data))
	return yaml.Unmarshal(data, t)
}
