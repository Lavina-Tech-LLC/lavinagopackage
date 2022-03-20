package llog

import "fmt"

const (
	InfoF    = "\033[1;34m[INFO]\033[0m %s"
	NoticeF  = "\033[1;36m[Notice]\033[0m %s"
	WarningF = "\033[1;33m[Warning]\033[0m %s"
	ErrorF   = "\033[1;31m[Error]\033[0m %s"
	DebugF   = "\033[0;36m[Debug]\033[0m %s"
)

func Info(s interface{}) {
	fmt.Printf(InfoF, s)
	fmt.Println("")
}

func Notice(s interface{}) {
	fmt.Printf(NoticeF, s)
	fmt.Println("")
}

func Warning(s interface{}) {
	fmt.Printf(WarningF, s)
	fmt.Println("")
}

func Error(s interface{}) {
	fmt.Printf(ErrorF, s)
	fmt.Println("")
}

func Debug(s interface{}) {
	fmt.Printf(DebugF, s)
	fmt.Println("")
}
