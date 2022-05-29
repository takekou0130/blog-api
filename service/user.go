package service

import "blog-api/model"

type UserService struct{}

func (UserService) GetAllUsers() []model.User {
	tests := make([]model.User, 0)
	err := DbEngine.Distinct("id", "first_name", "last_name").Limit(10, 0).Find(&tests)
	if err != nil {
		return make([]model.User, 0)
	}

	return tests
}
