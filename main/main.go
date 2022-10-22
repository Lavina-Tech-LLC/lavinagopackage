package main

import (
	"fmt"

	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/logger"
)

func main() {
	lggr := logger.New(true)
	lggr.Use(func(l *logger.Log) {
		fmt.Println("MIDDLEWARE > " + l.Message)
		if l.Level == logger.Critical {
			l.Message = "****** critical ******\n" + l.Message
		}
	})

	lggr.Infof("This is info %s", logger.Info)
	lggr.Debugf("This is debug")
	lggr.Warningf("This is warning")
	lggr.Noticef("This is notice")
	lggr.Errorf("This is error")
	lggr.Criticalf("This is critical")

	lggr.Info(lggr)

}
