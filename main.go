package main

import (
	"github.com/automation-go/bin/helper/common"
	"github.com/joho/godotenv"
)

func init() {
	env := godotenv.Load()
	common.LogPanicln(env)
}
