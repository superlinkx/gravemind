package main

// Page Structure
type Page struct {
	BusinessName    string
	TransLastMod    int64
	ROALastMod      int64
	PaymentsLastMod int64
	Totals          Totals
	Payments        DispPayments
	TransWarning    string
	ROAWarning      string
	PaymentsWarning string
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
	ROAFile           string `json:"roafile"`
	PaymentsFile      string `json:"paymentsfile"`
	DashboardTemplate string `json:"dashboard_template"`
	Port              string `json:"port"`
}

// Arguments Structure
type Arguments struct {
	Config string
}

// Payments Structure
type Payments struct {
	PaymentNumber string `csv:"#"`
	PaymentName   string `csv:"Payment Name"`
	TotalPayments string `csv:"Count"`
	TotalAmount   string `csv:"Total"`
}

// DispPayments Structure for displayable Payments
type DispPayments struct {
	Cash             string
	NumCash          string
	PaidOut          string
	NumPaidOut       string
	NetCashDep       string
	Checks           string
	NumChecks        string
	CashCheckDep     string
	NumCashCheckDep  string
	CCDep            string
	NumCCDep         string
	TotalDailyDep    string
	NumTotalDailyDep string
	AR               string
	NumAR            string
	RewardsRedeem    string
	NumRewardsRedeem string
	ARAdj            string
	NumARAdj         string
	GiftCards        string
	NumGiftCards     string
	Other            string
	NumOther         string
	TotalNonCash     string
	NumTotalNonCash  string
	GrandTotal       string
	ROAPaidOut       string
	ROAAR            string
	ROARewardsRedeem string
	ROAARAdj         string
	ROAGiftCards     string
	ROAOther         string
	PaymentsSales    string
}
