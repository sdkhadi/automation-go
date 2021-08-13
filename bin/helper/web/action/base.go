/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-09 00:06:00
 * @modify date 2021-08-09 00:06:00
 * @desc [description]
 */
package webaction

import (
	"github.com/automation-go/bin/helper/common"
	"github.com/tebeka/selenium"
)

// AcceptAlert : accepting alert
func (s Page) AcceptAlert() error {
	return s.driver().AcceptAlert()
}

// DismissAlert : dismissing alert
func (s Page) DismissAlert() error {
	return s.driver().DismissAlert()
}

// BackBrowser : back to previous page
func (s Page) BackBrowser() error {
	return s.driver().Back()
}

// GetCurrentURL : get current browser url
func (s Page) GetCurrentURL() string {
	element, err := s.driver().CurrentURL()
	common.LogPanicln(err)

	return element
}

// ButtonDown : button down
func (s Page) ButtonDown() error {
	return s.driver().ButtonDown()
}

// ButtonUp : buttom up
func (s Page) ButtonUp() error {
	return s.driver().ButtonUp()
}

// CloseWindow : close browser window
func (s Page) CloseWindow(locator string) error {
	return s.driver().CloseWindow(locator)
}

// AddCookie : add cookies in browser
func (s Page) AddCookie(cookie *selenium.Cookie) error {
	return s.driver().AddCookie(cookie)
}

// DeleteCookie : delete cookies in browser
func (s Page) DeleteCookie(locator string) error {
	return s.driver().DeleteCookie(locator)
}

// DeleteAllCookies : delete all cookies in browser
func (s Page) DeleteAllCookies(locator string) error {
	return s.driver().DeleteAllCookies()
}

// DoubleClick : double click in browser
func (s Page) DoubleClick() error {
	return s.driver().DoubleClick()
}

// Refresh : refresh browser
func (s Page) Refresh() error {
	return s.driver().Refresh()
}

// ResizeWindow : set browser size
func (s Page) ResizeWindow(locator string, w int, h int) error {
	return s.driver().ResizeWindow(locator, w, h)
}

// MaxWindow : set browser to maximum size
func (s Page) MaxWindow(locator string) error {
	return s.driver().MaximizeWindow(locator)
}

// TakeScreenshot : take error screenshot
func (s Page) TakeScreenshot() []byte {
	element, err := s.driver().Screenshot()
	common.LogPanicln(err)

	return element
}
