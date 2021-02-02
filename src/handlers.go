package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) handleGetAssetFlexValCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset Flexval Condition Been Called...")

		nodeid := r.URL.Query().Get("nodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/assetflexvalcondition?nodeid=" + nodeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get assetflexvalcondition")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get assetflexvalcondition"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get assetflexvalcondition" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []AFVCondition{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		//convert struct back to JSON.
		js, jserr := json.Marshal(assetsList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
