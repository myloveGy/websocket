package repo

import (
	"github.com/jmoiron/sqlx"

	"websocket/models"
	"websocket/utils"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	if err := u.db.Get(user, "select * from `user` where `username` = ?", username); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) UpdateAccessToken(userId int64, accessToken string) (int64, error) {
	result, err := u.db.Exec(
		"update `user` SET `access_token` = ?, updated_at = ?  WHERE `user_id` = ?",
		accessToken,
		utils.DateTime(),
		userId,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
