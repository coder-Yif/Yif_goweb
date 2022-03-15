package dto

import "awesomeProject3/model"

type UserDto struct {
	Username string
	Phone    string
}

func GetUser(user model.User) UserDto {
	return UserDto{
		user.Name,
		user.Phone,
	}
}
