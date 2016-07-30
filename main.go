package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/gocarina/gocsv"
)

//Global for settings
var settings Server

// Page Structure
type Page struct {
	BusinessName string
	LastMod      int64
	Totals       Totals
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

func handler(w http.ResponseWriter, r *http.Request) {
	transactions := []Transaction{}

	f, err := os.OpenFile(settings.TransactionsFile, os.O_RDONLY, os.ModePerm)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer f.Close()

	statInfo, err := f.Stat()
	if err != nil {
		fmt.Printf("Error getting statinfo: %v\n", err)
		return
	}
	lastMod := statInfo.ModTime().Unix()
	currentTime := time.Now().Unix()

	lastMod = (currentTime - lastMod) / 60

	if err = gocsv.UnmarshalFile(f, &transactions); err != nil {
		fmt.Printf("Error unmarshalling file: %v\n", err)
		return
	}

	totals := calcTotals(transactions)

	p := &Page{
		BusinessName: settings.BusinessName,
		LastMod:      lastMod,
		Totals:       totals,
	}

	t, err := template.ParseFiles(settings.DashboardTemplate)
	if err != nil {
		fmt.Printf("Error parsing template file: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, p)
}

func calcTotals(transactions []Transaction) Totals {
	var invCount decimal.Decimal
	var salesTotal decimal.Decimal
	var taxTotal decimal.Decimal
	var grandTotal decimal.Decimal
	var costTotal decimal.Decimal

	var sales decimal.Decimal
	var subtot decimal.Decimal
	var discount decimal.Decimal
	var ship decimal.Decimal
	var tax decimal.Decimal
	var total decimal.Decimal
	var cost decimal.Decimal

	greatestTime, _ := time.Parse("3:04 pm", "12:00 am")
	leastTime, _ := time.Parse("3:04 pm", "11:59 pm")

	invCount = decimal.NewFromFloat(float64(len(transactions)))

	for _, transaction := range transactions {
		subtot, _ = decimal.NewFromString(transaction.SubTotal)
		discount, _ = decimal.NewFromString(transaction.Discount)
		ship, _ = decimal.NewFromString(transaction.Ship)
		sales = subtot.Sub(discount).Add(ship)
		tax, _ = decimal.NewFromString(transaction.Tax)
		total, _ = decimal.NewFromString(transaction.Total)
		cost, _ = decimal.NewFromString(transaction.Cost)

		salesTotal = salesTotal.Add(sales)
		taxTotal = taxTotal.Add(tax)
		grandTotal = grandTotal.Add(total)
		costTotal = costTotal.Add(cost)

		transTime, _ := time.Parse("3:04 pm", strings.TrimSpace(transaction.Time))

		if transTime.After(greatestTime) {
			greatestTime = transTime
		}

		if transTime.Before(leastTime) {
			leastTime = transTime
		}
	}

	totalTime := greatestTime.Sub(leastTime).Hours()

	invPerHr := invCount.Div(decimal.NewFromFloat(totalTime))
	salesPerHr := salesTotal.Div(decimal.NewFromFloat(totalTime))
	salesPerInv := salesTotal.Div(invCount)

	profit := salesTotal.Sub(costTotal)

	return Totals{
		Sales:          salesTotal.StringFixed(2),
		Tax:            taxTotal.StringFixed(2),
		Total:          grandTotal.StringFixed(2),
		InvCount:       invCount.StringFixed(0),
		InvPerHr:       invPerHr.StringFixed(4),
		SalesPerHr:     salesPerHr.StringFixed(2),
		SalesPerInv:    salesPerInv.StringFixed(2),
		FirstTransTime: leastTime.Format("3:04 pm"),
		LastTransTime:  greatestTime.Format("3:04 pm"),
		Cost:           costTotal.StringFixed(2),
		Profit:         profit.StringFixed(2),
	}
}

func getArgs(params *Arguments) error {
	flag.StringVar(&params.Config, "config", "/etc/gravemind/gravemind.json", "Location of config file (default is /etc/gravemind/gravemind.json)")

	flag.Parse()

	return nil
}

func readConfig(config string, server *Server) error {
	file, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return err
	}

	if err := json.Unmarshal(file, &server); err != nil {
		fmt.Printf("JSON Unmarshalling error: %v\n", err)
		return err
	}

	return nil
}

func main() {
	var args Arguments
	if err := getArgs(&args); err != nil {
		fmt.Printf("Error getting arguments: %v\n", err)
	}

	if err := readConfig(args.Config, &settings); err != nil {
		fmt.Printf("Error reading configuration: %v\n", err)
	}

	address := ":" + settings.Port

	fmt.Println(address)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
