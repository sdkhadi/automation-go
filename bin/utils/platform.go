/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-09 00:05:52
 * @modify date 2021-08-09 00:05:52
 * @desc [description]
 */
package utils

import (
	webaction "github.com/automation-go/bin/helper/web/action"
	config "github.com/automation-go/bin/helper/web/config"
)

var Web webaction.Page

func WebConnect() error {
	config.DriverConnect()
	Web = webaction.Page{Action: config.Driver}

	return nil
}
