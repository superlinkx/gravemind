package main

// Page Structure
type Page struct {
	BusinessName string
	LastMod      int64
	Totals       Totals
	ROA          ROA
	Payments     Payments
	Warning      string
}

// Totals storage
type Totals struct {
	Sales          string
	Tax            string
	Total          string
	InvCount       string
	InvPerHr       string
	SalesPerHr     string
	SalesPerInv    string
	FirstTransTime string
	LastTransTime  string
	Cost           string
	Profit         string
}

// Transaction Structure for CSVData
type Transaction struct {
	TransID  string `csv:"Trans#"`
	Date     string `csv:"Date"`
	Time     string `csv:"Time"`
	SubTotal string `csv:"SubTot"`
	Tax      string `csv:"Tax"`
	Total    string `csv:"Total"`
	Cost     string `csv:"Cost"`
	Ship     string `csv:"Ship"`
	Discount string `csv:"Disc"`
	NumTrans string `csv:"#Trns"`
}

// Server Parameters Structure
type Server struct {
	BusinessName      string `json:"businessname"`
	TransactionsFile  string `json:"transfile"`
	DashboardTemplate string `json:"dashboard_template"`
	Port              string `json:"port"`
}

// Arguments Structure
type Arguments struct {
	Config string
}

// ROA Structure
type ROA struct {
}

// Payments Structure
type Payments struct {
}
