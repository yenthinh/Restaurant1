package main

import (
	"restaurant/controller"
	"restaurant/models"
)

func main() {
	models.Migrate()
	controller.InitController()
}
