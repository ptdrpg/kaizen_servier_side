package repository

import "KageNoEn/model"

func (r *Repository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Preload("Rank").Preload("Role").Preload("Status").Find(&users).Error; err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (r *Repository) GetUserById(id string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("id = ?", id).Preload("Rank").Preload("Role").Preload("Status").Find(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("username = ?", username).Preload("Rank").Preload("Role").Preload("Status").Find(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) CreateUser(user model.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateUser(user model.User) error {
	if err := r.DB.Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteUser(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}
