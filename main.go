package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

//Global for settings
var settings Server

// Page Structure
type Page struct {
	BusinessName string
	DateTime     string
	Data         []Data
}

// Data Structure for CSVData
type Data struct {
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
	data := []Data{}

	f, _ := os.OpenFile(settings.TransactionsFile, os.O_RDONLY, os.ModePerm)

	defer f.Close()

	err := gocsv.UnmarshalFile(f, &data)

	if err != nil {
		fmt.Println(err)
		return
	}

	datetime := time.Now().Format(time.UnixDate)

	p := &Page{
		BusinessName: settings.BusinessName,
		DateTime:     datetime,
		Data:         data,
	}

	t, _ := template.ParseFiles("dashboard.html")
	t.Execute(w, p)
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

func main() {
	settings = getArgs()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
