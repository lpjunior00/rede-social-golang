package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type User struct {
	db *sql.DB
}

// É como um construtor que recebe o banco de daods como parametro
func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (repository User) Create(user models.User) (uint64, error) {

	statement, erro := repository.db.Prepare("INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	idCreated, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(idCreated), nil

}

func (repository User) Find(nameOrNickName string) ([]models.User, error) {
	nameOrNickName = fmt.Sprintf("%%%s%%", nameOrNickName)

	linhas, erro := repository.db.Query("SELECT id, name, nickname, email, creationdate FROM users WHERE name LIKE ? OR nickname LIKE ?", nameOrNickName, nameOrNickName)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	//Create a users slice
	var users []models.User

	//Iterate over lines found
	for linhas.Next() {
		var user models.User

		//for every line, scan and set into a user created
		if erro := linhas.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreationDate,
		); erro != nil {
			return nil, erro
		}

		//append the user into a list of users
		users = append(users, user)
	}

	return users, nil
}

func (repository *User) FindById(id uint64) (models.User, error) {
	linha, erro := repository.db.Query("SELECT id, name, nickname, email, creationdate FROM users WHERE id = ?", id)
	if erro != nil {
		return models.User{}, erro
	}

	defer linha.Close()

	var user models.User
	if linha.Next() {
		if erro = linha.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreationDate,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repository User) UpdateUser(user models.User, userId uint64) (models.User, error) {

	statement, erro := repository.db.Prepare("UPDATE users SET name = ?, nickName = ?, email = ? WHERE id = ?")
	if erro != nil {
		return models.User{}, erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nickname, user.Email, userId); erro != nil {
		return models.User{}, erro
	}

	return user, nil
}

func (repository User) DeleteUser(userId uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(userId); erro != nil {
		return nil
	}

	return nil
}

func (repository User) FindByEmail(email string) (models.User, error) {
	linha, erro := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}

	defer linha.Close()

	var user models.User
	if linha.Next() {
		if erro = linha.Scan(
			&user.ID,
			&user.Password,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repository *User) FollowUser(userId, followerId uint64) error {
	statement, erro := repository.db.Prepare("INSERT IGNORE INTO followers (userId, followerId) VALUES (?,?)")
	defer statement.Close()

	if erro != nil {
		return erro
	}

	if _, erro := statement.Exec(userId, followerId); erro != nil {
		return erro
	}

	return nil
}

func (repository *User) UnfollowUser(userId, followerId uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM followers WHERE userId = ? AND followerId = ?")
	defer statement.Close()

	if erro != nil {
		return erro
	}

	if _, erro := statement.Exec(userId, followerId); erro != nil {
		return erro
	}

	return nil
}

func (repository *User) FollowersByUser(userId uint64) ([]models.User, error) {
	linhas, erro := repository.db.Query(`select u.id, u.name, u.nickname, u.email, u.creationDate FROM followers f 
					  	   INNER JOIN users u ON f.followerId = u.id
	 					   WHERE userId = ?`, userId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var followers []models.User
	for linhas.Next() {
		var follower models.User

		if erro := linhas.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nickname,
			&follower.Email,
			&follower.CreationDate,
		); erro != nil {
			return nil, erro
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository *User) Following(userId uint64) ([]models.User, error) {
	linhas, erro := repository.db.Query(`select u.id, u.name, u.nickname, u.email, u.creationDate FROM followers f 
										INNER JOIN users u ON f.followerId = u.id
										WHERE followerId = ?`, userId)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var following []models.User
	for linhas.Next() {
		var user models.User
		if erro := linhas.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreationDate,
		); erro != nil {
			return nil, erro
		}

		following = append(following, user)
	}

	return following, nil
}

func (repository *User) UpdatePassword(userId uint64, newPassword string) error {
	statement, erro := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ? ")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(newPassword, userId); erro != nil {
		return erro
	}

	return nil
}

func (repository User) FindPassword(userId uint64) (string, error) {
	linha, erro := repository.db.Query("SELECT password FROM users WHERE id = ?", userId)
	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var password string
	if linha.Next() {
		if erro = linha.Scan(
			&password,
		); erro != nil {
			return "", erro
		}
	}

	return password, nil
}
