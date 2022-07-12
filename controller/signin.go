package controller

import (
	"fmt"
	"net/http"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)

func Signin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signin.html", nil) // nil==null
}

func SigninAction(ctx *gin.Context) {
	fmt.Println("DA VO DAY")
	var acc models.Account
	err := ctx.ShouldBind(&acc)
	if err != nil {
		panic(err)
	}
	fmt.Print(acc.Email)
	fmt.Print(acc.Password)
	db := models.OpenDB()
	err = models.CreateAccount(db, acc)
	if err != nil {
		panic(err)
	}
	ctx.Redirect(http.StatusSeeOther, "/login")
}
