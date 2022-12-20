package models

import (
	"errors"
	"time"
)

type Post struct {
	Id             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorId       uint64    `json:"authorid,omitempty"`
	AuthorNickname string    `json:"authorNickname,omitempty"`
	Likes          uint64    `json:"likes"`
	CreationDate   time.Time `json:"creationdate,omitempty"`
}

func (post *Post) Prepare() error {

	if post.Title == "" {
		return errors.New("Title is a required field and cannot be empty")
	}

	if post.Content == "" {
		return errors.New("Content is a required field and cannot be empty")
	}

	if post.AuthorId == 0 {
		return errors.New("Author is a required field and cannot be empty")
	}

	return nil

}
