/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-09 00:06:22
 * @modify date 2021-08-09 00:06:22
 * @desc [description]
 */
package main

import (
	"github.com/automation-go/bin/helper/common"
	"github.com/joho/godotenv"
)

func init() {
	env := godotenv.Load()
	common.LogPanicln(env)
}
func main() {

}
