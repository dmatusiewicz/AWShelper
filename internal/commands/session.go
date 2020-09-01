package commands

import (
	log "github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:     "session",
	Aliases: []string{"ses"},
	Short:   "Gets session credentials from AWS API after MFA authentication.",
	Long:    "Wrapper that will export AWS credentials fetched from AWS after MFA authentication.",
	Run:     session,
}

func init() {
	sessionCmd.PersistentFlags().StringP("role", "r", "", "IAM Role to be assumed.")
	cfg.AppConfig.BindPFlag("role", sessionCmd.PersistentFlags().Lookup("role"))

	rootCmd.AddCommand(sessionCmd)
}

func session(cmd *cobra.Command, args []string) {

	svc := sts.New(cfg.AwsSession)
	mfaCode := cfg.AppConfig.GetString("mfa")
	if mfaCode == "" {
		log.Info().Msg("--mfa / -m is not set. Get session operation might fail.")
	}
	if 
	gsto, err := svc.GetSessionToken(&sts.GetSessionTokenInput{
		SerialNumber: cfg.MfaSerial[0].SerialNumber,
		TokenCode:    &mfaCode,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failied to get new session token.")
	}
	SessionCredentials = gsto.Credentials
}
