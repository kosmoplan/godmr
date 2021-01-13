package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"godmr/radioid"
	"io"
	"io/ioutil"
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

func DumpJSON(response *http.Response) error {
	usersFile, err := os.Create("./users.json")
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(usersFile)

	teeReader := io.TeeReader(response.Body, writer)
	_, err = ioutil.ReadAll(teeReader)

	return err
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
	//URL, _ := url.Parse("https://database.radioid.net/api/dmr/user/?country=Spain")

	URL, _ := url.Parse("https://database.radioid.net/api/dmr/user/?country=%")

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

	usersFile, err := os.Create("./user_completa_20200113.csv")
	csvWriter := csv.NewWriter(usersFile)
	defer csvWriter.Flush()

	switch resp.StatusCode {
	case http.StatusOK:
		users, err := ReadJSON(resp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Downloaded %d users\n", len(users))

		// write header
		_ = csvWriter.Write([]string{"Radio ID", "Callsign", "Name", "City", "State", "Country", "Remarks", "Call Type", "Call Alert"})
		for _, user := range users {
			var record []string
			record = append(record, strconv.FormatInt(user.RadioId, 10))
			record = append(record, user.Callsign)
			record = append(record, strings.Title(strings.ToLower(user.Name)))
			record = append(record, strings.Title(strings.ToLower(user.City)))
			record = append(record, user.State)
			record = append(record, user.Country)
			record = append(record, user.Remarks)
			record = append(record, "0")
			record = append(record, "0")
			_ = csvWriter.Write(record)
		}
	default:
		panic(errors.New("call to radioid.net not worked"))
	}
}
