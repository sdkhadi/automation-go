/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-07 21:31:54
 * @modify date 2021-08-07 21:31:54
 * @desc [description]
 */
package webaction

import (
	"strconv"

	"github.com/automation-go/bin/helper/common"
)

func (s Page) SendKeys(findElementBy string, locator string, text string) error {
	element, err := s.driver().FindElement(findElementBy, locator)
	common.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) Click(findElementBy string, locator string) error {
	element, err := s.driver().FindElement(findElementBy, locator)
	common.LogPanicln(err)

	return element.Click()
}

func (s Page) Numpad(numPadKey string, locator string) int {
	element, err := strconv.Atoi(numPadKey)
	common.LogPanicln(err)

	s.driver().Click(element)

	return element
}

func (s Page) Clear(findElementBy string, locator string) error {
	element, err := s.driver().FindElement(findElementBy, locator)
	common.LogPanicln(err)

	return element.Clear()
}

func (s Page) GetText(findElementBy string, locator string) (string, error) {
	element, err := s.driver().FindElement(findElementBy, locator)
	common.LogPanicln(err)

	return element.Text()
}
