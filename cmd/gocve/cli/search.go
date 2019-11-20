// +build linux,amd64 darwin,amd64

package cli

import (
	"fmt"
	"log"

	dbWrapper "github.com/jimmyislive/gocve/internal/pkg/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cveid, desc string

var searchCmd = &cobra.Command{
	Use:   "search [search text]",
	Short: "Searches the CVE DB for a pattern",
	Long:  `Searches the CVE DB for a pattern`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		v := viper.New()
		cfg, err := initConfig(v)
		if err != nil {
			log.Fatal(err)
		}

		records := dbWrapper.SearchCVE(cfg, args[0])

		for i := 0; i < len(records); i++ {
			fmt.Println(records[i][0])
			for j := 0; j < len(records[i][0]); j++ {
				fmt.Print("=")
			}
			fmt.Println()
			fmt.Println(records[i][1])
			fmt.Println()
		}

	},
}
