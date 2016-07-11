package main

//Transaction contains the structure of an imported transaction
type Transaction struct {
	Number                    string
	NumberOnly                string
	SubNumber                 string
	Date                      string
	Time                      string
	CustNumber                string
	CustName                  string
	CustOrderNumber           string
	SubTotal                  string
	Discount                  string
	DiscountPer               string
	Taxable                   string
	SalesTax                  string
	Shipping                  string
	Total                     string
	AmountDue                 string
	TaxPer                    string
	TaxTable                  string
	ClerkNumber               string
	RegNum                    string
	StoreNum                  string
	PriceTable                string
	OriginalOrderDate         string
	TransTypeCode             string
	PrevTransTypeCode         string
	SalesPersonNumber         string
	TotalCost                 string
	ContactID                 string
	ReversingTrans            string
	SalesTaxOverrideFlag      string
	NonTaxable                string
	HasBeenPrinted            string
	LineItemSalespersonUsed   string
	ShippingIsTaxable         string
	Ignore1                   string
	Ignore2                   string
	DateManuallyEntered       string
	TransHasBeenReversed      string
	OriginalOrderTime         string
	ShipOnDate                string
	NotBeforeDate             string
	NotAfterDate              string
	CustCode                  string
	CustGroup                 string
	ContactNumber             string
	ContactName               string
	ContactGroup              string
	DeliveredToCode           string
	DiscountTakenAmount       string
	AmountTendered            string
	PaidInFullDateDate        string
	PaidInFullDateTime        string
	NumPayments               string
	FirstPaymentNumber        string
	FirstPaymentType          string
	FillRatioByQty            string
	FillRatioByEntireLineItem string
	FillRatioByPrice          string
	MachineNumber             string
	PrevTransDateDate         string
	PrevTransDateTime         string
	PrevTransNumber           string
	PrevTransSubNumber        string
	CreatedDate               string
	CreatedTime               string
	LastChangedDate           string
	LastChangedTime           string
	NumLineItems              string
	NonSaleTotal              string
	Comment                   string
}

//Payment contains the structure of an imported payment
type Payment struct {
	Date                  string `csv:"Date"`
	Time                  string `csv:"Time"`
	TransNumber           string `csv:"TransNumber"`
	TransNumberOnly       string `csv:"TransNumberOnly"`
	TransSubNumber        string `csv:"TransSubNumber"`
	TransDate             string `csv:"TransDate"`
	CustomerNum           string `csv:"Customer#"`
	PaymentNumber         string `csv:"PaymentNumber"`
	PaymentName           string `csv:"PaymentName"`
	Reference             string `csv:"Reference"`
	ExpDate               string `csv:"ExpDate"`
	AuthCode              string `csv:"AuthCode"`
	Amount                string `csv:"Amount"`
	RegisterNum           string `csv:"Register#"`
	StoreNum              string `csv:"Store#"`
	ClerkNum              string `csv:"Clerk#"`
	PaymentType           string `csv:"PaymentType"`
	TransType             string `csv:"TransType"`
	OrigTransType         string `csv:"OrigTransType"`
	PaymentFlags          string `csv:"PaymentFlags"`
	TransTime             string `csv:"TransTime"`
	TypeFlags             string `csv:"TypeFlags"`
	ROAPayment            string `csv:"ROAPayment"`
	CapturedPayment       string `csv:"CapturedPayment"`
	TransferedPayment     string `csv:"TransferedPayment"`
	ReversingPayment      string `csv:"ReversingPayment"`
	AmountCredit          string `csv:"Amount-Credit"`
	AmountCash            string `csv:"Amount-Cash"`
	AmountOther           string `csv:"Amount-Other"`
	AuthID                string `csv:"AuthID"`
	SwipedCard            string `csv:"SwipedCard"`
	GiftCard              string `csv:"GiftCard"`
	AmountPayment         string `csv:"AmountPayment"`
	Card                  string `csv:"Card"`
	AmountGiftCard        string `csv:"Amount-Gift Card"`
	GiftCardTriggerNumber string `csv:"GiftCardTriggerNumber"`
	GiftCardAction        string `csv:"GiftCardAction"`
	PurgedPayment         string `csv:"PurgedPayment"`
	CreatedDate           string `csv:"Created-Date"`
	CreatedTime           string `csv:"Created-Time"`
	LastChangedDate       string `csv:"LastChanged-Date"`
	LastChangedTime       string `csv:"LastChanged-Time"`
	TermsCode             string `csv:"TermsCode"`
	AmountDue             string `csv:"AmountDue"`
	PaidInFullDate        string `csv:"PaidInFull-Date"`
	PaidInFullTime        string `csv:"PaidInFull-Time"`
	DueDate               string `csv:"DueDate"`
	CardHolderName        string `csv:"CardHolderName"`
	HasSignature          string `csv:"HasSignature"`
	GiftCardBalance       string `csv:"GiftCardBalance"`
	Gratuity              string `csv:"Gratuity"`
	Purchase              string `csv:"Purchase"`
	OpenGratuity          string `csv:"OpenGratuity"`
	HasGratuityPayment    string `csv:"HasGratuityPayment"`
	IsGratuityPayment     string `csv:"IsGratuityPayment"`
	SalesPersonNumber     string `csv:"SalesPersonNumber"`
	CustomerName          string `csv:"CustomerName"`
	Comment               string `csv:" Comment"`
}

