// +build linux,amd64 darwin,amd64

package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the gocve config settings",
	Long:  `Show the gocve config settings`,
	Run: func(cmd *cobra.Command, args []string) {

		v := viper.New()
		cfg, err := initConfig(v)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("dbtype: ", cfg.DBtype)
		fmt.Println("dbhost: ", cfg.DBhost)
		fmt.Println("dbname: ", cfg.DBname)
		fmt.Println("dbport: ", cfg.DBport)
		fmt.Println("dbuser: ", cfg.DBuser)
		fmt.Println("tablename: ", cfg.Tablename)
		if cfg.DBtype == "postgres" {
			fmt.Println("password: ", cfg.Password)
		}

	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
}
