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
	Crc                  string `json:"crc"`
	Drc                  string `json:"drc"`
	Assetflexvaluesorted string `json:"Assetflexvaluesorted"`
	Recordcount          string `json:"Recordcount"`
}

type Portfolio struct {
	AssetTypeLevel1 string `json:"assettypelvl1"`
	AssetTypeLevel2 string `json:"assettypelvl2"`
	AssetTypeLevel3 string `json:"assettypelvl3"`
	AssetTypeLevel4 string `json:"assettypelvl4"`
	AssetTypeLevel5 string `json:"assettypelvl5"`
	AssetTypeLevel6 string `json:"assettypelvl6"`
	CRC             string `json:"crc"`
}

type YearReplacement struct {
	AssetTypeLevel1 string `json:"assettypelvl1"`
	AssetTypeLevel2 string `json:"assettypelvl2"`
	AssetTypeLevel3 string `json:"assettypelvl3"`
	AssetTypeLevel4 string `json:"assettypelvl4"`
	AssetTypeLevel5 string `json:"assettypelvl5"`
	AssetTypeLevel6 string `json:"assettypelvl6"`
	Rul             string `json:"rulyears"`
	CRC             string `json:"replacementcost"`
}

type RenewalProfile struct {
	Rul                string `json:"rulyears"`
	Costopeningbalance string `json:"costopeningbalance"`
	CRC                string `json:"replacementcost"`
}

type RiskCriticality struct {
	Consequence string `json:"Consequence"`
	Likelyhood  string `json:"likelyhood"`
	CRC         string `json:"replacementcost"`
}

type RiskCriticalityDD struct {
	Name        string `json:"a"`
	Consequence string `json:"b"`
	Likelyhood  string `json:"c"`
	CRC         string `json:"d"`
}

type ReplacementByCondition struct {
	RULYears  string `json:"rulyearsears"`
	Condition string `json:"condition"`
	CRC       string `json:"replacementcost"`
}
