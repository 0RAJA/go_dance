package v1

type group struct {
	PostApi
	TopicApi
}

var Group = new(group)
