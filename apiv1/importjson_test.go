package apiv1

import (
	"reflect"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestImportJSON(t *testing.T) {
	/*
		importJSON takes in a string from the API request representing a
		JSON object. It runs validation, trims excess, stores to database, and generates a response with success/fail information
	*/
	t.Log("Testing Import Function")

	input := "{\"\"}"
	expected := "{\"success\":true}"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectCommit()

	ret := importJSON(db, input)
	if ret != expected {
		t.Error("Expected", expected, "got", ret)
	}
}

func TestValidateJSON(t *testing.T) {
	//Validates the JSON string from the request
	input := "{\"id\": 7, \"md5sum\": 6f5902ac237024bdd0c176cb93063dc4, \"prod_id\": 1234}"
	expected := false

	ret := validateJSON(input)
	if ret != expected {
		t.Error("Expected", expected, "got", ret)
	}
}

func TestConvertJSON(t *testing.T) {
	//Converts the JSON string to our data structure
	input := "{\"id\": 7, \"md5sum\": 6f5902ac237024bdd0c176cb93063dc4, \"prod_id\": 1234}"
	expected := []jsonData{}

	expected = append(expected, jsonData{
		ID:     7,
		MD5sum: "6f5902ac237024bdd0c176cb93063dc4",
		ProdID: 1234,
	})

	ret := convertJSON(input)

	if reflect.DeepEqual(ret, expected) != true {
		t.Error("Expected", expected, "got", ret)
	}
}

func TestTrimJSON(t *testing.T) {
	//Trims transactions that are already in the DB
}

func TestWriteJSON(t *testing.T) {
	//Writes the new data to the DB

}

func TestImportResponse(t *testing.T) {
	//Creates a JSON response based on database success/fail

}
