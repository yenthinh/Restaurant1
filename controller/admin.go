package controller

import (
	"fmt"
	"log"
	"net/http"
	"restaurant/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func ChefInfo(ctx *gin.Context) {
	var people []models.People
	db := models.OpenDB()
	models.GetPeoples(db, &people)
	ctx.HTML(http.StatusOK, "admin.html", people)
}
func AdminMenu(ctx *gin.Context) {
	var product []models.Product
	db := models.OpenDB()
	models.GetAllProduct(db, &product)
	log.Println(product)
	ctx.HTML(http.StatusOK, "admin_menu.html", product)
}
func AddMenu(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add_menu.html", nil)
}
func AddMenuAction(ctx *gin.Context) {
	fmt.Println("DA VAO MENU ACTION")
	var pd models.Product
	err := ctx.ShouldBind(&pd)
	if err != nil {
		panic(err)
	}
	fmt.Print(pd.Name)
	fmt.Print(pd.Ingredients)
	fmt.Print(pd.Price)
	db := models.OpenDB()
	err = models.CreateProduct(db, pd)
	if err != nil {
		panic(err)
	}
	ctx.Redirect(http.StatusSeeOther, "/admin/menu")
}
func AddChef(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add_chef.html", nil)
}
func AddChefAction(ctx *gin.Context) {
	fmt.Println("DA VAO CHEF ACTION")
	var pp models.People
	img, err := imageupload.Process(ctx.Request, "file")
	if err != nil {
		panic(err)
	}
	currentImage = img
	fmt.Print(pp.Names)
	fmt.Print(pp.Position)
	fmt.Print(pp.Picture, currentImage)
	db := models.OpenDB()
	err = models.CreatePeople(db, pp)
	if err != nil {
		panic(err)
	}
	ctx.Redirect(http.StatusSeeOther, "/admin")
}
func UpdateMenu(ctx *gin.Context) {
	fmt.Print("Update Menu")
	idStr := ctx.Param("id_product")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	db := models.OpenDB()
	models.UpdateHide(db, id, false)
	ctx.Redirect(http.StatusSeeOther, "/admin/menu")
}
func DisplayMenu(ctx *gin.Context) {
	fmt.Print("Display Menu")
	idStr := ctx.Param("id_product")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	db := models.OpenDB()
	models.DeleteHide(db, id, true)
	ctx.Redirect(http.StatusSeeOther, "/admin/menu")
}

func Image(ctx *gin.Context) {
	if currentImage == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	currentImage.Write(ctx.Writer)
}
func Thumbnail(ctx *gin.Context) {
	if currentImage == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
	t, err := imageupload.ThumbnailJPEG(currentImage, 600, 600, 100)
	if err != nil {
		panic(err)
	}
	t.Write(ctx.Writer)
}
