package service

import "blog-api/model"

type UserService struct{}

func (UserService) GetAllUsers() ([]model.User, error) {
	tests := make([]model.User, 0)
	err := DbEngine.Distinct("id", "first_name", "last_name").Limit(10, 0).Find(&tests)
	if err != nil {
		return nil, err
	}

	return tests, nil
}

func (UserService) CreateUser(user *model.User) error {
	_, err := DbEngine.Insert(user)
	if err != nil {
		return err
	}
	return nil
}
