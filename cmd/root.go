package cmd

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"godmr/ine"
	"godmr/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godmr",
	Short: "Gets digital contact lists and tries tidying them up",
	Long: `Gets digital contact lists from radioid.net, making it available to
filter by country or get the main, all-database dump`,
		Run: func(cmd *cobra.Command, args []string) {
			reader := bytes.NewReader(ine.Data)
			lines, err := csv.NewReader(reader).ReadAll()
			if err != nil {
				panic(err)
			}

			var misses = make(map[string]int)
			var dictionary = make(map[string]string)
			for _, line := range lines {
				dictionary[strings.ToLower(line[4])] = line[1]
			}

			//URL, _ := url.Parse("https://database.radioid.net/api/dmr/user/" +
			//	"?country=Spain" +
			//	"&country=Peru" +
			//	"&country=Ecuador" +
			//	"&country=Colombia" +
			//	"&country=Chile" +
			//	"&country=Bolivia" +
			//	"&country=Mexico" +
			//	"&country=Guatemala" +
			//	"&country=El%20Salvador" +
			//	"&country=Nicaragua" +
			//	"&country=Costa%20Rica" +
			//	"&country=Panama" +
			//	"&country=Honduras" +
			//	"&country=Cuba" +
			//	"&country=Dominican%20Republic" +
			//	"&country=Venezuela" +
			//	"&country=Paraguay" +
			//	"&country=Uruguay" +
			//	"&country=Argentina%20Republic")
			URL, _ := url.Parse("https://database.radioid.net/api/dmr/user/?country=Spain")

			req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
			if err != nil {
				panic("error creating request")
			}
			req.Header.Set("Accept-Encoding", "utf-8")
			req.Header.Set("Content-Type", "application/json")

			client := http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			now := time.Now()
			fileName := "users_" + now.Format("20060102150405") + ".csv"
			usersFile, err := os.Create(fileName)
			csvWriter := csv.NewWriter(usersFile)
			defer csvWriter.Flush()

			switch resp.StatusCode {
			case http.StatusOK:
				users, err := json.ReadJSON(resp)
				if err != nil {
					panic(err)
				}

				fmt.Printf("Downloaded %d users...\n", len(users))
				fmt.Printf("Tidying up...\n")

				var hits int
				_ = csvWriter.Write([]string{"Radio ID", "Callsign", "Name", "City", "State", "Country", "Remarks", "Call Type", "Call Alert"})
				for _, user := range users {
					var record []string
					record = append(record, strconv.FormatInt(user.RadioId, 10))
					record = append(record, user.Callsign)
					record = append(record, strings.Title(strings.ToLower(user.Name)))
					record = append(record, strings.Title(strings.ToLower(user.City)))

					key := strings.Trim(strings.ToLower(user.City), " ")
					if province, ok := dictionary[key]; ok {
						hits = hits + 1
						record = append(record, province)
					} else {
						record = append(record, user.State)
						// data is just from spain, so don't print anything from other countries. Also
						// 214 won't be -I guess- enough filter when things escalate
						if strings.HasPrefix(strconv.FormatInt(user.RadioId, 10), "214") {
							misses[key] = misses[key]+1
							if verbose {
								fmt.Printf("Callsign %s not found, was %s\n", user.Callsign, user.City)
							}
						}
					}
					record = append(record, user.Country)
					record = append(record, user.Remarks)
					record = append(record, "0")
					record = append(record, "0")
					_ = csvWriter.Write(record)
				}

				// Get to fix top down so that we concentrate on the most common cases
				if verbose {
					fmt.Printf("Hits were: %d\n", hits)
					for k, v := range misses {
						if v > 10 {
							fmt.Printf("%s -> %d\n", k, v)
						}
					}
				}

				fmt.Printf("Wrote file %s\nDone.\n", fileName)
			default:
				panic(errors.New("call to radioid.net did not work"))
			}
		},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print debugging messages")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".godmr" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".godmr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
