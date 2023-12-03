package main

import (
	"mostafaba29/intialization"
	"mostafaba29/routes"
)

func init() {
	intialization.ConnectToDB()
	intialization.MigrateDB()
	intialization.LoadDotEnv()
}

func main() {
	routes.Setup()
}
