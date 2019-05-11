package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func getServerURL() string {
	var url string
	url = os.Getenv("VALIDATION_PATH")
	if len(url) == 0 {
		url = "http://localhost:9000/"
	}
	return url
} // .getServerURL

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/validate", validate)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error starting the server on port 8080")
	} // .if
} // .main

func validate(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")

	// --------------------------------------------------------------
	start := time.Now()
	r.ParseMultipartForm(32 << 20)
	updfile, _, err := r.FormFile("updfile")
	if err != nil {
		log.Printf("Error at upload file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} // .if
	defer updfile.Close()
	uploadDur := time.Since(start)
	log.Printf("upload duration: %v\n", uploadDur)

	// --------------------------------------------------------------
	start = time.Now()
	data, orgCode, err := scanFile(updfile)
	if err != nil {
		log.Printf("Error at scan file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} // .if

	scanCsvDur := time.Since(start)
	log.Printf("scan duration: %v\n", scanCsvDur)

	// --------------------------------------------------------------
	start = time.Now()

	validationRequest := validationRequest{
		Metadata: metaData{
			OrgCode: orgCode,
		},
		Current:  data,
		Previous: nil,
	} // .validationRequest

	// sending to server service

	validationRequestBytes, err := json.Marshal(&validationRequest)
	if err != nil {
		log.Printf("Error at converting to JSON: %v\n", err)
	} // .if
	resp, err := http.Post(getServerURL(), "application/json", bytes.NewBuffer(validationRequestBytes))
	if err != nil {
		log.Printf("Error receive response from validation: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var valResp []validationResponse
	err = json.NewDecoder(resp.Body).Decode(&valResp)

	validationDur := time.Since(start)
	log.Printf("server response duration: %v\n", validationDur)

	uploadResponse := uploadResponse{
		UploadDur:          fmtDur(uploadDur),
		ScanCsvDur:         fmtDur(scanCsvDur),
		ValidationDur:      fmtDur(validationDur),
		ValidationResponse: valResp,
	} // .uploadResponse

	json.NewEncoder(w).Encode(&uploadResponse)
} // .validate

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
    <html>
    <head>
        <title>validator</title>
    </head>
    <body>
    <form action="/validate" enctype="multipart/form-data" method="post">
        <input type="file" name="updfile">
        <input type="submit" name="submit">
    </form>
    </body>
    </html>`)
} // .home

func fmtDur(d time.Duration) string {
	d = d.Round(time.Microsecond)

	s := d / time.Second
	ms := d / time.Millisecond

	return fmt.Sprintf("%02ds:%02dms", s, ms)
} // .fmtDur
