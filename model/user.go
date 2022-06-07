package model

import (
	"TikTok/db"
	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Name          string `gorm:"column:name;type:varchar(20);not null"`     // 用户名称
	Password      string `gorm:"column:password;type:varchar(20);not null"` //用户密码
	FollowCount   int64  `gorm:"column:follow_count;type:int;default:0"`    // 关注总数
	FollowerCount int64  `gorm:"column:follower_count;type:int;default:0"`  // 粉丝总数
}

// UserResp 响应结构体
type UserResp struct {
	Id            uint   `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

// ToResp 转化为响应结构体，默认关注
func (U User) ToResp() (UR UserResp) {
	UR.Id = U.ID
	UR.Name = U.Name
	UR.FollowCount = U.FollowerCount
	UR.FollowerCount = U.FollowerCount
	UR.IsFollow = true
	return UR
}

// IsFollowJudge 关注校验，视情况调用
func (UR *UserResp) IsFollowJudge(UserId uint) {
	var FR FollowRelation
	db.DB.Where("follower_id = ? AND user_id = ?", UR.Id, UserId).First(&FR)
	if FR.Id <= 0 {
		(*UR).IsFollow = false
	}
}
