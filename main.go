package main

import (
	"github.com/guilhermelribeiroo/go-continuousintegration/database"
	"github.com/guilhermelribeiroo/go-continuousintegration/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
