package logic

import (
	"bluebell1/dao/mysql"
	"bluebell1/models"
	"bluebell1/pkg/jwt"
	"bluebell1/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户存不存在
	if err = mysql.IsUserExist(p.UserName); err != nil {
		return err
	}
	//生成UID
	userID := snowflake.GetID()
	user := &models.User{
		UserID:   userID,
		Password: p.PassWord,
		UserName: p.UserName,
		Email:    p.Email,
	}
	//将用户数据保存进数据库
	err = mysql.InsertInto(user)
	if err != nil {
		return err
	}
	return
}

func Login(p *models.ParamLogin) (Token string, err error) {
	user := &models.User{
		UserName: p.UserName,
		Password: p.PassWord,
	}
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.UserName)
}
