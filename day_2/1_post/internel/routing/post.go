package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go_dance/day_2/1_post/internel/api/v1"
)

type PostRouter struct{}

func (router *PostRouter) InitPostRouter(group *gin.RouterGroup) {
	postGroup := group.Group("post")
	{
		postGroup.POST("public", v1.Group.PostApi.PublicPost)
	}
}
