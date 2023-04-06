package tests

import (
	"testing"

	lvn "github.com/Lavina-Tech-LLC/lavinagopackage/v2"
	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/logger"
)

func TestLoggerUse(t *testing.T) {
	lvn.Logger.Use(logger.Timestamper)
	lvn.Logger.Notice("This is the log")
}
