package commands

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/dmatusiewicz/AWShelper/internal/config"
	"github.com/spf13/cobra"
)

// var cfg config.Configuration
var cfg = config.Load()
var rootCmd = &cobra.Command{
	Use:   "AWShelper command",
	Short: "AWShelper is an AWS API authentication wrapper.",
	Long:  "Wrapper that will export AWS credentials fetched from AWS after MFA authentication or role assumption.",
	Run:   root,
}

// SessionCredentials holds for each action
var SessionCredentials *sts.Credentials

// Run start of the program
func Run() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug level of logging.")
	rootCmd.PersistentFlags().StringP("mfa", "m", "", "MFA code.")
	rootCmd.PersistentFlags().StringP("mfaDeviceSerial", "e", "", "MFA device serial number.")
	cfg.AppConfig.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	cfg.AppConfig.BindPFlag("mfa", rootCmd.PersistentFlags().Lookup("mfa"))
	cfg.AppConfig.BindPFlag("mfaDeviceSerial", rootCmd.PersistentFlags().Lookup("mfaDeviceSerial"))

	rootCmd.Execute()
	printOutput()
}

func root(cmd *cobra.Command, args []string) {
	fmt.Print(cmd.UsageString())
}

func printOutput() {
	if SessionCredentials != nil {
		fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", *SessionCredentials.AccessKeyId)
		fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", *SessionCredentials.SecretAccessKey)
		fmt.Printf("export AWS_SESSION_TOKEN=%s\n", *SessionCredentials.SessionToken)
	}
}
