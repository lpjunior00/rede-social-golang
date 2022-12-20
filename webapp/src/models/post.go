package models

import "time"

type Post struct {
	Id             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorId       uint64    `json:"authorId,omitempty"`
	AuthorNickname string    `json:"authorNickname,omitempty"`
	Likes          uint64    `json:"likes"`
	CreationDate   time.Time `json:"creationDate,omitempty"`
}
