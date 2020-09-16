package main

import "github.com/gorilla/mux"

//create structs for JSON objects recieved and responses

type InsertSheet struct {
	SheetID string `json:"sheetid"`
}

type InsertSheetResponse struct {
	Message string `json:"message"`
}

//touter service struct
type Server struct {
	router *mux.Router
}

//config struct
type Config struct {
	SHEETHost string
	SHEETPort string
}

type Application struct {
	Name               string
	Description        string
	SerialNo           string
	Size               string
	Type               string
	Class              string
	Dimension1Val      string
	Dimension2Val      string
	Dimension3Val      string
	Dimension4Val      string
	Dimension5Val      string
	Dimension6Val      string
	Extent             string
	ExtentConfidence   string
	DeRecognitionvalue string
}
