package smtp

const (
	EmailTypeText = "text"
	EmailTypeHTML = "html"
)

type Email struct {
	Type        string       `json:"type"`
	To          []string     `json:"to"`
	Cc          []string     `json:"cc"`
	Bcc         []string     `json:"bcc"`
	Subject     string       `json:"subject"`
	Body        string       `json:"body"`
	Attachments []Attachment `json:"attachments"`
	// To replace the value in subject and body, set the key in like {{key}} format in subject and body
	Values map[string]string `json:"values"`
}

type Attachment struct {
	Filename string `json:"filename"`
	Data     []byte `json:"data"`
}
