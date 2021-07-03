package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var configFileName, configFileType = "AWShelper.yaml", "yaml"

// Configuration structure that holds entire app configuration.
type Configuration struct {
	AppConfig  *viper.Viper
	AwsSession *session.Session
	MfaSerial  []*iam.MFADevice
}

// Conf global configuration
var Conf Configuration

// Load configuration from files and environment variables.
func Load() Configuration {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	Conf.AppConfig = configFiles()
	zerolog.SetGlobalLevel(zerolog.Level(Conf.AppConfig.GetInt("logLevel")))
	Conf.AwsSession, Conf.MfaSerial = awsSession()

	return Conf
}

func configFiles() *viper.Viper {
	viper.AddConfigPath("/etc/.AWShelper/")
	viper.AddConfigPath("$HOME/.AWShelper")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(configFileName)
	viper.AutomaticEnv()
	viper.Set("configFileName", configFileName)
	viper.Set("configFileType", configFileType)
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Err(err).Msg("Run 'AWShelper configure' subcommand")
		} else {
			log.Fatal().Err(err).Msg("Failed to load config file. Please run $ AWShelper configure")
		}
	}

	return viper.GetViper()
}

func awsSession() (*session.Session, []*iam.MFADevice) {

	awsConfig := aws.NewConfig()
	r := Conf.AppConfig.GetString("region")
	awsConfig.Region = &r
	awsConfig.STSRegionalEndpoint = 2
	b := true
	awsConfig.CredentialsChainVerboseErrors = &b

	awsOptions := &session.Options{
		Config:  *awsConfig,
		Profile: Conf.AppConfig.GetString("AWS_PROFILE"),
	}

	sess := session.Must(session.NewSessionWithOptions(*awsOptions))
	svc := iam.New(sess)
	output, err := svc.ListMFADevices(&iam.ListMFADevicesInput{})

	if err != nil {
		if e, ok := err.(awserr.RequestFailure); ok {
			log.Info().Err(e).Msg("Failed to fetch MFA devices. Please provide it via flag -e / --mfaDeviceSerial or configuration file. In order to disable this message in each run - set logLevel to '2' in configuration file.")
			var mfaDevicesList []*iam.MFADevice
			return sess, mfaDevicesList
		}
		log.Fatal().Err(err).Msg("Failied to list MFA devices. Make surre that you have set AWS_PROFILE variable or your 'default' profile works.")
	}

	return sess, output.MFADevices
}
