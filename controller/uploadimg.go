package controller

import (
	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("avatar")
	if err != nil {
		panic(err)
	}
	fileName := file.Filename
	ctx.SaveUploadedFile(file, "./views/assets/img/chefs/" + fileName)
	url :="http://localhost:8080/assets/img/chefs/" + fileName
	ctx.Redirect(302, url)
}
