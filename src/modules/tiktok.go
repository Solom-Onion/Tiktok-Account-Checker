package modules

import (
	"fmt"
	"log"
	"solomon/bin/utils"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var (
	followerXPath  = "/html/body/div[1]/div[3]/div[2]/div/div[1]/h3/div[2]/strong"
	followingXPath = "/html/body/div[1]/div[3]/div[2]/div/div[1]/h3/div[1]/strong"
	usernameXPath  = "/html/body/div[1]/div[3]/div[2]/div/div[1]/div[1]/div[2]/h2"
	userBioXPath   = "/html/body/div[1]/div[3]/div[2]/div/div[1]/div[1]/div[2]/h1"
	likesXPath     = "/html/body/div[1]/div[3]/div[2]/div/div[1]/h3/div[3]/strong"
	setNameXPath   = "/html/body/div[1]/div[3]/div[2]/div/div[1]/h2"
)

type Element struct {
	Name       string
	XPath      string
	IsRequired bool
}

func ExecuteTiktokChecker() {
	cookies, err := utils.GetCookies()
	if err != nil {
		// TODO: handle error
	}
	opts := []selenium.ServiceOption{}
	chromeCaps := chrome.Capabilities{
		Path: "",
	}
	chromeCaps.Args = append(chromeCaps.Args, "--disable-blink-features=AutomationControlled")
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeCaps)

	service, err := selenium.NewChromeDriverService("./chromedriver", 9515, opts...)
	if err != nil {
		fmt.Println("Error starting the ChromeDriver server: ", err)
		return
	}
	defer service.Stop()

	for _, cookieBatch := range cookies {
		// Create Chrome WebDriver
		caps := selenium.Capabilities{"browserName": "chrome"}
		caps.AddChrome(chromeCaps)

		wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
		if err != nil {
			log.Println("Error creating Chrome WebDriver:", err)
			return
		}
		defer wd.Quit()

		err = wd.Get("https://www.tiktok.com/")
		if err != nil {
			log.Println("Error navigating to URL:", err)
			return
		}

		// Inject cookies into the page
		for _, cookie := range cookieBatch {
			cookieJS := fmt.Sprintf("document.cookie='%s=%s;domain=%s;path=%s;expires=%s';", cookie.Name, cookie.Value, cookie.Domain, cookie.Path, "")
			_, err := wd.ExecuteScript(cookieJS, nil)
			if err != nil {
				// log.Println("Error injecting cookie:", err)
				continue
			}
		}

		// Refresh the page
		err = wd.Refresh()
		if err != nil {
			log.Println("Error refreshing page:", err)
			return
		}

		// Click the element with ID header-more-menu-icon
		moreMenu, err := wd.FindElement(selenium.ByID, "header-more-menu-icon")
		if err != nil {
			panic(err)
		}
		if err := moreMenu.Click(); err != nil {
			panic(err)
		}

		log.Println("Checking profile details...")

		// Click elm with profile info
		viewProfileSpan, err := wd.FindElement(selenium.ByXPATH, "//span[contains(text(), 'View profile')]")
		if err != nil {
			log.Fatal("Invalid account!")
		}
		err = viewProfileSpan.Click()
		if err != nil {
			log.Fatal(err)
		}

		// Wait for the page to load
		time.Sleep(10 * time.Second)
		log.Println("Valid account!")

		// Define the elements to retrieve
		elements := []Element{
			{"Follower count", followerXPath, true},
			{"Following count", followingXPath, true},
			{"Username", usernameXPath, true},
			{"User bio", userBioXPath, false},
			{"Likes count", likesXPath, true},
			{"Set name", setNameXPath, true},
		}

		// Retrieve the element values and log them
		for _, element := range elements {
			value, err := utils.GetElementText(wd, element.XPath)
			if err != nil && element.IsRequired {
				log.Printf("Error getting %s: %s", element.Name, err)
			} else {
				log.Printf("%s: %s", element.Name, value)
			}
		}
	}
}
