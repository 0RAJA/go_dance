package routing

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	root := r.Group("/")
	{
		Group.InitPostRouter(root)
		Group.InitTopicRouter(root)
	}
	return r
}
