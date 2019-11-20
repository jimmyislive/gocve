// +build linux,amd64 darwin,amd64

package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "DB Commands",
	Long:  `Commands related to setting up the CVE DB`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	origHelpFunc := dbCmd.HelpFunc()
	dbCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.Flags().MarkHidden("configFile")
		cmd.Flags().MarkHidden("configDir")
		origHelpFunc(cmd, args)
	})
}
