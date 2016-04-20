package apiv1

import "database/sql"

type jsonData struct {
	ID     int    `json:"id"`
	MD5sum string `json:"md5sum"`
	ProdID int    `json:"prod_id"`
}

type success struct {
	success bool
}

type jsonResp struct {
	success bool
}

func importJSON(db *sql.DB, req string) string {
	return "{\"success\":true}"
}

func validateJSON(data string) bool {
	return true
}

func convertJSON(json string) []jsonData {
	var data []jsonData
	return data
}

//retrieveLastTransaction should be a db handler
func retrieveLastTransaction() string {
	return ""
}

func trimJSON(req string) jsonData {
	var data jsonData
	return data
}

func writeJSON(db *sql.DB, data jsonData) success {
	return success{
		success: true,
	}
}

func importResponse(resp success) string {
	return ""
}
