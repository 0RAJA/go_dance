package routing

import (
	"github.com/gin-gonic/gin"
	v1 "go_dance/day_2/1_post/internel/api/v1"
)

type TopicRouter struct{}

func (router *TopicRouter) InitTopicRouter(group *gin.RouterGroup) {
	topicGroup := group.Group("topic")
	{
		topicGroup.POST("public", v1.Group.TopicApi.PublicTopic)
		topicGroup.GET("get", v1.Group.TopicApi.QueryTopic)
	}
}
