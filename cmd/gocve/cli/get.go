// +build linux,amd64 darwin,amd64

package cli

import (
	"fmt"
	"log"

	dbWrapper "github.com/jimmyislive/gocve/internal/pkg/db"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details about a particular CVE",
	Long:  `Get details about a particular CVE`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		v := viper.New()
		cfg, err := initConfig(v)
		if err != nil {
			log.Fatal(err)
		}

		record := dbWrapper.GetCVE(cfg, args[0])

		if len(record) == 0 {
			fmt.Println("No CVE with this id")
		} else {
			fmt.Println(record[0])
			for j := 0; j < len(record[0]); j++ {
				fmt.Print("=")
			}
			fmt.Println()
			fmt.Print("Status: ")
			fmt.Println(record[1])

			fmt.Println()
			fmt.Print("Description: ")
			fmt.Println(record[2])

			fmt.Println()
			fmt.Print("Reference: ")
			fmt.Println(record[3])

			fmt.Println()
			fmt.Print("Phase: ")
			fmt.Println(record[4])

			fmt.Println()
			fmt.Print("Category: ")
			fmt.Println(record[5])
			fmt.Println()
		}
	},
}
