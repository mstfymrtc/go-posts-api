package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/mstfymrtc/go-posts-api/utils"
)

type Post struct {
	gorm.Model
	Content string `json:"content"`
	UserId  uint   `json:"user_id"`
}

func (post *Post) Validate() (map[string]interface{}, bool) {
	if post.Content == "" {
		return u.Message(false, "Post content cannot be empty!"), false
	}
	if post.UserId <= 0 {
		return u.Message(false, "User not recognized!"), false
	}

	return u.Message(true, "Post validation successful"), true
}

func (post *Post) Create() (map[string]interface{}) {
	//post not valid
	if resp, ok := post.Validate(); !ok {
		return resp
	}

	GetDB().Create(post)

	resp := u.Message(true, "Post created successfully!")
	resp["post"] = post
	return resp

}

func GetPost(id uint) (*Post) {
	post := &Post{}
	err := GetDB().Table("posts").Where("id=?", id).First(post).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return post
}

func GetUserPosts(userId uint) ([]*Post) {
	posts := make([]*Post, 0)
	err := GetDB().Table("posts").Where("user_id=?", userId).Find(&posts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return posts
}
