package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

func GetCookies() ([][]*selenium.Cookie, error) {
	// Read cookies from file
	cookiesFile, err := os.Open("./data/cookies.json")
	if err != nil {
		return nil, err
	}
	defer cookiesFile.Close()

	var cookies [][]*selenium.Cookie
	err = json.NewDecoder(cookiesFile).Decode(&cookies)
	if err != nil {
		return nil, err
	}

	for _, cookieBatch := range cookies {
		fmt.Println("Cookie Batch: ", cookieBatch)
	}

	return cookies, nil
}
