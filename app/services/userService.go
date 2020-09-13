package services

import (
	"../models"
	"../repositories"
)

var UserService _UserService

func init() {

}

func (_UserService) Login(user models.User) bool {
	_, err := repositories.UserRepository.GetUser(user)
	if err != nil {
		return false
	}
	return true
}
