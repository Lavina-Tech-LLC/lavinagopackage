package tests

import (
	"fmt"
	"testing"

	lvn "github.com/Lavina-Tech-LLC/lavinagopackage/v2"
)

func TestResSelected(t *testing.T) {
	bodyIn := struct {
		Field1 string
		Field2 string
		Field3 struct {
			Field4 string
			Field5 string
		}
	}{
		Field1: "Field1",
		Field2: "Field2",
		Field3: struct {
			Field4 string
			Field5 string
		}{
			Field4: "Field4",
			Field5: "Field5",
		},
	}
	code, message, body := lvn.ResSelected(200, bodyIn, "message", "field3")

	fmt.Println(code, message, string(body))
}
