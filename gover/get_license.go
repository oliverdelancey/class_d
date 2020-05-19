package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handle_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write_file(name string, contents string) {
	cwd, err := os.Getwd()
	file_name := fmt.Sprintf("%s/%s", cwd, name)
	f, err := os.Create(file_name)
	handle_error(err)
	defer f.Close()
	_, err = f.WriteString(contents)
	handle_error(err)
	f.Sync()
}

func read_lb() []byte {
	content, err := ioutil.ReadFile("license_table.txt")
	handle_error(err)
	return content
}

func ggr1() string {
	resp, err := http.Get("https://api.github.com/licenses")
	handle_error(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handle_error(err)
	return string(body)
}

func main() {
	type LicenseTableEntry struct {
		Key     string
		Name    string
		Spdx_id string
		Url     string
		Node_id string
	}
	raw_data := read_lb()
	var table []LicenseTableEntry
	err := json.Unmarshal(raw_data, &table)
	handle_error(err)

	var licenseIDs []string
	for _, entry := range table {
		licenseIDs = append(licenseIDs, entry.Key)
	}
	fmt.Println(licenseIDs)
}
