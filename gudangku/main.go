package main

import (
	"gudangku/packages/database"
	"gudangku/routes"
)

func main() {
	// Run App
	database.Init()

	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":1323"))
}
