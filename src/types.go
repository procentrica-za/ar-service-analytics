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
	CRUDHost  string
	CRUDPort  string
}

type AFVCondition struct {
	Assetname            string `json:"assetname,omitempty"`
	Crc                  string `json:"crc,omitempty"`
	Drc                  string `json:"drc,omitempty"`
	Assetflexvaluesorted string `json:"Assetflexvaluesorted,omitempty"`
}
