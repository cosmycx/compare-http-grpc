package main

import (
	context "context"
	"fmt"
	"log"
	"net/http"
	"time"

	grpc "google.golang.org/grpc"
)

type server struct {
	grpcClient ValidatePBServiceClient
}

func main() {

	fmt.Println("Started client, open localhost:8080 and upload file ...")
	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer grpcConn.Close()

	grpcClient := NewValidatePBServiceClient(grpcConn)

	s := server{
		grpcClient: grpcClient,
	} // .s

	http.HandleFunc("/", s.home)
	http.HandleFunc("/validate", s.validate)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Can't start server")
	} // .err

} // .main

func (s *server) doClientStreaming(participants map[string][]map[int][]string) {
	fmt.Println("Starting streaming participants ... ")

	stream, err := s.grpcClient.Validate(context.Background())
	if err != nil {
		log.Fatalf("error while calling Validation: %v", err)
	}

	for participID, dataRows := range participants {

		var rows []*Row

		for _, oneRow := range dataRows {
			var row *Row

			for k, v := range oneRow {
				row = &Row{
					RowNo: int64(k),
					RowData: &RowData{
						Field: v,
					},
				} // .row
			} // .for

			rows = append(rows, row)
		} // .for

		oneParticip := &ValidationPBRequest{
			ParticipID: participID,
			Datarows:   rows,
		} // .oneParticip

		//log.Printf("Sending req for participID: %v\n", participID)
		stream.Send(oneParticip)
		//time.Sleep(5 * time.Millisecond)

	} // .for

	_, err = stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	//log.Printf("Validation Response: %v\n", res)
} // .doClientStreaming

func (s *server) validate(w http.ResponseWriter, r *http.Request) {

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
	data, _, err := scanFile(updfile)
	if err != nil {
		log.Printf("Error at scan file: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} // .if

	scanCsvDur := time.Since(start)
	log.Printf("scan duration: %v\n", scanCsvDur)

	// --------------------------------------------------------------
	start = time.Now()
	// sending to server GRPC service
	s.doClientStreaming(data)

	serverRespDur := time.Since(start)
	log.Printf("response duration: %v\n", serverRespDur)

} // .validate

func (s *server) home(w http.ResponseWriter, r *http.Request) {
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
