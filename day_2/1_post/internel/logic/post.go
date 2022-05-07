package logic

import (
	"github.com/gin-gonic/gin"
	"go_dance/day_2/1_post/internel/dao"
	"go_dance/day_2/1_post/internel/global"
	"go_dance/day_2/1_post/internel/model/request"
	"go_dance/day_2/1_post/internel/pkg/app/errcode"
)

type Post struct{}

func (post *Post) PublicPost(ctx *gin.Context, params *request.PublicPost) (int64, errcode.Err) {
	postData, err := dao.Store.QueryTopic(ctx, params.ParentId)
	if err != nil {
		global.Log.Println(err)
		return 0, errcode.ErrQuery
	}
	if postData == nil {
		return 0, errcode.NotFound
	}
	id, err := dao.Store.SavePost(ctx, params.ParentId, params.Content)
	if err != nil {
		global.Log.Println(err)
		return 0, errcode.ErrSave
	}
	return id, nil
}
