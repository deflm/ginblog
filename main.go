package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	model.InitDB()
	routes.InitRouter()
}
