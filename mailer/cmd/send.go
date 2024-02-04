package cmd

import (
	"log"
	"mailer/internal/app"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	var from, subject, templatePath, recipientsPath, configPath string
	var globalParams []string

	var sendCmd = &cobra.Command{
		Use:   "send",
		Short: "Send emails from a template to a list of recipients",
		Run: func(cmd *cobra.Command, args []string) {
			app.SendEmails(from, subject, templatePath, recipientsPath, configPath, parseParams(globalParams))
		},
	}

	sendCmd.Flags().StringVarP(&from, "from", "f", "", "'From' email address")
	sendCmd.Flags().StringVarP(&subject, "subject", "s", "", "Email Subject")
	sendCmd.Flags().StringVarP(&templatePath, "template", "t", "", "Path to the email template")
	sendCmd.Flags().StringVarP(&recipientsPath, "recipients", "r", "", "Path to the recipients CSV file")
	sendCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to the application configuration file")
	sendCmd.Flags().StringSliceVarP(&globalParams, "param", "p", []string{}, "Global parameters in the format key=value")

	rootCmd.AddCommand(sendCmd)
}

// parseParams converts a slice of strings in the format key=value into a map
func parseParams(paramsSlice []string) map[string]string {
	paramsMap := make(map[string]string)
	for _, param := range paramsSlice {
		parts := strings.SplitN(param, "=", 2)
		if len(parts) != 2 {
			log.Fatalf("invalid parameter format: %s", param)
		}
		paramsMap[parts[0]] = parts[1]
	}
	return paramsMap
}
