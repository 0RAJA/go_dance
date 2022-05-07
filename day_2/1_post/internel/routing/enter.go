package routing

type routingGroup struct {
	PostRouter
	TopicRouter
}

var Group = new(routingGroup)
