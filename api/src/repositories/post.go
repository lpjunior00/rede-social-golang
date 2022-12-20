package repositories

import (
	"api/src/models"
	"database/sql"
)

type Post struct {
	db *sql.DB
}

// É como um construtor que recebe o banco de daods como parametro
func NewPostRepository(db *sql.DB) *Post {
	return &Post{db}
}

func (repository Post) Create(post models.Post) (uint64, error) {

	statement, erro := repository.db.Prepare("INSERT INTO posts (title, content, authorId) VALUES (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(post.Title, post.Content, post.AuthorId)
	if erro != nil {
		return 0, erro
	}

	postId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(postId), nil
}

func (repository Post) FindAll(userId uint64) ([]models.Post, error) {

	linhas, erro := repository.db.Query(`SELECT p.id, p.title, p.content, p.authorId, u.nickname, p.likes, p.creationDate FROM posts p 
										INNER JOIN users u ON p.authorId = u.id
										LEFT JOIN followers f ON p.authorId = f.userId
										WHERE u.id = ? or f.followerId = ?
										ORDER BY p.creationDate DESC`, userId, userId)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var posts []models.Post
	for linhas.Next() {

		var post models.Post
		if erro := linhas.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorNickname,
			&post.Likes,
			&post.CreationDate,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Post) FindSpecificPost(postId uint64) (models.Post, error) {

	linha, erro := repository.db.Query(`SELECT p.id, p.title, p.content, p.authorId, u.nickname, p.likes FROM posts p 
										INNER JOIN users u ON p.authorId = u.id
										WHERE p.id = ?`, postId)
	if erro != nil {
		return models.Post{}, erro
	}

	defer linha.Close()

	var post models.Post
	if linha.Next() {

		if erro := linha.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorNickname,
			&post.Likes,
		); erro != nil {
			return models.Post{}, erro
		}
	}

	return post, nil

}

func (repository Post) DeletePost(postId uint64) error {

	statement, erro := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(postId); erro != nil {
		return erro
	}

	return nil

}

func (repository Post) UpdatePost(postId uint64, post models.Post) error {

	statement, erro := repository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(post.Title, post.Content, postId); erro != nil {
		return erro
	}

	return nil
}

func (repository Post) FindPostsByUser(authorId uint64) ([]models.Post, error) {

	linhas, erro := repository.db.Query(`SELECT p.id, p.title, p.content, p.authorId, u.nickname, p.likes FROM posts p 
										INNER JOIN users u ON p.authorId = u.id
										WHERE p.authorId = ?`, authorId)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var posts []models.Post
	for linhas.Next() {

		var post models.Post
		if erro := linhas.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorNickname,
			&post.Likes,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)

	}

	return posts, nil
}

func (repository Post) LikePost(postId uint64) error {
	statement, erro := repository.db.Prepare("UPDATE posts SET likes=likes +1 WHERE id = ?")
	if erro != nil {
		return erro
	}

	if _, erro := statement.Exec(postId); erro != nil {
		return erro
	}

	return nil

}

func (repository Post) UnlikePost(postId uint64) error {
	statement, erro := repository.db.Prepare("UPDATE posts SET likes=CASE WHEN likes > 0 then likes -1 ELSE 0 end WHERE id = ?")
	if erro != nil {
		return erro
	}

	if _, erro := statement.Exec(postId); erro != nil {
		return erro
	}

	return nil
}
