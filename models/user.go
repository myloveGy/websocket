package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	UserId      int64     `db:"user_id" json:"user_id"`
	Username    string    `db:"username" json:"username"`
	Phone       string    `db:"phone" json:"phone"`
	Password    string    `db:"password" json:"password"`
	Status      int       `db:"status" json:"status"`
	AccessToken string    `db:"access_token" json:"access_token"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

func (u *User) FindByUsername(db *sqlx.DB, username string) error {
	return db.Get(u, "select * from `user` where `username` = ?", username)
}

func (u *User) UpdateAccessToken(db *sqlx.DB) (int64, error) {
	result, err := db.Exec("update `user` SET `access_token` = ? WHERE `user_id` = ?", u.AccessToken, u.UserId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
