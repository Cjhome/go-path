package user

import (
	"errors"
	"main.go/dto"
	"main.go/pkg/hash"
	user2 "main.go/repositories/user"
)

var NewHash = hash.Bcrypt{}

type UserService struct {

}

//Register 注册
func (user UserService) Register(userDto dto.UserDto) error {
	//密码加密
	bytes, err := NewHash.Make([]byte(userDto.Password))
	if err != nil {
		return errors.New(err.Error())
	}
	userDto.Password = string(bytes)
	return user2.CreateUser(userDto)
}

func (user UserService)Login(userDto dto.UserDto) error {
	// 根据昵称查询
	model := user2.GetUserByUsername(userDto.Username)
	if model.ID == 0 {
		return errors.New("账号不存在")
	}
	err := NewHash.Check([]byte(model.Password), []byte(userDto.Password))
	if err != nil {
		return errors.New("密码错误")
	}
	return nil
}
