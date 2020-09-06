package commands

import (
	"io/ioutil"

	log "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Create sample configuration in current working directory. File name: ./" + cfg.AppConfig.GetString("configFileName") + ".",
	Run:   configure,
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

func configure(cmd *cobra.Command, args []string) {
	sampleConfig := "---\n#List of roles that can be assumed by shortname.\nroles:\n" +
		"  name0: arn:aws:iam::123456789:role/name0\n" +
		"  name1: arn:aws:iam::123456789:role/name1\n" +
		"  name2: arn:aws:iam::987654321:role/name2\n" +
		"\n#Limit API requests to one region. Often used for compliance reasons in gov projects.\nregion: eu-west-2\n" +
		"\n#Log levels available in the application\n# 1: Info\n# 0: Debug\n# 2: Warning (show only critical problems)\nlogLevel: 1\n" +
		"mfaDeviceSerial: arn:aws:iam::<ACCOUNT_ID>:mfa/<IAM_USER_NAME>\n"
	err := ioutil.WriteFile(cfg.AppConfig.GetString("configFileName"), []byte(sampleConfig), 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write configuration file.")
	}
	log.Info().Msg("Configuration file has been generated.")
}
