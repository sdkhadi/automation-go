package webaction

import (
	"os"

	"github.com/magiconair/properties"
	"github.com/tebeka/selenium"
)

// Driver global variables
var Driver selenium.WebDriver

// DriverConnect : connect to driver
func DriverConnect() selenium.WebDriver {
	p := properties.MustLoadFile("${GOPATH}/src/github.com/automation-go/capabilities-web.properties", properties.UTF8)
	caps := selenium.Capabilities{
		"browserName":    p.MustGetString("browserName"),
		"browserVersion": p.MustGetString("browserVersion"),
		"enableVNC":      p.MustGetBool("enableVNC"),
		"enableVideo":    p.MustGetBool("enableVideo"),
		"videoName":      p.MustGetString("videoName"),
		"name":           p.MustGetString("name"),
	}
	Driver, _ = selenium.NewRemote(caps, os.Getenv("SELENIUM_URL"))

	return Driver
}
