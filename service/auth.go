package service

import (
	. "IxDServer/common"
	"IxDServer/db"
	"IxDServer/param/auth"
	"IxDServer/util"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

//1:注册
/*func AuthSignup(p *auth.SigninAndSignUp) (string, error) {
	//检查邮箱是否存在
	has, err := db.HasEmail(p.Email)
	if err != nil {
		return "", err
	}

	if has {
		return "", fmt.Errorf(MULTIPLE_STR)
	}

	//生成唯一用户id
	id := uuid.NewV4().String()

	//加密密码
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password+p.Email), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf(PASSWORD_BCRYPT_STR)
	}

	//保存用户
	err = db.InsertUser(id, p.Email, string(hash))
	if err != nil {
		return "", err
	}

	//生成token
	token, err := util.MarshalToken(id)
	if err != nil {
		return "", err
	}

	return token, nil
}*/

//2登录
func AuthSignin(p *auth.Signin) (string, error) {
	//检查用户
	user, err := db.SelectUserByEmail(p.Email)
	if err != nil {
		return "", err
	}
	if user["status"].(uint32) != 0 {
		return "", fmt.Errorf(LOCK_STR)
	}
	userId := user["id"].(string)

	//校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user["password"].(string)), []byte(p.Password+userId))
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf(PASSWORD_STR)
	}
	//token
	token, err := util.MarshalToken(userId)
	if err != nil {
		return "", err
	}
	return token, nil
}

//登录和注册
/*func AuthSigninAndSignup(p *auth.SigninAndSignUp) (string, error) {
	//检查邮箱
	has, err := db.HasEmail(p.Email)
	if err != nil {
		return "", err
	}

	if has {
		//登录
		return AuthSignin(p)
	} else {
		//注册
		return AuthSignup(p)
	}
}*/

//解密token
func AuthUnToken(token string) (string, error) {
	if token == "" || len(token) < 100 {
		return "", fmt.Errorf(AUTH_STR)
	}
	id, err := util.UnMarshalToken(token)
	if err != nil {
		return "", err
	}
	//TODO 权限
	return id, err
}

//续期
func AuthProlongToken(token string) (map[string]interface{}, error) {
	//解出id
	id, err := util.UnMarshalToken(token)
	if err != nil {
		return nil, err
	}

	//检查用户
	user, err := db.SelectUserById(id)
	if err != nil {
		return nil, err
	}

	if user["status"].(uint32) != 0 {
		return nil, fmt.Errorf(LOCK_STR)
	}

	//加密
	newToken, err := util.MarshalToken(id)
	if err != nil {
		return nil, err
	}

	//ip
	return map[string]interface{}{
		"token":    newToken,
		"datetime": time.Now().String(),
	}, nil
}
