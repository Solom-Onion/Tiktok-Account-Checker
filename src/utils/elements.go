package utils

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func checkAndRemoveElement(wd selenium.WebDriver, xpath string) error {
	exists, err := wd.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		return err
	}

	if exists != nil {
		_, err = wd.ExecuteScript(fmt.Sprintf("document.evaluate('%s', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue.remove()", xpath), nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetElementText(wd selenium.WebDriver, xpath string) (string, error) {
	el, err := wd.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		return "", err
	}
	return el.Text()
}
