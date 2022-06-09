package function

import (
	"TikTok/dbfunc"
	"TikTok/model"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// CommentAction 评论操作
func CommentAction(c *gin.Context) (commentResp model.CommentResp, err error) {

	userId := uint(1) //后期更改

	videoId64, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
	videoId := uint(videoId64)
	actionType := c.Query("action_type")
	if actionType == "1" {
		commentText := c.Query("comment_text")
		tempComment := model.Comment{
			VideoId: videoId,
			UserId:  userId,
			Content: commentText,
		}
		err = dbfunc.AddComment(tempComment)
		if err != nil {
			return commentResp, err
		}
		tempComment.CreatedAt = time.Now()
		return tempComment.ToResp(userId), nil
	} else if actionType == "2" {
		commentId64, _ := strconv.ParseUint(c.Query("comment_id"), 10, 64)
		commentId := uint(commentId64)
		err = dbfunc.DeleteComment(commentId, userId, videoId)
		if err != nil {
			return commentResp, err
		}
		return commentResp, nil
	}
	return commentResp, errors.New("未知错误")

}

// CommentList 评论列表
func CommentList(c *gin.Context) (model.CommentListResponse, error) {

	userId := uint(1) //后期更改

	videoId64, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
	videoId := uint(videoId64)
	commentList, err := dbfunc.CommentList(videoId, userId)
	var commentListResponse model.CommentListResponse
	commentListResponse.CommentList = commentList
	if err != nil {
		return commentListResponse, err
	}
	return commentListResponse, nil
}
