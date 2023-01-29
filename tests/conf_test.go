package tests

import (
	"fmt"
	"math/rand"
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
		Field2: rand.Intn(1000),
	}
	confBeforeLoad := conf.Get[confType]("")
	conf.Load[confType]("")

	res := []testsRes[confType]{
		{Out: conf.Get[confType](),
			Want: confBeforeLoad,
			Test: "Config set and get"},
	}

	conf.Set(config)
	res = append(res, testsRes[confType]{
		Out:  conf.Get[confType](),
		Want: config,
		Test: "Config set and get"},
	)
	check(res, t)
}

func TestGetPath(t *testing.T) {
	path := "/Volumes/drive/GitHub/lavinagopackage/tests/"
	pathUp1 := "/Volumes/drive/GitHub/lavinagopackage/"
	pathUp2 := "/Volumes/drive/GitHub/"

	res := []testsRes[string]{
		{
			Out:  conf.GetPath(),
			Want: path,
			Test: "Up 0",
		},
	}
	res = append(res,
		testsRes[string]{
			Out:  conf.GetPath(1),
			Want: pathUp1,
			Test: "Up 1",
		},
		testsRes[string]{
			Out:  conf.GetPath(2),
			Want: pathUp2,
			Test: "Up 2",
		},
	)
	check(res, t)

}

func TestUpperConf(t *testing.T) {
	type (
		confType struct {
			Field1 string
			Field2 int
		}
	)

	config := conf.Load[confType]("../conf/")
	fmt.Println(config)
}
