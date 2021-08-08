/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-09 00:06:29
 * @modify date 2021-08-09 00:06:29
 * @desc [description]
 */
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tebeka/selenium"
)

func main() {
	errror := godotenv.Load()
	if errror != nil {
		log.Fatal("Error loading .env file")
	}
	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome", "browserVersion": "87.0", "enableVNC": true, "enableVideo": true, "name": "Golang Automation", "videoName": "golang-test.mp4"})
	if webDriver, err = selenium.NewRemote(caps, os.Getenv("SELENIUM_URL")); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	err = webDriver.Get("https://katalon-demo-cura.herokuapp.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	if title, err := webDriver.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		fmt.Printf("Failed to get page title: %s", err)
		return
	}

	err = webDriver.MaximizeWindow("")
	if err != nil {
		if err != nil {
			fmt.Printf("Failed to maximize window: %s\n", err)
			return
		}
	}
	err = webDriver.SetImplicitWaitTimeout(30)
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}
	var elem selenium.WebElement

	elem, err = webDriver.FindElement(selenium.ByID, "btn-make-appointment")
	elem.Click()
	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return
	}

	elem, err = webDriver.FindElement(selenium.ByID, "txt-username")
	elem.SendKeys("John Doe")
	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return
	}
	elem, err = webDriver.FindElement(selenium.ByID, "txt-password")
	elem.SendKeys("ThisIsNotAPassword")
	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return
	}
	elem, err = webDriver.FindElement(selenium.ByID, "btn-login")
	elem.Click()
	if err != nil {
		fmt.Printf("Failed to click button: %s\n", err)
		return
	}
}
