// +build linux,amd64 darwin,amd64

package cli

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	url, path string
	force     bool
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download cve data",
	Long:  `Downloads cve data from the mitre website`,
	Run: func(cmd *cobra.Command, args []string) {
		cveDBURL := viper.GetString("url")
		path := viper.GetString("path")

		err := downloadDB(cveDBURL, path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("CVE DB successfully downloaded.")
	},
}

func init() {

	downloadCmd.Flags().StringVar(&url, "url", "https://cve.mitre.org/data/downloads/allitems.csv.gz", "cve db url (default is https://cve.mitre.org/data/downloads/allitems.csv.gz)")
	downloadCmd.Flags().BoolVar(&force, "force", false, "force overwrite, if file exists")
	downloadCmd.Flags().StringVar(&path, "path", ".", "cve db path (default is current directory)")
	viper.BindPFlag("url", downloadCmd.Flags().Lookup("url"))
	viper.BindPFlag("force", downloadCmd.Flags().Lookup("force"))
	viper.BindPFlag("path", downloadCmd.Flags().Lookup("path"))

	dbCmd.AddCommand(downloadCmd)
}

func downloadDB(cveDBURL, path string) error {

	// Get the filename from the URL
	fileNameSuffix := strings.Split(cveDBURL, "/")
	fileName := fileNameSuffix[len(fileNameSuffix)-1]
	fileNameWithPath := filepath.Join(path, fileName)

	// Check if it exists first, only if --force is false
	if !viper.GetBool("force") {
		if _, err := os.Stat(fileNameWithPath); err != nil {
			// If its a system error, return it
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			return fmt.Errorf("File %v already exists. Please remove and try again. (or use the --force flag)", fileNameWithPath)
		}
	}

	var resp *http.Response
	var err error

	fmt.Printf("Downloading cve db from %s\n", cveDBURL)

	if strings.HasSuffix(cveDBURL, ".gz") {
		tr := &http.Transport{
			DisableCompression: true,
		}
		client := &http.Client{Transport: tr}

		resp, err = client.Get(cveDBURL)
		if err != nil {
			return err
		}
	} else {
		resp, err = http.Get(cveDBURL)
		if err != nil {
			return err
		}
	}

	defer resp.Body.Close()

	out, err := os.Create(fileNameWithPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
