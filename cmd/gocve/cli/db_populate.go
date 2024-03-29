// +build linux,amd64 darwin,amd64

package cli

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	dbWrapper "github.com/jimmyislive/gocve/internal/pkg/db"
	ds "github.com/jimmyislive/gocve/internal/pkg/ds"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	fileName string
)

var populateCmd = &cobra.Command{
	Use:   "populate",
	Short: "Populate cve db with cve data",
	Long:  `Populate cve db with cve data`,
	Run: func(cmd *cobra.Command, args []string) {
		v := viper.New()
		cfg, err := initConfig(v)
		if err != nil {
			log.Fatal(err)
		}

		if fileName == "" {
			fmt.Println("--fileName is a required flag")
			fmt.Println()
			cmd.Help()
			os.Exit(1)
		}

		populateDB(cfg, fileName)
	},
}

func init() {
	populateCmd.Flags().StringVar(&fileName, "fileName", "", "File that contains the cve data to populate the db with")
	viper.BindPFlag("fileName", populateCmd.Flags().Lookup("fileName"))

	populateCmd.PersistentFlags().MarkHidden("configFile")

	dbCmd.AddCommand(populateCmd)
}

func populateDB(cfg *ds.Config, fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	// Read and skip the first 10 lines as they are headers
	fmt.Println("Skipping first 10 lines of header...")
	scanner := bufio.NewScanner(f)
	for i := 0; i < 10; i++ {
		scanner.Scan()
		fmt.Println("Skipped header: ", scanner.Text())
	}

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fileContent := strings.Join(lines, "\n")

	r := csv.NewReader(strings.NewReader(fileContent))
	var recordsList [][]string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		// Check if the field count is proper
		if err != nil {
			fmt.Println("Record found with incorrect number of fields:")
			fmt.Println(record)
			log.Fatal(err)
		}

		recordsList = append(recordsList, record)
	}

	err = dbWrapper.PopulateDB(cfg, recordsList)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
