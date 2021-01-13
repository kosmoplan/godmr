package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"godmr/esgob"
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
	reader := bytes.NewReader(esgob.GetData())
	lines, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		panic(err)
	}

	var dictionary = make(map[string]string)

	// Loop through lines & turn into object
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

	//URL, _ := url.Parse("https://database.radioid.net/api/dmr/user/?country=%")

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

			if province, ok := dictionary[strings.Trim(strings.ToLower(user.City), " ")]; ok {
				hits = hits + 1
				//fmt.Printf("Got %s, inserted %s\n", user.City, province)
				record = append(record, province)
			} else {
				record = append(record, user.State)
				if strings.HasPrefix(strconv.FormatInt(user.RadioId, 10), "214") {
					fmt.Printf("Callsign %s not found, was %s\n", user.Callsign, user.City)
				}
			}
			record = append(record, user.Country)
			record = append(record, user.Remarks)
			record = append(record, "0")
			record = append(record, "0")
			_ = csvWriter.Write(record)
		}

		fmt.Printf("Hits was: %d\n", hits)
	default:
		panic(errors.New("call to radioid.net not worked"))
	}
}
