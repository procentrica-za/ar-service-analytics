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

func (s *Server) handleanalyseassets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle analyse assets has been called")

		//Get Asset ID from URL
		assettypeid := r.URL.Query().Get("assettypeid")
		sheetid := r.URL.Query().Get("sheetid")

		//Check if Asset ID provided is null
		if assettypeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Asset Type ID not properly provided in URL")
			fmt.Println("Asset Type ID not proplery provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/analyseassets?assettypeid=" + assettypeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve advertisement information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get asset data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get asset data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []AssetRegisterResponse{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get Messages response ")
			return
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(assetsList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}
		// Unmarshal JSON data
		var jsonData []Application
		err := json.Unmarshal([]byte(js), &jsonData)

		if err != nil {
			fmt.Println(err)
		}
		//generate UUID as a container name
		fileUUID, _ := newUUID()
		fileName := fileUUID + ".csv"

		//Create CSV file
		csvFile, err := os.Create(fileName)
		fmt.Println("Created file")
		if err != nil {
			fmt.Println(err)
		}
		defer csvFile.Close()

		writer := csv.NewWriter(csvFile)

		//Add data into rows
		for _, usance := range jsonData {
			var row []string
			row = append(row, usance.Name)
			row = append(row, usance.Description)
			row = append(row, usance.SerialNo)
			row = append(row, usance.Size)
			row = append(row, usance.Type)
			row = append(row, usance.Class)
			row = append(row, usance.Dimension1Val)
			row = append(row, usance.Dimension2Val)
			row = append(row, usance.Dimension3Val)
			row = append(row, usance.Dimension4Val)
			row = append(row, usance.Dimension5Val)
			row = append(row, usance.Dimension6Val)
			row = append(row, usance.Extent)
			row = append(row, usance.ExtentConfidence)
			row = append(row, usance.DeRecognitionvalue)
			writer.Write(row)
		}
		fmt.Println("Populated CSV")
		fmt.Println(fileName)

		// flush the writer
		writer.Flush()
		//Open file
		Openfile, err := os.Open(fileName)

		//return to beginning  of array
		Openfile.Seek(0, 0)

		var dlterror = os.Remove(fileName)
		if dlterror != nil {
			fmt.Println(dlterror)
		}
		fmt.Println("File has been deleted -->" + fileName)
		//Send the file

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
		spreadsheetID := sheetid
		writeRange := "A2"
		var vr sheets.ValueRange

		//csv try

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened CSV file")
		defer Openfile.Close()

		csvLines, err := csv.NewReader(Openfile).ReadAll()
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

		sheetResponse.Message = "Google sheets has been successfully updated"
		fmt.Println("Insert into google sheets was successfull")

		js1, jserr := json.Marshal(sheetResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from Google shheet result to populate sheet")
			return
		}
		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js1)
	}
}
