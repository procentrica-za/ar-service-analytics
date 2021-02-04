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
	Crc                  string `json:"crc,omitempty"`
	Drc                  string `json:"drc,omitempty"`
	Assetflexvaluesorted string `json:"Assetflexvaluesorted,omitempty"`
	Recordcount          string `json:"Recordcount,omitempty"`
}

type Portfolio struct {
	AssetTypeLevel1 string `json:"assettypelvl1,omitempty"`
	AssetTypeLevel2 string `json:"assettypelvl2,omitempty"`
	AssetTypeLevel3 string `json:"assettypelvl3,omitempty"`
	AssetTypeLevel4 string `json:"assettypelvl4,omitempty"`
	AssetTypeLevel5 string `json:"assettypelvl5,omitempty"`
	AssetTypeLevel6 string `json:"assettypelvl6,omitempty"`
	CRC             string `json:"crc,omitempty"`
}

type YearReplacement struct {
	AssetTypeLevel1 string `json:"assettypelvl1,omitempty"`
	AssetTypeLevel2 string `json:"assettypelvl2,omitempty"`
	AssetTypeLevel3 string `json:"assettypelvl3,omitempty"`
	AssetTypeLevel4 string `json:"assettypelvl4,omitempty"`
	AssetTypeLevel5 string `json:"assettypelvl5,omitempty"`
	AssetTypeLevel6 string `json:"assettypelvl6,omitempty"`
	Rul             string `json:"rulyears,omitempty"`
	CRC             string `json:"replacementcost,omitempty"`
}
