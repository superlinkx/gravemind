package main

import "github.com/shopspring/decimal"

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
	TransError      string
	ZeroMoney       string
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
	SupportEmail      string `json:"supportemail"`
	ToEmail           string `json:"toemail"`
	EmailPass         string `json:"emailpass"`
}

// Arguments Structure
type Arguments struct {
	Config string
}

// Payment Raw Structure
type Payment struct {
	PaymentNumber string `csv:"#"`
	PaymentName   string `csv:"Payment Name"`
	TotalPayments string `csv:"Count"`
	TotalAmount   string `csv:"Total"`
}

// Payments contains raw extracted data
type Payments struct {
	PaidOut    PaymentData
	CCPayments PaymentData
	CCPulled   PaymentData
	Cash       PaymentData
	Check      PaymentData
	Rewards    PaymentData
	Other      PaymentData
	OnAccount  PaymentData
	GiftCard   PaymentData
	ARAdj      PaymentData
}

// PaymentData Structure
type PaymentData struct {
	Num   int
	Total decimal.Decimal
}

// PaymentTotals contains totalling for payments/roas
type PaymentTotals struct {
	NumPaidOut       int
	NumNetCash       int
	NumCashCheckDep  int
	NumCCDep         int
	NumTotalDailyDep int
	NumTotalNonCash  int
	PaidOut          decimal.Decimal
	NetCashDep       decimal.Decimal
	CashCheckDep     decimal.Decimal
	CCDep            decimal.Decimal
	TotalDailyDep    decimal.Decimal
	TotalNonCash     decimal.Decimal
	GrandTotal       decimal.Decimal
	ROAPay           decimal.Decimal
	ROAARAdj         decimal.Decimal
	PaymentsSales    decimal.Decimal
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
	OnAccount        string
	NumOnAccount     string
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
	ROAOnAccount     string
	ROAARAdj         string
	PaymentsSales    string
}
