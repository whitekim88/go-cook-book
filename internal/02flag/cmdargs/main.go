package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const version = "1.0.0"
const usage = `Usage:
%s [options]
Commands:
	Greet
	Version
`
const greetUsage = `Usage:
%s greet name [flag]
Positional Arguments:
	name
		the name to greet
Flags:
`

// MenuConf 구조체는 중첩된 명령중에 대한 모들 레벨을 저장한다.
type MenuConf struct {
	Goodbye bool
}

// SetupMenu 함수는 기본 플래그를 설정한다.
func (m *MenuConf) SetupMenu() *flag.FlagSet {
	menu := flag.NewFlagSet("menu", flag.ExitOnError)
	menu.Usage = func() {
		fmt.Printf(usage, os.Args[0])
		menu.PrintDefaults()
	}
	return menu
}

// GetSubMenu 함수는 하위 메뉴에 대한 플래그 모음을 반환한다.
func (m *MenuConf) GetSubMenu() *flag.FlagSet {
	submenu := flag.NewFlagSet("submenu", flag.ExitOnError)
	submenu.BoolVar(&m.Goodbye, "goodbye", false, "say goodbye instead of hello")
	submenu.Usage = func() {
		fmt.Printf(greetUsage, os.Args[0])
		submenu.PrintDefaults()
	}
	return submenu
}

// Greet 함수는 greet 명령에 의해 호출된다.
func (m *MenuConf) Greet(name string) {
	if m.Goodbye {
		fmt.Printf("Goodbye %s\n", name+"!")
	} else {
		fmt.Printf("Hello %s\n", name+"!")
	}
}

// Version 함수는 상수(const)로 저장된 현재 버전을 출력한다.
func (m *MenuConf) Version() {
	fmt.Println("Version:", version)
}

func main() {
	c := MenuConf{}
	menu := c.SetupMenu()
	if err := menu.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing params %s, error: %v", os.Args[1:], err)
		return
	}

	// 명령을 전환하는 데 매개변수를 사용한다.
	// 플래그 또한 매개변수로 사용된다.
	if len(os.Args) > 1 {
		// 대소문자 구별은 하지 않는다.
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			f := c.GetSubMenu()
			if len(os.Args) < 3 {
				f.Usage()
				return
			}
			if len(os.Args) > 3 {
				if err := f.Parse(os.Args[3:]); err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing params %s, error: %v", os.Args[3:], err)
					return
				}
			}
			c.Greet(os.Args[2])
		default:
			fmt.Println("Invalid command")
			menu.Usage()
			return
		}
	} else {
		menu.Usage()
		return
	}
}
