package websteps

import (
	webpages "github.com/automation-go/bin/testobjects/web"
	"github.com/automation-go/bin/utils"
)

func UserIsInHomePage() error {
	web := webpages.LoginPage{Page: utils.Web}

	web.InputUsername().InputPassword().ClickLogin()

	return nil
}
