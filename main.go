package main

import (
	"KachProxyAPI/config"
	"KachProxyAPI/server"
	"log"
)

func main() {
	log.Println("Initing Config...")
	config := config.InitConfig("theat")
	log.Println("Initing Scrapper...")
	scrapper := server.InitScrapper()
	log.Println("Initing Server...")
	httpServer := server.InitHttpServer(config, scrapper)

	httpServer.Start()
}