package user

import "main.go/dto"

type UserContract interface {
	Register(dto dto.UserDto) error
	Login(dto dto.UserDto) error
}
