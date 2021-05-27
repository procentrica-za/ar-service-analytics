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

type RenewalProfileDetails struct {
	Rul                float32 `json:"rulyears"`
	Costopeningbalance float32 `json:"costopeningbalance"`
	CRC                float32 `json:"replacementcost"`
}

type RiskCriticality struct {
	Consequence string `json:"Consequence"`
	Likelyhood  string `json:"likelyhood"`
	CRC         string `json:"replacementcost"`
}

type RiskCriticalityDD struct {
	Name          string `json:"a"`
	Consequence   string `json:"b"`
	Likelyhood    string `json:"c"`
	Crc           string `json:"d"`
	Description   string `json:"e"`
	Cuname        string `json:"f"`
	Typename      string `json:"g"`
	SerialNo      string `json:"h"`
	Extent        string `json:"i"`
	Drc           string `json:"j"`
	Cost          string `json:"k"`
	CarryingValue string `json:"l"`
	TakeOnDate    string `json:"m"`
	Rulyears      string `json:"n"`
	TypeFriendly  string `json:"o"`
	Size          string `json:"p"`
}

type RiskCriticalityDetails struct {
	Consequence       string  `json:"consequence,omititempty"`
	ConsequenceWeight string  `json:"cweight,omititempty"`
	Likelyhood        string  `json:"likelyhood,omititempty"`
	LikelyhoodWeight  string  `json:"lweight,omititempty"`
	CRC               float32 `json:"replacementcost,omititempty"`
	Score             string  `json:"score,omititempty"`
}

type RiskCriticalityFilter struct {
	Name          string  `json:"name,omititempty"`
	Consequence   string  `json:"consequence,omititempty"`
	Likelyhood    string  `json:"likelyhood,omititempty"`
	Crc           float32 `json:"crc,omititempty"`
	Description   string  `json:"description,omititempty"`
	Cuname        string  `json:"cuname,omititempty"`
	Typename      string  `json:"typename,omititempty"`
	SerialNo      string  `json:"serialno,omititempty"`
	Extent        float32 `json:"extent,omititempty"`
	Drc           float32 `json:"drc,omititempty"`
	Cost          float32 `json:"cost,omititempty"`
	CarryingValue float32 `json:"carryingvalue,omititempty"`
	TakeOnDate    string  `json:"takeondate,omititempty"`
	Rulyears      float32 `json:"rulyears,omititempty"`
	TypeFriendly  string  `json:"typefriendly,omititempty"`
	Size          float32 `json:"size,omititempty"`
}

type ReplacementByCondition struct {
	RULYears  string `json:"rulyearsears"`
	Condition string `json:"condition"`
	CRC       string `json:"replacementcost"`
}

type FlattenedHierarchyFilter struct {
	NodeID      string `json:"nodeid,omitempty"`
	Likelyhood  string `json:"likelyhood,omitempty"`
	Consequence string `json:"consequence,omitempty"`
	AssettypeID string `json:"assettypeid,omitempty"`
	Rulyears    int    `json:"rulyears,omitempty"`
	Condition   string `json:"condition,omitempty"`
}

type PortfolioListCost struct {
	Portfolio []PortfolioDD `json:"levels"`
	Cost      float32       `json:"cost"`
}

type PortfolioListHigherCost struct {
	PortfolioHigher []PortfolioListCost `json:"portfolio"`
}

type PortfolioDD struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PortfolioList struct {
	Portfolio []PortfolioDD `json:"levels"`
	CRC       float32       `json:"crc"`
}

type PortfolioListHigher struct {
	PortfolioHigher []PortfolioList `json:"portfolio"`
}

type YearReplacementDetails struct {
	Rul float32 `json:"rulyears"`
	CRC float32 `json:"replacementcost"`
}

type ReplacementByConditionDetails struct {
	RULYears  float32 `json:"rulyears"`
	Condition string  `json:"condition"`
	CRC       float32 `json:"replacementcost"`
}

type AFVConditionDetails struct {
	Crc                  float32 `json:"crc"`
	Drc                  float32 `json:"drc"`
	Assetflexvaluesorted string  `json:"assetflexvaluesorted"`
	Recordcount          float32 `json:"recordcount"`
}
