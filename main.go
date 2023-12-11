package main

import (
	"mostafaba29/intialization"
	"mostafaba29/middleware"
)

func init() {
	intialization.ConnectToDB()
	intialization.MigrateDB()
	intialization.LoadDotEnv()
}

func main() {
	middleware.Setup()
}
