// +build linux,amd64 darwin,amd64

package cli

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var buildVersion string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of gocve",
	Long:  `Prints the version of gocve`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("GoCVE Version: %s %s/%s\n", buildVersion, runtime.GOOS, runtime.GOARCH)
		return nil
	},
}
