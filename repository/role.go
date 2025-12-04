package repository

import "KageNoEn/model"

func (r *Repository) GetAllRoles() ([]model.Role, error) {
	var roles []model.Role
	if err := r.DB.Find(&roles).Error; err != nil {
		return []model.Role{}, err
	}
	return roles, nil
}

func (r *Repository) GetRole(id string) (model.Role, error) {
	var role model.Role
	if err := r.DB.First(&role, id).Error; err != nil {
		return model.Role{}, err
	}
	return role, nil
}

func (r *Repository) CreateRole(role model.Role) error {
	if err := r.DB.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteRole(id string) error {
	if err := r.DB.Delete(&model.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}
