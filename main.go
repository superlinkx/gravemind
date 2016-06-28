package main

import (
	"flag"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
	Sales          float64
	Tax            float64
	Total          float64
	InvCount       float64
	InvPerHr       float64
	SalesPerHr     float64
	FirstTransTime string
	LastTransTime  string
	Cost           float64
	Profit         float64
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
	NumTrans string `csv:"#Trns"`
}

// Server Parameters Structure
type Server struct {
	BusinessName     string
	TransactionsFile string
}

func handler(w http.ResponseWriter, r *http.Request) {
	transactions := []Transaction{}

	f, _ := os.OpenFile(settings.TransactionsFile, os.O_RDONLY, os.ModePerm)

	defer f.Close()

	statInfo, _ := f.Stat()
	lastMod := statInfo.ModTime().Unix()
	currentTime := time.Now().Unix()

	lastMod = (currentTime - lastMod) / 60

	err := gocsv.UnmarshalFile(f, &transactions)

	if err != nil {
		fmt.Println(err)
		return
	}

	totals := calcTotals(transactions)

	p := &Page{
		BusinessName: settings.BusinessName,
		LastMod:      lastMod,
		Totals:       totals,
	}

	t, _ := template.ParseFiles("dashboard.html")
	t.Execute(w, p)
	fmt.Println(totals)
}

func calcTotals(transactions []Transaction) Totals {
	var invCount float64
	var salesTotal float64
	var taxTotal float64
	var grandTotal float64
	var costTotal float64

	var sales float64
	var tax float64
	var total float64
	var cost float64

	greatestTime, _ := time.Parse("3:04 pm", "12:00 am")
	leastTime, _ := time.Parse("3:04 pm", "11:59 pm")

	invCount = float64(len(transactions))

	for _, transaction := range transactions {
		sales, _ = strconv.ParseFloat(transaction.SubTotal, 64)
		tax, _ = strconv.ParseFloat(transaction.Tax, 64)
		total, _ = strconv.ParseFloat(transaction.Total, 64)
		cost, _ = strconv.ParseFloat(transaction.Cost, 64)

		salesTotal += sales
		taxTotal += tax
		grandTotal += total
		costTotal += cost

		transTime, _ := time.Parse("3:04 pm", strings.TrimSpace(transaction.Time))

		if transTime.After(greatestTime) {
			greatestTime = transTime
		}

		if transTime.Before(leastTime) {
			leastTime = transTime
		}
	}

	totalTime := greatestTime.Sub(leastTime).Hours()

	invPerHr := invCount / totalTime
	salesPerHr := salesTotal / totalTime

	profit := salesTotal - costTotal

	return Totals{
		Sales:          roundDollars(salesTotal),
		Tax:            roundDollars(taxTotal),
		Total:          roundDollars(grandTotal),
		InvCount:       invCount,
		InvPerHr:       roundDollars(invPerHr),
		SalesPerHr:     roundDollars(salesPerHr),
		FirstTransTime: leastTime.Format("3:04 pm"),
		LastTransTime:  greatestTime.Format("3:04 pm"),
		Cost:           roundDollars(costTotal),
		Profit:         roundDollars(profit),
	}
}

func getArgs() Server {
	var params Server

	flag.StringVar(&params.BusinessName, "businessname", "", "Name for current business")
	flag.StringVar(&params.TransactionsFile, "transfile", "", "Path to transactions.csv")

	flag.Parse()

	if (params.BusinessName == "") || (params.TransactionsFile == "") {
		panic("Required flags: -businessname, -transfile")
	}

	return params
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func preciseRound(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return round(f*shift) / shift
}

func roundDollars(f float64) float64 {
	return preciseRound(f, 2)
}

func main() {
	settings = getArgs()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
