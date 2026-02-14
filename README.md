# smtp

A Go package for sending emails via SMTP with a fluent builder API. Built on top of [gopkg.in/mail.v2](https://github.com/go-mail/mail).

## Installation

```bash
go get github.com/MetaDiv-AI/smtp
```

## Features

- **Fluent builder API** for both dialer and email configuration
- **Text and HTML** email support
- **Attachments** with automatic MIME type detection
- **Template values** — replace `{{key}}` placeholders in subject, body, and attachment filenames
- **CC and BCC** support

## Quick Start

### Sending a Simple Email

```go
package main

import (
	"log"

	"github.com/MetaDiv-AI/smtp"
)

func main() {
	dialer := smtp.NewDialer("My App").
		Host("smtp.example.com").
		Port(587).
		Username("user@example.com").
		Password("your-password")

	email := smtp.NewEmail(smtp.EmailTypeHTML).
		Subject("Hello!").
		Body("<h1>Welcome</h1><p>Thanks for signing up.</p>").
		To("recipient@example.com").
		Build()

	if err := dialer.Send(email); err != nil {
		log.Fatal(err)
	}
}
```

### Using Template Values

Replace placeholders in subject, body, and attachment filenames with `{{key}}`:

```go
email := smtp.NewEmail(smtp.EmailTypeText).
	Subject("Order {{order_id}} confirmed").
	Body("Hi {{name}}, your order {{order_id}} has been confirmed.").
	To("customer@example.com").
	Values(map[string]string{
		"name":     "Alice",
		"order_id": "12345",
	}).
	Build()
```

### Sending with Attachments

```go
email := smtp.NewEmail(smtp.EmailTypeHTML).
	Subject("Your invoice").
	Body("<p>Please find your invoice attached.</p>").
	To("customer@example.com").
	Attachments(smtp.Attachment{
		Filename: "invoice.pdf",
		Data:     pdfBytes,
	}).
	Build()
```

### Full Example with Optional Fields

```go
email := smtp.NewEmail(smtp.EmailTypeHTML).
	Subject("Project Update").
	Body("<h1>Weekly Report</h1><p>See attached.</p>").
	To("team@example.com").
	Cc("manager@example.com").
	Bcc("archive@example.com").
	AddAttachment(smtp.Attachment{Filename: "report.pdf", Data: reportData}).
	AddValue("project", "Alpha").
	Build()
```

## API Reference

### Email Types

- `smtp.EmailTypeText` — plain text email
- `smtp.EmailTypeHTML` — HTML email

### Dialer Builder

| Method | Description |
|--------|-------------|
| `NewDialer(displayName)` | Start building a dialer |
| `Host(host)` | SMTP server host |
| `Port(port)` | SMTP server port (e.g., 587 for TLS) |
| `Username(username)` | SMTP username |
| `Password(password)` | SMTP password |

### Email Builder

| Method | Description |
|--------|-------------|
| `NewEmail(emailType)` | Start building an email |
| `Subject(subject)` | Email subject |
| `Body(body)` | Email body |
| `To(to...)` | Recipient addresses |
| `Cc(cc...)` | CC addresses |
| `Bcc(bcc...)` | BCC addresses |
| `Attachments(attachments...)` | Attachments |
| `AddAttachment(attachment)` | Add a single attachment |
| `Values(values)` | Template replacement map |
| `AddValue(key, value)` | Add a single template value |
| `Build()` | Build the email |

## Requirements

- Go 1.23.2 or later
