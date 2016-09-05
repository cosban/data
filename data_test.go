package data

import (
	"database/sql"
	"testing"
)

func TestSetStringNoOtherwise(t *testing.T) {
	expected := "success"
	input := sql.NullString{
		Valid:  true,
		String: "success",
	}
	result := SetString(input, "failure")
	if expected != result {
		t.Log("Expected %s", expected)
		t.Log("Received %s", result)
		t.FailNow()
	}
}

func TestSetStringOtherwise(t *testing.T) {
	expected := "success"
	input := sql.NullString{
		Valid:  false,
		String: "failure",
	}
	result := SetString(input, "success")
	if expected != result {
		t.Log("Expected %s", expected)
		t.Log("Received %s", result)
		t.FailNow()
	}
}

func TestSetIntNoOtherwise(t *testing.T) {
	expected := 1
	input := sql.NullInt64{
		Valid: true,
		Int64: 1,
	}
	result := SetInt(input, 0)
	if expected != result {
		t.Log("Expected %s", expected)
		t.Log("Received %s", result)
		t.FailNow()
	}
}

func TestSetIntOtherwise(t *testing.T) {
	expected := 1
	input := sql.NullInt64{
		Valid: false,
		Int64: 0,
	}
	result := SetInt(input, 1)
	if expected != result {
		t.Log("Expected %s", expected)
		t.Log("Received %s", result)
		t.FailNow()
	}
}
