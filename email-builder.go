package smtp

// NewEmail creates a new email builder
// emailType can be "text" or "html"
func NewEmail(emailType string) *subjectBuilder {
	return &subjectBuilder{
		emailType: emailType,
	}
}

type subjectBuilder struct {
	emailType string
}

// Subject sets the subject of the email
func (b *subjectBuilder) Subject(subject string) *bodyBuilder {
	return &bodyBuilder{
		emailType: b.emailType,
		subject:   subject,
	}
}

type bodyBuilder struct {
	emailType string
	subject   string
}

// Body sets the body of the email
func (b *bodyBuilder) Body(body string) *toBuilder {
	return &toBuilder{
		emailType: b.emailType,
		subject:   b.subject,
		body:      body,
	}
}

type toBuilder struct {
	emailType string
	subject   string
	body      string
}

// To sets the to address of the email
func (b *toBuilder) To(to ...string) *optionalBuilder {
	return &optionalBuilder{
		emailType: b.emailType,
		subject:   b.subject,
		body:      b.body,
		to:        to,
	}
}

type optionalBuilder struct {
	emailType   string
	subject     string
	body        string
	to          []string
	cc          []string
	bcc         []string
	attachments []Attachment
	values      map[string]string
}

// Cc sets the cc address of the email
func (b *optionalBuilder) Cc(cc ...string) *optionalBuilder {
	b.cc = cc
	return b
}

// Bcc sets the bcc address of the email
func (b *optionalBuilder) Bcc(bcc ...string) *optionalBuilder {
	b.bcc = bcc
	return b
}

// Attachments sets the attachments of the email
func (b *optionalBuilder) Attachments(attachments ...Attachment) *optionalBuilder {
	b.attachments = attachments
	return b
}

// Values sets the values of the email
// To replace the value in subject and body, set the key in like {{key}} format in subject and body
func (b *optionalBuilder) Values(values map[string]string) *optionalBuilder {
	b.values = values
	return b
}

// AddAttachment adds an attachment to the email
func (b *optionalBuilder) AddAttachment(attachment Attachment) *optionalBuilder {
	if b.attachments == nil {
		b.attachments = make([]Attachment, 0)
	}
	b.attachments = append(b.attachments, attachment)
	return b
}

// AddValue adds a value to the email
// if the key already exists, the value will be overwritten
// To replace the value in subject and body, set the key in like {{key}} format in subject and body
func (b *optionalBuilder) AddValue(key, value string) *optionalBuilder {
	if b.values == nil {
		b.values = make(map[string]string)
	}
	b.values[key] = value
	return b
}

// Build builds the email
func (b *optionalBuilder) Build() *Email {
	return &Email{
		Type:        b.emailType,
		To:          b.to,
		Cc:          b.cc,
		Bcc:         b.bcc,
		Subject:     b.subject,
		Body:        b.body,
		Attachments: b.attachments,
		Values:      b.values,
	}
}
