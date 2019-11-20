package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	ds "github.com/jimmyislive/gocve/internal/pkg/ds"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config, configFile, configDir string

var rootCmd = &cobra.Command{
	Use:   "gocve",
	Short: "Gocve is cli tool to get CVE details",
	Long:  `Gocve is cli tool and rest api server to view CVE details`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GoCVE ! (See help for usage)")
	},
}

func addCommonFlags(cmd *cobra.Command, configFile string, configDir string) {
	cmd.Flags().StringVar(&configFile, "configFile", "", "gocve.yaml")
	cmd.Flags().StringVar(&configDir, "configDir", "", "~/.gocve")
	viper.BindPFlag("configFile", cmd.Flags().Lookup("configFile"))
	viper.BindPFlag("configDir", cmd.Flags().Lookup("configDir"))
}

func setCfgDetails(v *viper.Viper) {
	v.SetConfigFile(config)
	v.AutomaticEnv()
}

func initConfig(v *viper.Viper) (*ds.Config, error) {
	v.SetConfigFile(config)
	v.SetEnvPrefix("GOCVE")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", v.ConfigFileUsed())
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not set up as yet. Use set-db")
		}
	}

	c := ds.Config{
		DBtype:    v.GetString("dbtype"),
		DBhost:    v.GetString("dbhost"),
		DBname:    v.GetString("dbname"),
		DBport:    v.GetInt("dbport"),
		DBuser:    v.GetString("dbuser"),
		Tablename: v.GetString("tablename"),
		Password:  v.GetString("password"),
	}

	return &c, nil
}

func readInConfig(v *viper.Viper) error {
	setCfgDetails(v)
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("Config file not set up as yet. Use set-db")
		}
		return err
	}

	return nil
}

func init() {

	cfgFile := filepath.Join(os.Getenv("HOME"), ".gocve", "gocve.yaml")

	rootCmd.PersistentFlags().StringVar(&config, "config", cfgFile, "Defaults to "+cfgFile)
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(getCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
