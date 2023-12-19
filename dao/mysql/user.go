package mysql

import (
	"bluebell1/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "Orimiya123"

func IsUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return UserExist
	}
	return
}

func InsertInto(user *models.User) (err error) {
	//对密码加密
	user.Password = encryptPassword(user.Password)
	//执行插入语句
	sqlStr := `insert into user (user_id, username, password, email) values (?,?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.Password, user.Email)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write([]byte(oPassword))
	return hex.EncodeToString(h.Sum(nil))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.UserName)
	if errors.Is(err, sql.ErrNoRows) {
		return UserNotExist
	}
	if err != nil {
		return
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return InvalidPassword
	}
	return
}
