package apiv1

import (
	"database/sql"
	"testing"
)

func TestImportJSON(t *testing.T) {
	/*
		importJSON takes in a string from the API request representing a
		JSON object. It runs validation, trims excess, stores to database, and generates a response with success/fail information
	*/
	var input = "{\"\"}"
	var expected = "{\"success\":true}"

	var DB *sql.DB

	ret := importJSON(DB, input)
	if ret != expected {
		t.Error("Expected", expected, "got", ret)
	}
}

func TestValidateJSON(t *testing.T) {
	//Validates the JSON string from the request

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
