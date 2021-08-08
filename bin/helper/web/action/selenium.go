/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-07 21:31:54
 * @modify date 2021-08-07 21:31:54
 * @desc [description]
 */
package webaction

import (
	"github.com/automation-go/bin/helper/common"
	"github.com/tebeka/selenium"
)

func (s Page) SendKeys(locator string, text string) error {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	common.LogPanicln(err)

	return element.SendKeys(text)
}
