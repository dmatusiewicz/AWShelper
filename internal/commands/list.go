package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"lst"},
	Short:   "List avalible roles from configuration",
	Run:     list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {

	if cfg.AppConfig.Get("roles") != "" {
		for i, o := range cfg.AppConfig.GetStringMap("roles") {
			fmt.Printf("%s %s\n", i, o)
		}
	}
}
