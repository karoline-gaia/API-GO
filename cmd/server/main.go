package main

import "github.com/karoline-gaia/API-GO/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
