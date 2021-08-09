/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-09 00:05:39
 * @modify date 2021-08-09 00:05:39
 * @desc [description]
 */
package websteps

import (
	webpages "github.com/automation-go/bin/testobjects/web"
	"github.com/automation-go/bin/utils"
)

func UserLogin() error {
	login := webpages.LoginPage{Page: utils.Web}

	login.InputUsername().
		InputPassword().
		ClickLogin()

	return nil
}
