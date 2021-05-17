package main

import (
	"bytes"
	"compress/gzip"
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
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetPortfolio() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset portfolio has Been Called...")

		nodeid := r.URL.Query().Get("nodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/portfolio?nodeid=" + nodeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get portfolio")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get portfolio"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get portfolio" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []Portfolio{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetYearReplacement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Year Replacement has Been Called...")

		nodeid := r.URL.Query().Get("nodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/yearreplacement?nodeid=" + nodeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get year replacement")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get year replacement"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get year replacement" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []YearReplacement{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetRenewalProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Renewal profile has Been Called...")

		nodeid := r.URL.Query().Get("nodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/renewalprofile?nodeid=" + nodeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get Renewal profile")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get Renewal profile"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get Renewal profile" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []RenewalProfile{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetRiskCriticality() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Risk Criticality has Been Called...")

		nodeid := r.URL.Query().Get("nodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/riskcriticality?nodeid=" + nodeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get Risk Criticality")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get Risk Criticality"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get Risk Criticality" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []RiskCriticality{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetReplacementByCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get replacement by condition has Been Called...")

		nodeid := r.URL.Query().Get("nodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/replacementbycondition?nodeid=" + nodeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get Risk Replacement by condition")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get Risk Replacement by condition"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get Risk Replacement by condition" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []ReplacementByCondition{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetRiskCriticalityDrillDown() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Risk Criticality DD has Been Called...")

		nodeid := r.URL.Query().Get("nodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/riskcriticalitydd?nodeid=" + nodeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get Risk Criticality")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get Risk Criticality"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get Risk Criticality" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []RiskCriticalityDD{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetRiskCriticalityDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Risk Criticality has Been Called...")
		///get JSON payload
		filterparams := FlattenedHierarchyFilter{}
		err := json.NewDecoder(r.Body).Decode(&filterparams)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error with hierarchy filter parameters")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(filterparams)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/riskcriticalitydetails", "application/json", bytes.NewBuffer(requestByte))
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve nodefunclocs spatial filtered.")
			return
		}

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get nodefunclocs spatial filtered")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get nodefunclocs spatial filtered"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get nodefunclocs spatial filtered" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []RiskCriticalityDetails{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetRiskCriticalityFilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Risk Criticality filter has Been Called...")

		nodeid := r.URL.Query().Get("nodeid")
		likelyhood := r.URL.Query().Get("likelyhood")
		consequence := r.URL.Query().Get("consequence")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Node ID not properly provided in URL")
			fmt.Println("Node ID not properly provided in URL")
			return
		}

		if likelyhood == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "likelyhood not properly provided in URL")
			fmt.Println("likelyhood not properly provided in URL")
			return
		}

		if consequence == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "consequence not properly provided in URL")
			fmt.Println("consequence not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/riskcriticalityfilter?nodeid=" + nodeid + "&&likelyhood=" + likelyhood + "&&consequence=" + consequence)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get Risk Criticality")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get Risk Criticality"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get Risk Criticality" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []RiskCriticalityFilter{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetPortfolioFilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset portfolio filtered has Been Called...")

		///get JSON payload
		filterparams := FlattenedHierarchyFilter{}
		err := json.NewDecoder(r.Body).Decode(&filterparams)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error with hierarchy filter parameters")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(filterparams)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/portfoliofiltered", "application/json", bytes.NewBuffer(requestByte))
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve nodefunclocs spatial filtered.")
			return
		}

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get nodefunclocs spatial filtered")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get nodefunclocs spatial filtered"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get nodefunclocs spatial filtered" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		portfolioListHigher := PortfolioListHigher{}
		portfolioListHigher.PortfolioHigher = []PortfolioList{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&portfolioListHigher)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get location response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(portfolioListHigher)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetRenewalProfileDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Renewal profile details has Been Called...")

		///get JSON payload
		filterparams := FlattenedHierarchyFilter{}
		err := json.NewDecoder(r.Body).Decode(&filterparams)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error with hierarchy filter parameters")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(filterparams)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/renewalprofiledetails", "application/json", bytes.NewBuffer(requestByte))
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve nodefunclocs spatial filtered.")
			return
		}

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get nodefunclocs spatial filtered")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get nodefunclocs spatial filtered"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get nodefunclocs spatial filtered" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := []RenewalProfileDetails{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		// create header
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.WriteHeader(200)
		// Gzip data
		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(assetsList)
		gz.Close()
	}
}

func (s *Server) handleGetPortfolioFilterCost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset portfolio cost filtered has Been Called...")

		///get JSON payload
		filterparams := FlattenedHierarchyFilter{}
		err := json.NewDecoder(r.Body).Decode(&filterparams)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error with hierarchy filter parameters")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(filterparams)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/portfoliofilteredcost", "application/json", bytes.NewBuffer(requestByte))
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve nodefunclocs spatial filtered.")
			return
		}

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get nodefunclocs spatial filtered")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get nodefunclocs spatial filtered"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get nodefunclocs spatial filtered" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		portfolioListHigher := PortfolioListHigherCost{}
		portfolioListHigher.PortfolioHigher = []PortfolioListCost{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&portfolioListHigher)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get location response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(portfolioListHigher)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}
