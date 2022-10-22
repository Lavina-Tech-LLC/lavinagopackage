package tests

import (
	"testing"

	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/logger"
)

func testLlog(t *testing.T) {
	lggr := logger.New(true)
	lggr.Use(func(l *logger.Log) {
		l.Message = "MIDDLEWARE > " + l.Message
		if l.Level == logger.Critical {
			l.Message = "****** critical ******\n" + l.Message
		}
	})

	lggr.Infof("This is info %s", logger.Info)
}
