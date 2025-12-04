package repository

import "KageNoEn/model"

func (r *Repository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Preload("Rank").Preload("Role").Find(&users).Error; err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (r *Repository) GetUserById(id string) model.User {
	var user model.User
	if err := r.DB.Where("id = ?", id).Preload("Rank").Preload("Role").Find(&user).Error; err != nil {
		return model.User{}
	}

	return user
}

func (r *Repository) CreateUser(user model.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) Updateuser(user model.User) (model.User, error) {
	if err := r.DB.Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) DeleteUser(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}
