package tests

import (
	"testing"

	lvn "github.com/Lavina-Tech-LLC/lavinagopackage/v2"
)

type ()

func TestTernary(t *testing.T) {

	res := []testsRes[int]{
		{
			Out:  lvn.Ternary(true, 0, 1),
			Want: 0,
			Test: "Ternary with bool",
		},
		{
			Out:  lvn.Ternary(0, 0, 1),
			Want: 1,
			Test: "Ternary with int",
		},
		{
			Out:  lvn.Ternary("", 0, 1),
			Want: 1,
			Test: "Ternary with string",
		},
	}
	check(res, t)

}

func TestGetValue(t *testing.T) {

	type nest2Type struct {
		StringField string
	}
	type nestType struct {
		StringField string
		ArrayField  []string
		Nest2       nest2Type
	}
	type testType struct {
		StringField string
		IntField    int
		ArrayField  []string
		Nest        nestType
	}

	val := testType{
		StringField: "Level1 string",
		IntField:    1,
		ArrayField: []string{
			"level1 array1",
			"level1 array2",
		},
		Nest: nestType{
			StringField: "Level2 string",
			ArrayField: []string{
				"level2 array1",
				"level2 array2",
			},
			Nest2: nest2Type{
				StringField: "Level3 string",
			},
		},
	}

	res := []testsRes[string]{
		{
			Out:  lvn.GetValue[string](val, "StringField"),
			Want: val.StringField,
			Test: "GetValue 1st Level",
		},
		{
			Out:  lvn.GetValue[string](val, "Nest", "StringField"),
			Want: val.Nest.StringField,
			Test: "GetValue 1st nest",
		},
		{
			Out:  lvn.GetValue[[]string](val, "Nest", "ArrayField")[0],
			Want: val.Nest.ArrayField[0],
			Test: "GetValue array",
		},
		{
			Out:  lvn.GetValue[string](val, "Nest", "Nest2", "StringField"),
			Want: val.Nest.Nest2.StringField,
			Test: "GetValue 2nd nest",
		},
	}
	check(res, t)
}
