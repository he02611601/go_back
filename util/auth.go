package auth

import (
	"errors"
)

var userData map[string]string

func init() {
	userData = map[string]string {
		"userName": "password",
	}
}

func checkUserIsExist(userName string) bool {
	_, isExist := userData[userName]
	return isExist
}

func checkUserPassword(inputPassword string, getPassword string) error {
	if inputPassword == getPassword {
		return nil
	}
	return errors.New("密碼錯誤喔！")
}

func Auth(username string, password string) error {
	if isExist := checkUserIsExist(username); isExist {
		return checkUserPassword(userData[username], password)
	} else {
		return errors.New("使用者不存在！")
	}
}
