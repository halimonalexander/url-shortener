package main

import (
	"fmt"
	"link_shortener/internal/config"
)

func main() {
	cfg := initConfig()
	fmt.Println(cfg)
	initLogger()
	initStorage()
	initRouter()

	// TODO run server
}

func initConfig() *config.Config {
	return config.MustLoad()
}

func initLogger() {

}

func initStorage() {

}

func initRouter() {

}
