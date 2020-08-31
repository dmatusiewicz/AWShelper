package commands

import (
	"os"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var roleCmd = &cobra.Command{
	Use:     "role",
	Aliases: []string{"rl"},
	Short:   "Gets session credentials from AWS API after MFA authentication to role.",
	Long:    "Wrapper that will export AWS credentials fetched from AWS after MFA role authentication.",
	Run:     role,
}
var roleToBeAssumed string

func init() {
	roleCmd.PersistentFlags().StringP("role", "r", "", "IAM Role to be assumed.")
	cfg.AppConfig.BindPFlag("role", roleCmd.PersistentFlags().Lookup("role"))
	rootCmd.AddCommand(roleCmd)
}

func role(cmd *cobra.Command, args []string) {

	// fmt.Println(cfg.AppConfig.Get("roles"))
	// fmt.Println(cfg.AppConfig.GetString("role"))
	// fmt.Println(cfg.AppConfig.Get("roles." + cfg.AppConfig.GetString("role")))
	if cfg.AppConfig.GetString("role") == "" {
		log.Fatal().Msg("--role / -r  has not been set.")
	} else {
		if cfg.AppConfig.Get("roles."+cfg.AppConfig.GetString("role")) == nil {
			log.Info().Msg("Configuartion lookup of 'role." + cfg.AppConfig.GetString("role") + "' failed. Using role flag litterally.")
			roleToBeAssumed = cfg.AppConfig.GetString("role")
		} else {
			roleToBeAssumed = cfg.AppConfig.GetString("roles." + cfg.AppConfig.GetString("role"))
		}
	}

	user := os.Getenv("USER")
	svc := sts.New(cfg.AwsSession)
	mfaCode := cfg.AppConfig.GetString("mfa")
	if mfaCode == "" {
		log.Warn().Msg("--mfa / -m is not set. Get session operation might fail.")
	}
	gsto, err := svc.AssumeRole(&sts.AssumeRoleInput{
		SerialNumber:    cfg.MfaSerial[0].SerialNumber,
		TokenCode:       &mfaCode,
		RoleArn:         &roleToBeAssumed,
		RoleSessionName: &user,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failied to get new session token.")
	}
	SessionCredentials = gsto.Credentials
}
