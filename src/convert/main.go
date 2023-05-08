package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cookie struct {
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Path    string `json:"path"`
	Expires int64  `json:"expires"`
}

func main() {
	// open cookies.txt file for reading
	file, err := os.Open("cookies.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cookies []*Cookie

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		if len(fields) < 6 {
			continue
		}

		domain := fields[0]
		name := fields[5]
		value := fields[6]
		path := fields[2]
		expirationStr := fields[4]

		expiry, err := strconv.ParseInt(expirationStr, 10, 64)

		if err != nil {
			// ignore invalid expiry date
			continue
		}

		// create new cookie
		cookie := &Cookie{
			Domain:  domain,
			Name:    name,
			Value:   value,
			Path:    path,
			Expires: expiry,
		}

		// append to cookies slice
		cookies = append(cookies, cookie)
	}

	// write cookies to cookies.json file
	cookiesFile, err := os.Create("cookies.json")
	if err != nil {
		panic(err)
	}
	defer cookiesFile.Close()

	encoder := json.NewEncoder(cookiesFile)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(cookies)
	if err != nil {
		panic(err)
	}

	fmt.Println("Cookies written to cookies.json")
}
