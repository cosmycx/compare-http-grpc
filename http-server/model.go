package main

type uploadResponse struct {
	UploadDur          string               `json:"uploadDuration"`
	ScanCsvDur         string               `json:"scanCsvDuration"`
	ValidationDur      string               `json:"validationDuration"`
	ValidationResponse []validationResponse `json:"validationResponse"`
} // .uploadResponse

type validationRequest struct {
	Metadata metaData     `json:"metadata"`
	Current  currentData  `json:"current"`
	Previous previousData `json:"previous"`
} // .validationRequest

type currentData map[string][]map[int][]string
type previousData map[int][]string

type metaData struct {
	OrgCode int `json:"orgcode"`
} // .metadata

type validationResponse struct {
	SeverityType     string
	IssueType        string
	IssueScopeType   string
	RowNumber        int
	FieldName        int
	FieldValue       string
	IssueDescription string
} // .validationRequest
