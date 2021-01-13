package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"godmr/ine"
	"godmr/radioid"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
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

	usersFile, err := os.Create("./user_es_20200113.csv")
	csvWriter := csv.NewWriter(usersFile)
	defer csvWriter.Flush()

	switch resp.StatusCode {
	case http.StatusOK:
		users, err := ReadJSON(resp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Downloaded %d users\n", len(users))

		var hits int

		// write header
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
					fmt.Printf("Callsign %s not found, was %s\n", user.Callsign, user.City)
				}
			}
			record = append(record, user.Country)
			record = append(record, user.Remarks)
			record = append(record, "0")
			record = append(record, "0")
			_ = csvWriter.Write(record)
		}

		// Get to fix top down so that we concentrate on the most common cases
		fmt.Printf("Hits were: %d\n", hits)
		for k, v := range misses {
			if v > 10 {
				fmt.Printf("%s -> %d\n", k, v)
			}
		}
	default:
		panic(errors.New("call to radioid.net did not work"))
	}
}
