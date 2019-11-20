// +build linux,amd64 darwin,amd64

package cli

import (
	"fmt"
	"log"

	dbWrapper "github.com/jimmyislive/gocve/internal/pkg/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all cve ids",
	Long:  `Lists all cve ids`,
	Run: func(cmd *cobra.Command, args []string) {
		v := viper.New()
		cfg, err := initConfig(v)
		if err != nil {
			log.Fatal(err)
		}

		records := dbWrapper.ListCVE(cfg)

		for i := 0; i < len(records); i++ {
			max := len(records[i][1])
			if max > 100 {
				max = 100
			}
			fmt.Println(records[i][0], "\t", records[i][1][:max])
		}
	},
}
