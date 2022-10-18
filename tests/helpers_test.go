package tests

import (
	"testing"

	lvn "github.com/Lavina-Tech-LLC/lavinagopackage/v2"
)

func TestTernary(t *testing.T) {

	k := lvn.Ternary(true, 0, 1)
	if k != 0 {
		t.Errorf("Res was incorrect, got: %d, want: %d.", k, 0)
	}
	k = lvn.Ternary(0, 0, 1)
	if k != 1 {
		t.Errorf("Res was incorrect, got: %d, want: %d.", k, 1)
	}

	k = lvn.Ternary("", 0, 1)
	if k != 1 {
		t.Errorf("Res was incorrect, got: %d, want: %d.", k, 1)
	}

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

	have := lvn.GetValue[string](val, "StringField")
	want := val.StringField
	if have != want {
		t.Errorf("Res was incorrect, got: %s, want: %s.", have, want)
	}

	have = lvn.GetValue[string](val, "Nest", "StringField")
	want = val.Nest.StringField
	if have != want {
		t.Errorf("Res was incorrect, got: %s, want: %s.", have, want)
	}

	have = lvn.GetValue[[]string](val, "Nest", "ArrayField")[0]
	want = val.Nest.ArrayField[0]
	if have != want {
		t.Errorf("Res was incorrect, got: %s, want: %s.", have, want)
	}

	have = lvn.GetValue[string](val, "Nest", "Nest2", "StringField")
	want = val.Nest.Nest2.StringField
	if have != want {
		t.Errorf("Res was incorrect, got: %s, want: %s.", have, want)
	}

}
