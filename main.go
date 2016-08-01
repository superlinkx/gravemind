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
	"strconv"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
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

func handler(w http.ResponseWriter, r *http.Request) {
	transactions := []Transaction{}
	var totals Totals
	var statInfo os.FileInfo
	var lastMod int64
	var p *Page
	emptyFile := false
	misalignedFile := false

	f, err := os.OpenFile(settings.TransactionsFile, os.O_RDONLY, os.ModePerm)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		emptyFile = true
	}

	defer f.Close()

	if !emptyFile {
		var e error
		statInfo, e = f.Stat()
		if e != nil {
			fmt.Printf("Error getting statinfo: %v\n", err)
			lastMod = 0
			emptyFile = true
		} else {
			lastMod = statInfo.ModTime().Unix()
			currentTime := time.Now().Unix()

			lastMod = (currentTime - lastMod) / 60
		}

		if err = gocsv.UnmarshalFile(f, &transactions); err != nil {
			fmt.Printf("Error unmarshalling file: %v\n", err)
			misalignedFile = true
		}

		if !misalignedFile {
			totals = calcTotals(transactions)

			p = &Page{
				BusinessName: settings.BusinessName,
				LastMod:      lastMod,
				Totals:       totals,
				Warning:      "",
			}
		}
	}

	if emptyFile || misalignedFile {
		totals = emptyTotals()
		p = &Page{
			BusinessName: settings.BusinessName,
			LastMod:      lastMod,
			Totals:       totals,
			Warning:      "Empty set. Either no data yet or problem with transaction import. Contact Sysadmin if persists.",
		}
	}

	t, err := template.ParseFiles(settings.DashboardTemplate)
	if err != nil {
		http.Error(w, "Error parsing template file", 500)
		fmt.Printf("Error parsing template file: %v\n", err)
		return
	}

	t.Execute(w, p)
}

func emptyTotals() Totals {
	emptyTime, _ := time.Parse("3:04 pm", "12:00 am")
	emptyTimeString := emptyTime.Format("3:04 pm")
	return Totals{
		Sales:          "0.00",
		Tax:            "0.00",
		Total:          "0.00",
		InvCount:       "0",
		InvPerHr:       "0.000",
		SalesPerHr:     "0.00",
		SalesPerInv:    "0.00",
		FirstTransTime: emptyTimeString,
		LastTransTime:  emptyTimeString,
		Cost:           "0.00",
		Profit:         "0.00",
	}
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
		Sales:          commaString(salesTotal, 2),
		Tax:            commaString(taxTotal, 2),
		Total:          commaString(grandTotal, 2),
		InvCount:       commaString(invCount, 0),
		InvPerHr:       commaString(invPerHr, 4),
		SalesPerHr:     commaString(salesPerHr, 2),
		SalesPerInv:    commaString(salesPerInv, 2),
		FirstTransTime: leastTime.Format("3:04 pm"),
		LastTransTime:  greatestTime.Format("3:04 pm"),
		Cost:           commaString(costTotal, 2),
		Profit:         commaString(profit, 2),
	}
}

func commaString(dec decimal.Decimal, places int32) string {
	floatConv, _ := strconv.ParseFloat(dec.StringFixed(places), 64)
	return humanize.Commaf(floatConv)
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
