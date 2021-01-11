package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"godmr/radioid"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type Results struct {
	Count   int               `json:"count"`
	Results []radioid.Contact `json:"results"`
}

func ReadJSON(response *http.Response) ([]radioid.Contact, error) {
	var results Results
	buffer := new(bytes.Buffer)
	_, _ = buffer.ReadFrom(response.Body)

	respByte := buffer.Bytes()
	if err := json.Unmarshal(respByte, &results); err != nil {
		return nil, err
	}

	return results.Results, nil
}

func main() {
	URL, _ := url.Parse("https://database.radioid.net/api/dmr/user/?country=Spain")

	client := http.Client{}
	resp, err := client.Get(URL.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	usersFile, err := os.Create("./users.csv")
	csvWriter := csv.NewWriter(usersFile)

	switch resp.StatusCode {
	case http.StatusOK:
		users, err := ReadJSON(resp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Downloaded %d users", len(users))

		// write header
		_ = csvWriter.Write([]string{"RADIO_ID", "CALLSIGN", "FIRST_NAME", "LAST_NAME", "CITY", "STATE", "COUNTRY", "REMARKS"})
		for _, user := range users {
			var record []string
			record = append(record, strconv.FormatInt(user.RadioId, 10))
			record = append(record, user.Callsign)
			record = append(record, user.Name)
			record = append(record, user.Surname)
			record = append(record, user.City)
			record = append(record, user.Province)
			record = append(record, user.Country)
			record = append(record, "DMR")
			_ = csvWriter.Write(record)
		}
		csvWriter.Flush()
	default:
		panic(errors.New("call to radioid.net not worked"))
	}
}