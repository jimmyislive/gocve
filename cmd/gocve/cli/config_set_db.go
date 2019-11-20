// +build linux,amd64 darwin,amd64

package cli

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dbType, dbHost, dbUser, dbName, password, tableName string
	dbPort                                              int
)

var configSetDBCmd = &cobra.Command{
	Use:   "set-db",
	Short: "Sets the url of the cve DB",
	Long:  `Sets the url of the cve DB`,
	Run: func(cmd *cobra.Command, args []string) {

		if dbType != "sqlite" && dbType != "postgres" {
			log.Fatal("DB type must be one of sqlite|postgres")
		}

		dir, _ := filepath.Split(config)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0700); err != nil {
				log.Fatal(err)
			}
		}

		v := viper.New()
		v.Set("dbType", dbType)
		v.Set("dbHost", dbHost)
		v.Set("dbPort", dbPort)
		v.Set("dbUser", dbUser)
		v.Set("dbName", dbName)
		v.Set("tableName", tableName)

		if err := v.WriteConfigAs(config); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Finished writing config at %s\n", config)
	},
}

func init() {
	configCmd.AddCommand(configSetDBCmd)

	configSetDBCmd.Flags().StringVar(&dbType, "dbType", "", "The type of the DB. Supported types are sqlite, postgres")
	configSetDBCmd.Flags().StringVar(&dbHost, "dbHost", "localhost", "The host on which your DB is running")
	configSetDBCmd.Flags().IntVar(&dbPort, "dbPort", 0, "The DB port")
	configSetDBCmd.Flags().StringVar(&dbUser, "dbUser", "", "The DB user to log in as")
	configSetDBCmd.Flags().StringVar(&dbName, "dbName", "cvedb", "The DB name")
	configSetDBCmd.Flags().StringVar(&tableName, "tableName", "cve", "The name of the table in the db")
	configSetDBCmd.Flags().StringVar(&password, "password", "", "The password for the db")

	configSetDBCmd.MarkFlagRequired("dbType")

	viper.BindPFlag("dbType", populateCmd.Flags().Lookup("dbType"))
	viper.BindPFlag("dbHost", populateCmd.Flags().Lookup("dbHost"))
	viper.BindPFlag("dbPort", populateCmd.Flags().Lookup("dbPort"))
	viper.BindPFlag("dbUser", populateCmd.Flags().Lookup("dbUser"))
	viper.BindPFlag("dbName", populateCmd.Flags().Lookup("dbName"))
	viper.BindPFlag("tableName", populateCmd.Flags().Lookup("tableName"))
	viper.BindPFlag("password", populateCmd.Flags().Lookup("password"))
}