//Customer contains the structure of an imported customer
type Customer struct {
	Number             string `csv:"Number"`
	BusinessName       string `csv:"BusinessName"`
	FirstName          string `csv:"FirstName"`
	LastName           string `csv:"LastName"`
	Salutation         string `csv:"Salutation"`
	Title              string `csv:"Title"`
	Phone              string `csv:"Phone"`
	Phone2             string `csv:"Phone2"`
	TaxTable           string `csv:"TaxTable"`
	TaxNumber          string `csv:"TaxNumber"`
	Address            string `csv:"Address"`
	City               string `csv:"City"`
	State              string `csv:"State"`
	Zip                string `csv:"Zip"`
	Note               string `csv:"Note"`
	Class              string `csv:"Class"`
	Flags              string `csv:"Flags"`
	ReqPO              string `csv:"ReqPO"`
	Code               string `csv:"Code"`
	CreditLimit        string `csv:"CreditLimit"`
	PriceTable         string `csv:"PriceTable"`
	DiscountPer        string `csv:"Discount-%"`
	HomePhone          string `csv:"HomePhone"`
	CellPhone          string `csv:"CellPhone"`
	TotalPurchased     string `csv:"TotalPurchased"`
	LastPurchaseDate   string `csv:"LastPurchase-Date"`
	AmountDue          string `csv:"AmountDue"`
	LastOpenItemDate   string `csv:"LastOpenItem-Date"`
	OldFinanceDate     string `csv:"OldFinanceDate"`
	CurrentFinanceDue  string `csv:"CurrentFinanceDue"`
	OldFinanceDue      string `csv:"OldFinanceDue"`
	GracePeriod        string `csv:"GracePeriod"`
	AltAddress         string `csv:"AltAddress"`
	LastPaymentDate    string `csv:"LastPayment-Date"`
	SalesPerson        string `csv:"Sales Person"`
	PaymentNumber      string `csv:"Payment Number"`
	Birthday           string `csv:"Birthday"`
	ExpDate            string `csv:"Exp Date"`
	SecurityLevelSales string `csv:"SecurityLevel-Sales"`
	CreatedDate        string `csv:"Created-Date"`
	LastChangedDate    string `csv:"LastChanged-Date"`
	Age                string `csv:"Age"`
	CreatedTime        string `csv:"Created-Time"`
	LastChangedTime    string `csv:"LastChanged-Time"`
	LastOpenItemTime   string `csv:"LastOpenItem-Time"`
	LastPurchaseTime   string `csv:"LastPurchase-Time"`
	LastPaymentTime    string `csv:"LastPayment-Time"`
	StoreNum           string `csv:"Store#"`
	SecurityLevelView  string `csv:"SecurityLevel-View"`
	LookupName         string `csv:"LookupName"`
	WebID              string `csv:"WebID"`
	Group              string `csv:"Group"`
	Fax                string `csv:"Fax"`
	EMail              string `csv:"EMail"`
	EmailAlt           string `csv:"EmailAlt"`
	County             string `csv:"County"`
	Country            string `csv:"Country"`
	Birthdate          string `csv:"Birthdate"`
	FirstPurchaseDate  string `csv:"FirstPurchaseDate"`
	NumPurchases       string `csv:"NumPurchases"`
	AvgPurchaseAmount  string `csv:"AvgPurchaseAmount"`
	AvgProfitAmount    string `csv:"AvgProfitAmount"`
	OnlineCustomer     string `csv:"OnlineCustomer"`
	PreferredCustomer  string `csv:"PreferredCustomer"`
	ResidentialAddress string `csv:"ResidentialAddress"`
	AvgPayDateDays     string `csv:"AvgPayDate-Days"`
	StatementMethod    string `csv:"StatementMethod"`
	UserShortInfo1     string `csv:"UserShortInfo1"`
	UserShortInfo2     string `csv:"UserShortInfo2"`
	UserInfo1          string `csv:"UserInfo1"`
	UserInfo2          string `csv:"UserInfo2"`
	UserLongInfo1      string `csv:"UserLongInfo1"`
	UserLongInfo2      string `csv:"UserLongInfo2"`
	UserNumber1        string `csv:"UserNumber1"`
	UserNumber2        string `csv:"UserNumber2"`
	UserDate1          string `csv:"UserDate1"`
	UserDate2          string `csv:"UserDate2"`
	UserFlags          string `csv:"UserFlags"`
	Hidden             string `csv:"Hidden"`
	PrintStatements    string `csv:"PrintStatements"`
	EmailStatements    string `csv:"EmailStatements"`
	Comment            string `csv:" Comment"`
}
