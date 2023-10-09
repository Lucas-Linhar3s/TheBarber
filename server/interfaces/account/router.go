package account

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/create", CreatedAccount)
	r.POST("/login", Login)
}
