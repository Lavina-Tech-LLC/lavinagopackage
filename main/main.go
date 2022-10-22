package main

import (
	"fmt"

	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/llog"
)

func main() {
	logger := llog.New(true)
	logger.Use(func(l *llog.Log) {
		fmt.Println("MIDDLEWARE > " + l.Message)
		if l.Level == llog.Critical {
			l.Message = "****** critical ******\n" + l.Message
		}
	})

	logger.Infof("This is info %s", llog.Info)
	logger.Debugf("This is debug")
	logger.Warningf("This is warning")
	logger.Noticef("This is notice")
	logger.Errorf("This is error")
	logger.Criticalf("This is critical")

	logger.Info(logger)

}
