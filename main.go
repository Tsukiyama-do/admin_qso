package main

import (
	//"admin_qso/config"
	"./routes"
	"./sessions"
	"os"
  "github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {

	//  パラメータチェック
	    if len(os.Args) == 2 {
	      if os.Args[1] == "release" {
	        gin.SetMode(gin.ReleaseMode)
	      }
	    }

	//

	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")

	store := sessions.NewDummyStore()
	router.Use(sessions.StartDefaultSession(store))

	user := router.Group("/user")
	{
		user.POST("/signup", routes.UserSignUp)
		user.POST("/login", routes.UserLogIn)
		user.POST("/logout", routes.UserLogOut)
	}

	qsl := router.Group("/qsl")
	{
		qsl.GET("/qslmain", routes.Qslmain)
		qsl.POST("/qslselect", routes.Qslselectp)
		qsl.POST("/qslupddel", routes.Qslupddelp)
		qsl.POST("/qslinsert", routes.Qslinsertp)
		qsl.GET("/uploads", routes.Qsluploads)
		qsl.POST("/fileselect", routes.Fileselectp)
		qsl.POST("/fileupload", routes.Fileuploadp)

	}

	router.GET("/", routes.Home)
	router.GET("/login", routes.LogIn)
	router.GET("/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)

//	router.Run(":8081")

//  Start server with graceful shutdown function
	endless.ListenAndServe(":50005", router)

}
