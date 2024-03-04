package authcontrollers

import (
	"net/http"

	"github.com/SaepudinJeh/go-gin-auth/src/models/user"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user user.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		ctx.Abort()
		return
	}
}
