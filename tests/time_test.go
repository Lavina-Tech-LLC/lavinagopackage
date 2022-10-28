package tests

import (
	"testing"
	"time"

	lvn "github.com/Lavina-Tech-LLC/lavinagopackage/v2"
)

func TestTime(t *testing.T) {

	inputTime, _ := (time.Parse(
		"2006-01-02T15:04:05-0700",
		"2022-04-15T15:04:05-0700",
	))

	startOfDay, _ := time.Parse(
		"2006-01-02T15:04:05-0700",
		"2022-04-15T00:00:00-0700")
	endOfDay, _ := time.Parse(
		"2006-01-02T15:04:05-0700",
		"2022-04-15T23:59:59-0700")

	startOfWeek, _ := time.Parse(
		"2006-01-02T15:04:05-0700",
		"2022-04-11T00:00:00-0700")
	endOfWeek, _ := time.Parse(
		"2006-01-02T15:04:05-0700",
		"2022-04-17T23:59:59-0700")

	startOfMonth, _ := time.Parse(
		"2006-01-02T15:04:05-0700",
		"2022-04-01T00:00:00-0700")
	endOfMonth, _ := time.Parse(
		"2006-01-02T15:04:05-0700",
		"2022-04-30T23:59:59-0700")

	startOfYear, _ := time.Parse(
		"2006-01-02T15:04:05-0700",
		"2022-01-01T00:00:00-0700")
	endOfYear, _ := time.Parse(
		"2006-01-02T15:04:05-0700",
		"2022-12-31T23:59:59-0700")

	input := lvn.Time(inputTime)

	res := []testsRes[time.Time]{
		{
			Out:  input.StartOfTheDay(),
			Want: startOfDay,
			Test: "StartOfDay",
		},
		{
			Out:  input.EndOfTheDay(),
			Want: endOfDay,
			Test: "EndOfTheDay",
		},
		{
			Out:  input.StartOfTheWeek(),
			Want: startOfWeek,
			Test: "StartOfTheWeek",
		},
		{
			Out:  input.EndOfTheWeek(),
			Want: endOfWeek,
			Test: "EndOfTheWeek",
		},
		{
			Out:  input.StartOfTheMonth(),
			Want: startOfMonth,
			Test: "StartOfTheMonth",
		},
		{
			Out:  input.EndOfTheMonth(),
			Want: endOfMonth,
			Test: "EndOfTheMonth",
		},
		{
			Out:  input.StartOfTheYear(),
			Want: startOfYear,
			Test: "StartOfTheYear",
		},
		{
			Out:  input.EndOfTheYear(),
			Want: endOfYear,
			Test: "EndOfTheYear",
		},
	}

	check(res, t)

}
