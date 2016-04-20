package apiv1

import "database/sql"

//ImportJSONHandler is mapped to the API method import_json
func ImportJSONHandler() {
	var db *sql.DB
	importJSON(db, "")
}
