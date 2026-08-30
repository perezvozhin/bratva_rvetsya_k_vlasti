package gem_service

import "time"

// note: took it from genai docs as to possible states
const (
	FileStateUnspecified FileState = "STATE_UNSPECIFIED"
	FileStateProcessing  FileState = "PROCESSING"
	FileStateActive      FileState = "ACTIVE"
	FileStateFailed      FileState = "FAILED"
)

type FileState string
type UploadFileStatus struct {
	Name           string    `json:"name,omitempty"`
	MIMEType       string    `json:"mimeType,omitempty"`
	SizeBytes      *int64    `json:"sizeBytes,omitempty,string"`
	CreateTime     time.Time `json:"createTime,omitempty"`
	ExpirationTime time.Time `json:"expirationTime,omitempty"`
	State          FileState `json:"state,omitempty"`
	Message        string    `json:"message,omitempty"`
	FileURI        string    `json:"uri,omitempty"`
}

type StreamResponse struct {
	Chunk string
	Error error
}

type Message struct {
	Role     string `json:"role"` // "user" | "model"
	Text     string `json:"text"`
	FileURI  string `json:"fileUri,omitempty"`
	MIMEType string `json:"mimeType,omitempty"`
}

type ChatHistory struct {
	ChatName  string    `json:"chatName"`
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ChatSummary struct {
	ChatName  string    `json:"chatName"`
	UpdatedAt time.Time `json:"updatedAt"`
}
