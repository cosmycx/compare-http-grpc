package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)

	log.Println("Starting server on :9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalln("Error starting the server on port 9000")
	} // .if
} // .main

func home(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		log.Println("Method not accepted")
		http.Error(w, "method not accepted", http.StatusMethodNotAllowed)
		return
	} // .if
	log.Println("Received request")

	var valReq validationRequest
	json.NewDecoder(r.Body).Decode(&valReq)
	log.Printf("Lines received to validate: %v\n", len(valReq.Current))
	valResp := validationResponse{} //validate(valReq)

	log.Println("Responded")

	json.NewEncoder(w).Encode(valResp)
} // .home

func validate(valReq validationRequest) []validationResponse {
	//log.Println(valReq)

	var errReport []validationResponse

	for rowNo := 0; rowNo < len(valReq.Current); rowNo += 100 {

		var oneErr validationResponse

		oneErr = validationResponse{
			SeverityType:     "severitytype",
			IssueType:        "issuetype",
			IssueScopeType:   "scopetype",
			RowNumber:        rowNo,
			FieldName:        20,
			FieldValue:       "fieldvalue",
			IssueDescription: "issuedescription",
		} // .valResp

		errReport = append(errReport, oneErr)
	} // .for
	log.Printf("Error report lines: %v\n", len(errReport))

	return errReport
} // .validate
