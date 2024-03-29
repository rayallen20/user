package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	. "github.com/rayallen20/user/proto/user"
)

// User 用户对象
type User struct {
}

// Login 登录接口
func (u *User) Login(ctx context.Context, request *LoginRequest, response *LoginResponse) error {
	account := request.Account
	password := request.Password

	if account != "admin" {
		return errors.New("用户名不对")
	}

	if password != "NEWnew123!@#" {
		return errors.New("密码不对")
	}

	// 生成一个md5加密的字符串作为jwt
	hash := md5.New()
	hash.Write([]byte("testMd5"))
	md5Str := hex.EncodeToString(hash.Sum(nil))

	response.Jwt = md5Str
	response.Role = &Role{
		Id:   1,
		Name: "super_admin",
	}

	fmt.Println(1)
	return nil
}
