/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-09 00:07:09
 * @modify date 2021-08-09 00:07:09
 * @desc [description]
 */
package common

import (
	"fmt"
	"log"
)

// LogPanicln : general error message using log.Panicln
func LogPanicln(err interface{}) error {
	if err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}
