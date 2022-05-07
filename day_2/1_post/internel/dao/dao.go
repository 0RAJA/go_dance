package dao

import (
	"context"
	"go_dance/day_2/1_post/internel/model/database"
)

type DB interface {
	SavePost(ctx context.Context, ParentId int64, Content string) (int64, error)
	QueryPostsByTopicID(ctx context.Context, topicID int64) ([]database.Post, error)
	SaveTopic(ctx context.Context, Title, Content string) (int64, error)
	QueryTopic(ctx context.Context, topicID int64) (*database.Topic, error)
}

type store struct {
	DB
}

var Store *store

func InitDB(db DB) {
	Store = &store{DB: db}
}
