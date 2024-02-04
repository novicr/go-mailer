package app

import (
	"log"
	"mailer/config"
	"mailer/internal/pkg/csvparser"
	"mailer/internal/pkg/emailer"
)

func SendEmails(from, subject, templatePath, recipientsPath, configPath string, globalParams map[string]string) {
	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	m, err := emailer.NewEmailer(from, subject, cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUsername, cfg.SMTPPassword, templatePath, globalParams)
	if err != nil {
		log.Fatalf("failed to create emailer: %v", err)
	}

	recipients, err := csvparser.ParseCSV(recipientsPath)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	for _, r := range recipients {
		err = m.SendEmail(r)
		if err != nil {
			log.Fatalf("failed to send email: %v", err)
		}
	}
}
