package postgresql

import (
	"server/core/infrastructure"
	"server/core/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db    sqlx.DB
	table string
}

const userSchema = `
	CREATE TABLE IF NOT EXISTS users (
		id BIGSERIAL PRIMARY KEY,
		session_token TEXT NOT NULL
	);`

func NewUserRepository(db sqlx.DB) infrastructure.UserRepository {
	db.MustExec(userSchema)
	return UserRepository{
		db:    db,
		table: "users",
	}
}

func (r UserRepository) Create(user models.User) (models.User, error) {
	_, err := r.db.NamedExec("INSERT INTO users (session_token) VALUES (:session_token)", user)
	if err == nil {
		updatedUser, err := r.GetBySessionToken(user.SessionToken)
		if err == nil {
			return updatedUser, nil
		}
	}
	return user, err
}

func (r UserRepository) GetBySessionToken(token string) (models.User, error) {
	user := models.User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE session_token=$1", token)
	return user, err
}
