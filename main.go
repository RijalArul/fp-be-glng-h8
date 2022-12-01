package main

import (
	"fp-be-glng-h8/configs"
	"fp-be-glng-h8/routes"
)

func main() {
	configs.StartDB()
	routes.Routes()
}
