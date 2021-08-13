/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-09 00:06:40
 * @modify date 2021-08-09 00:06:40
 * @desc [description]
 */
package webpages

import (
	webaction "github.com/automation-go/bin/helper/web/action"
)

var (
	editTextUsernameByID = "txt-username"
	editTextPasswordByID = "txt-password"
	btnLoginByID         = "btn-login"
)

type LoginPage struct {
	Page webaction.Page
}

func (s LoginPage) InputUsername() *LoginPage {
	s.Page.SendKeys(webaction.ByID, editTextUsernameByID, "test")

	return &LoginPage{Page: s.Page}
}

// InputPassword : fill in password
func (s LoginPage) InputPassword() *LoginPage {
	s.Page.SendKeys(webaction.ByID, editTextPasswordByID, "")

	return &LoginPage{Page: s.Page}
}

// ClickLogin : click Login button
func (s LoginPage) ClickLogin() {
	s.Page.Click(webaction.ByID, btnLoginByID)

	return
}
