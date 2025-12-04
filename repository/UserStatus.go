package repository

import "KageNoEn/model"

func (r *Repository) GetAllUserStatus() ([]model.UserStatus, error) {
	var userStatus []model.UserStatus
	if err := r.DB.Find(&userStatus).Error; err != nil {
		return []model.UserStatus{}, err
	}

	return userStatus, nil
}

func (r *Repository) GetUserStatusById(id string) (model.UserStatus, error) {
	var status model.UserStatus
	if err := r.DB.Where("id = ?", id).Find(&status).Error; err != nil {
		return model.UserStatus{}, err
	}

	return status, nil
}

func (r *Repository) CreateUserStatus(status model.UserStatus) error {
	if err := r.DB.Create(&status).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateUserStatus(status *model.UserStatus) error {
	if err := r.DB.Where("id = ?", status.Id).Updates(&status).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteUserStatus(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&model.UserStatus{}).Error; err != nil {
		return err
	}

	return nil
}
