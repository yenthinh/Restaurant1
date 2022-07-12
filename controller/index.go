package controller

import (
	"fmt"
	"net/http"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)
func Index(ctx *gin.Context) {
	type index struct{
		Products []models.Product
		Peoples  []models.People
	}
	var people []models.People
	var product []models.Product
	db := models.OpenDB()
	hide := true
	models.GetProducts(db, &product, hide)
	models.GetPeoples(db, &people)
	var data index
	data.Products = product
	data.Peoples=people
	ctx.HTML(http.StatusOK, "index.html", data) //nil == null
}
func ArrivalAction(ctx *gin.Context) {
	fmt.Print("DA VAO ARRIVAL")
}
