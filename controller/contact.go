package controller

import (
	"fmt"
	"net/http"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)

func Contact(ctx *gin.Context) {
	var cc []models.Contact
	db := models.OpenDB()
	models.GetContacts(db, &cc)
	ctx.HTML(http.StatusOK, "contact.html", cc)
}
func ContactAction(ctx *gin.Context) {
	fmt.Print("DA VAO ARRIVAL")
	var cc models.Contact
	err := ctx.ShouldBind(&cc)
	if err != nil {
		panic(err)
	}
	fmt.Print(cc.Yname)
	fmt.Print(cc.Ymail)
	fmt.Print(cc.Topic)
	fmt.Print(cc.Ymess)
	db := models.OpenDB()
	err = models.CreateContact(db, cc)
	if err != nil {
		panic(err)
	}
	ctx.Redirect(http.StatusSeeOther, "/")
}
