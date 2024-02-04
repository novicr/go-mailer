package emailer

import (
	"bytes"
	"html/template"
	"mailer/internal/pkg/csvparser"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Emailer struct {
	From        string
	Subject     string
	DataMap     map[string]string
	Template    *template.Template
	EmailConfig *EmailConfig
}

// EmailConfig holds configuration for the SMTP server
type EmailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

// NewEmailConfig creates a new EmailConfig
func NewEmailer(from, subject, host, port, username, password, templatePath string, data map[string]string) (*Emailer, error) {

	// Parse the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return &Emailer{}, err
	}

	return &Emailer{
		From:     from,
		Subject:  subject,
		Template: tmpl,
		DataMap:  data,
		EmailConfig: &EmailConfig{
			Host:     host,
			Port:     port,
			Username: username,
			Password: password,
		},
	}, nil
}

// SendEmail sends an email using the provided template, recipient, and parameters
func (m *Emailer) SendEmail(recipient csvparser.Recipient) error {
	e := email.NewEmail()
	config := m.EmailConfig

	// Set up email basic info
	e.From = m.From
	e.To = []string{recipient.To}
	e.Subject = m.Subject

	data := map[string]any{
		"global": m.DataMap,
		"params": recipient.Params,
	}

	//fmt.Printf("to: %v, data: %v\n", e.To, data)

	var buf bytes.Buffer
	m.Template.ExecuteTemplate(&buf, "body", data)
	e.HTML = buf.Bytes()

	//_, err := fmt.Printf("Output: \n%s\n", buf.String())
	//return err

	// Address of the SMTP server
	addr := config.Host + ":" + config.Port

	// Send the email
	return e.Send(addr, smtp.PlainAuth("", config.Username, config.Password, config.Host))
}
