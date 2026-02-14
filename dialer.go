package smtp

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/MetaDiv-AI/smtp/internal/mime"

	"gopkg.in/mail.v2"
)

type Dialer struct {
	DisplayName string `json:"display_name"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

// Send sends an email
func (d *Dialer) Send(email *Email) error {
	if email == nil {
		return errors.New("email is nil")
	}

	dialer := mail.NewDialer(d.Host, d.Port, d.Username, d.Password)

	msg, err := d.buildMessage(email)
	if err != nil {
		return err
	}

	return dialer.DialAndSend(msg)
}

func (d *Dialer) buildMessage(m *Email) (*mail.Message, error) {
	msg := mail.NewMessage()
	msg.SetHeader("From", d.sender())
	msg.SetHeader("To", m.To...)
	msg.SetHeader("Cc", m.Cc...)
	msg.SetHeader("Bcc", m.Bcc...)

	for key, value := range m.Values {
		m.Body = strings.ReplaceAll(m.Body, "{{"+key+"}}", value)
		m.Subject = strings.ReplaceAll(m.Subject, "{{"+key+"}}", value)
		for i := range m.Attachments {
			m.Attachments[i].Filename = strings.ReplaceAll(m.Attachments[i].Filename, "{{"+key+"}}", value)
		}
	}
	msg.SetHeader("Subject", m.Subject)
	switch m.Type {
	case EmailTypeText:
		msg.SetBody("text/plain", m.Body)
	case EmailTypeHTML:
		msg.SetBody("text/html", m.Body)
	}

	for _, att := range m.Attachments {
		convertor := mime.NewConvertor()
		mimeType := convertor.FileNameToMime(att.Filename)
		if mimeType == "" {
			mimeType = "application/octet-stream" // fallback for unknown types
		}

		msg.Attach(att.Filename,
			mail.SetCopyFunc(func(w io.Writer) error {
				_, err := w.Write(att.Data)
				return err
			}),
			mail.SetHeader(map[string][]string{
				"Content-Type": {mimeType},
			}),
		)
	}
	return msg, nil
}

func (d *Dialer) sender() string {
	if d.DisplayName == "" {
		return d.Username
	}
	return fmt.Sprintf("%s <%s>", d.DisplayName, d.Username)
}
