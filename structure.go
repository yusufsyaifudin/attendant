package attendant

// ReplyType defines reply type, whether it Error (ReplyError) or other type
type ReplyType string

const (
	// ReplyError is type error
	ReplyError ReplyType = "Error"
)

// ReplyErrorStructure is error structure when error response written to user
type ReplyErrorStructure struct {
	Code    string `json:"code,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}

// ReplyStructure is common JSON response structure if request is success
type ReplyStructure struct {
	Error *ReplyErrorStructure `json:"error,omitempty"`
	Type  ReplyType            `json:"type,omitempty"`
	Data  interface{}          `json:"data"`
}
