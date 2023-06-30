package main

import (
	"bubble/routes"
	"github.com/xylong/bingo"
)

func main() {
	bingo.Init().
		Mount("v1", routes.Controllers...)().
		Lunch(8080)
}
