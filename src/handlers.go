package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func (s *Server) handleaddsheetdata() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle add sheet data has been called")
		sheetData := InsertSheet{}
		err := json.NewDecoder(r.Body).Decode(&sheetData)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Improper cell data provided")
			return
		}

		b, err := ioutil.ReadFile("credentials.json")
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
			fmt.Println("Unable to read client secret file: %v", err)
		}
		// If modifying these scopes, delete your previously saved credentials
		// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
		config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
			fmt.Println("Unable to parse client secret file to config: %v", err)
		}
		client := getClient(config)
		srv, err := sheets.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve Sheets Client %v", err)
			fmt.Println("Unable to retrieve Sheets Client %v", err)
		}
		spreadsheetID := sheetData.SheetID
		writeRange := "A2"
		var vr sheets.ValueRange

		//csv try
		csvFile, err := os.Open("data.csv")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened CSV file")
		defer csvFile.Close()

		csvLines, err := csv.NewReader(csvFile).ReadAll()
		if err != nil {
			fmt.Println(err)
		}
		for _, line := range csvLines {
			asset := Application{
				Name:               line[0],
				Description:        line[1],
				SerialNo:           line[2],
				Size:               line[3],
				Type:               line[4],
				Class:              line[5],
				Dimension1Val:      line[6],
				Dimension2Val:      line[7],
				Dimension3Val:      line[8],
				Dimension4Val:      line[9],
				Dimension5Val:      line[10],
				Dimension6Val:      line[11],
				Extent:             line[12],
				ExtentConfidence:   line[13],
				DeRecognitionvalue: line[14],
			}
			myval := []interface{}{asset.Name, asset.Description, asset.SerialNo, asset.Size, asset.Type, asset.Class, asset.Dimension1Val, asset.Dimension2Val, asset.Dimension3Val, asset.Dimension4Val, asset.Dimension5Val, asset.Dimension6Val, asset.Extent, asset.ExtentConfidence, asset.DeRecognitionvalue}
			vr.Values = append(vr.Values, myval)
		}

		_, err = srv.Spreadsheets.Values.Update(spreadsheetID, writeRange, &vr).ValueInputOption("RAW").Do()

		var sheetResponse InsertSheetResponse

		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet. %v", err)
			sheetResponse.Message = "There was an issue whilst inserting data into the sheet"
		}

		sheetResponse.Message = "Cell successfully updated!"
		fmt.Println("Insert into google sheets was successfull")

		js, jserr := json.Marshal(sheetResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from Google shheet result to populate sheet")
			return
		}
		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
