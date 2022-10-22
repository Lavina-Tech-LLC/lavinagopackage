package tests

import (
	"testing"

	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/llog"
)

func testLlog(t *testing.T) {
	logger := llog.New(true)
	logger.Use(func(l *llog.Log) {
		l.Message = "MIDDLEWARE > " + l.Message
		if l.Level == llog.Critical {
			l.Message = "****** critical ******\n" + l.Message
		}
	})

	logger.Infof("This is info %s", llog.Info)
}
