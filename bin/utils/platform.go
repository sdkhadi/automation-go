/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-09 00:05:52
 * @modify date 2021-08-09 00:05:52
 * @desc [description]
 */
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/automation-go/bin/helper/common"
	webaction "github.com/automation-go/bin/helper/web/action"
	config "github.com/automation-go/bin/helper/web/config"
	"github.com/automation-go/bin/utils/models"
	"github.com/cucumber/godog"
)

var Web webaction.Page

var PWD string
var testCase models.ScenarioDetail
var path, pathname, filename string
var scenarioTags, featureTags, tags string

func WebConnect() error {
	config.DriverConnect()
	Web = webaction.Page{Action: config.Driver}

	return nil
}

func structDetail(scenario interface{}) error {
	data, _ := json.Marshal(scenario)

	return json.Unmarshal(data, &testCase)
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.BeforeScenario(func(scenario *godog.Scenario) {
		structDetail(scenario)

		i := 0

		for i < len(testCase.Tags) {
			scenarioTags += " " + testCase.Tags[i].Name

			i++
		}

		WebConnect()
	})

	s.AfterScenario(func(scenario *godog.Scenario, log error) {
		var err error

		structDetail(scenario)

		scenarioTags = ""
		PWD, err = os.Getwd()
		common.LogPanicln(err)

		if log != nil {
			filename = fmt.Sprintf("Screenshot - FAILED - %s - %s - %s.png", testCase.Name, testCase.Steps[0].Text, log)
			ssWeb()
		}
	})

}
func ssWeb() error {
	if config.Driver != nil {
		path = fmt.Sprintf("%s/screenshots/web", PWD)

		takeErrorWebPageImage()
	}

	return nil
}
func takeErrorWebPageImage() error {
	buffWeb := webaction.Page{Action: config.Driver}

	dirCheck()
	buffWeb.TakeScreenshot()
	ioutil.WriteFile(pathname, []byte(fmt.Sprintf("%v", buffWeb)), 0644)

	return nil
}
func dirCheck() error {
	pathname = filepath.Join(path, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}

	return nil
}
