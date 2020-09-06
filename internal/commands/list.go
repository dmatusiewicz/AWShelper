package commands

import (
	"fmt"
	"sort"

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
		roleMap := cfg.AppConfig.GetStringMap("roles")
		var sortedRoles []string
		for k := range roleMap {
			sortedRoles = append(sortedRoles, k)
		}
		sort.Strings(sortedRoles)

		for _, v := range sortedRoles {

			fmt.Printf("\t- [ %s ]\t%s\n", v, roleMap[v])
		}
	}
}
