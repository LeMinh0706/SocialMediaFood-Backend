package router

import "github.com/gin-gonic/gin"

func Static(r *gin.Engine) {
	r.Static("upload/post", "./upload/post")
	r.Static("upload/avatar", "./upload/avatar")
	r.Static("upload/background", "./upload/background")
}
