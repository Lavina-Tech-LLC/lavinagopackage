package tests

import (
	"encoding/json"
	"testing"

	lvn "github.com/Lavina-Tech-LLC/lavinagopackage/v2"
)

type ()

func TestTernary(t *testing.T) {
	var nilArray []string
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
		{
			Out:  lvn.Ternary(nilArray, 0, 1),
			Want: 1,
			Test: "Ternary with nil array",
		},
		{
			Out:  lvn.Ternary([]string{}, 0, 1),
			Want: 0,
			Test: "Ternary with array",
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

func TestMarshall(t *testing.T) {
	type (
		baseArrayItem struct {
			A_field1 string
			A_field2 int
			A_field3 bool
		}
		base struct {
			Field1 string
			Field2 string
			Field3 struct {
				Field4 int
				Field5 float64
			}
			Array []baseArrayItem
		}

		ccBaseArrayItem struct {
			A_field1 string `json:"a_field1"`
			A_field2 int    `json:"a_field2"`
			A_field3 bool   `json:"a_field3"`
		}
		ccBase struct {
			Field1 string `json:"field1"`
			Field2 string `json:"field2"`
			Field3 struct {
				Field4 int     `json:"field4"`
				Field5 float64 `json:"field5"`
			} `json:"field3"`
			Array []ccBaseArrayItem `json:"array"`
		}

		editedArrayItem struct {
			A_field1 string `json:"a_field1"`
			A_field3 bool   `json:"a_field3"`
		}
		editedBase struct {
			Field2 string `json:"field2"`
			Field3 struct {
				Field4 int `json:"field4"`
			} `json:"field3"`
			Array []editedArrayItem `json:"array"`
		}

		seasonedBase struct {
			Field2 string `json:"field2"`
		}
	)

	Base := base{
		Field1: "field1",
		Field2: "field2",
		Field3: struct {
			Field4 int
			Field5 float64
		}{
			Field4: 15,
			Field5: 0.35,
		},
		Array: []baseArrayItem{
			{
				A_field1: "a_field1",
				A_field2: 78,
				A_field3: true,
			},
			{
				A_field1: "a_field12",
				A_field2: 44,
				A_field3: false,
			},
			{
				A_field1: "a_field13",
				A_field2: 9,
				A_field3: true,
			},
		},
	}

	CCBase := ccBase{
		Field1: "field1",
		Field2: "field2",
		Field3: struct {
			Field4 int     `json:"field4"`
			Field5 float64 `json:"field5"`
		}{
			Field4: 15,
			Field5: 0.35,
		},
		Array: []ccBaseArrayItem{
			{
				A_field1: "a_field1",
				A_field2: 78,
				A_field3: true,
			},
			{
				A_field1: "a_field12",
				A_field2: 44,
				A_field3: false,
			},
			{
				A_field1: "a_field13",
				A_field2: 9,
				A_field3: true,
			},
		},
	}

	EditedBase := editedBase{
		Field2: "field2",
		Field3: struct {
			Field4 int `json:"field4"`
		}{
			Field4: 15,
		},
		Array: []editedArrayItem{
			{
				A_field1: "a_field1",
				A_field3: true,
			},
			{
				A_field1: "a_field12",
				A_field3: false,
			},
			{
				A_field1: "a_field13",
				A_field3: true,
			},
		},
	}

	SeasonedBase := seasonedBase{
		Field2: "field2",
	}

	res := []testsRes[string]{}

	got, _ := lvn.Marshal(Base)
	want, _ := json.Marshal(CCBase)
	res = append(res, testsRes[string]{
		Want: string(want),
		Out:  string(got),
		Test: "Camel case test",
	})

	got, _ = lvn.Marshal(Base, "a_field2", "field1", "field5")
	want, _ = json.Marshal(EditedBase)
	res = append(res, testsRes[string]{
		Want: string(want),
		Out:  string(got),
		Test: "Omit test",
	})

	got, _ = lvn.MarshalSelected(Base, "field2")
	want, _ = json.Marshal(SeasonedBase)
	res = append(res, testsRes[string]{
		Want: string(want),
		Out:  string(got),
		Test: "Select case test",
	})

	check(res, t)
}
