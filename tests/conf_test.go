package tests

import (
	"testing"

	"github.com/Lavina-Tech-LLC/lavinagopackage/v2/conf"
)

func TestConf(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("panic: %v", err)
		}
	}()
	type (
		confType struct {
			Field1 string
			Field2 int
		}
	)

	config := confType{
		Field1: "Hello world",
		Field2: 13,
	}

	conf.Load[confType]()
	conf.Set(config)
	res := []testsRes[confType]{
		{Out: conf.Get[confType](),
			Want: config,
			Test: "Config set and get"},
	}
	check(res, t)
}
