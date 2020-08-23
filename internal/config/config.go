package config

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var stsFS, roleFS *flag.FlagSet

// Load Parsing flags and loading config files.
func Load() {

	createFlagSets()
	// importFlags()
	// loadConfigFiles()

	fmt.Printf("%s\n", viper.Get("role"))
	fmt.Printf("%s\n", viper.Get("list"))
	fmt.Printf("%s\n", viper.Get("sts"))

}

func createFlagSets() {

	switch os.Args[1] {
	case "sts":
		{
			stsFS = flag.NewFlagSet("sts", 1)
			stsFS.Parse(os.Args[2:])
			viper.BindPFlags(stsFS)
		}
	case "role":
		{
			roleFS = flag.NewFlagSet("role", 1)
			roleFS.StringP("role", "r", "", "Role to be assumed.")
			roleFS.BoolP("list", "l", false, "Attempt to list roles.")
			roleFS.Parse(os.Args[2:])
			viper.BindPFlags(roleFS)
		}
	}
}

func loadConfigFiles() error {
	viper.AddConfigPath("/etc/.AWShelper/")
	viper.AddConfigPath("$HOME/.AWShelper")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
