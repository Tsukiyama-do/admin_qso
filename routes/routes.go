package routes

import (
	"../sessions"
	"../config"

	"net/http"
	"github.com/gin-gonic/gin"
)

const c_host string = "localhost:8081"

func Home(ctx *gin.Context) {
	var user *config.DummyUserModel

	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if !exists {
		println("Unhappy home")
		println("  sessionID: " + session.ID)
		session.Save()
		ctx.HTML(http.StatusOK, "index.html", gin.H{ "domainport": c_host, })
		return
	}

	user = buffer.(*config.DummyUserModel)
	println("Home sweet home")
	println("  sessionID: " + session.ID)
	println("  username: " + user.Username)
	println("  email: " + user.Email)

	session.Save()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"isLoggedIn": true,
		"username": user.Username,
		"email": user.Email,
		"domainport": c_host,
	})
}

func LogIn(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{ "domainport": c_host, })
}

func SignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{ "domainport": c_host, })
}

func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found", "domainport": c_host, })
}
