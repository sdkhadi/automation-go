package websteps

import (
	webpages "github.com/automation-go/bin/testobjects/web"
	"github.com/automation-go/bin/utils"
)

func UserLogin() error {
	login := webpages.LoginPage{Page: utils.MobileWeb}

	login.InputUsername().
		InputPassword().
		ClickLogin()

	return nil
}
