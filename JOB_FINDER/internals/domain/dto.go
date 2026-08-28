package domain

type Request struct {
	Text     string `json:"text"`
	FileURI  string `json:"fileUri"`
	MIMEType string `json:"MIMEType"`
}
