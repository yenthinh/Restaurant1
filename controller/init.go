package controller

import (
	"github.com/gin-gonic/gin"
)

func InitController() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./views/assets/")

	router.GET("/", Index)
	router.GET("/admin/menu", AdminMenu)
	router.GET("/admin/addmenu", AddMenu)
	router.GET("/admin/addchef", AddChef)
	router.GET("/admin", ChefInfo)
	router.GET("/signin", Signin)
	router.GET("/login", Login)
	router.POST("/signin/action", SigninAction)
	router.POST("login/action", LoginAction)
	router.POST("/admin/addmenu/action", AddMenuAction)
	router.POST("/admin/addchef/action", AddChefAction)
	router.GET("/admin/arrival",Arrival)
	router.POST("/admin/arrival/action",ArrivalAction)
	router.POST("/admin/updatemenu/action/:id_product", UpdateMenu) //API
	router.POST("/admin/displaymenu/action/:id_product", DisplayMenu)
	router.Run()
}
