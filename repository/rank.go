package repository

import "KageNoEn/model"

func (r *Repository) GetAllRanks() ([]model.Rank, error) {
	var ranks []model.Rank
	if err := r.DB.Find(&ranks).Error; err != nil {
		return []model.Rank{}, err
	}

	return ranks, nil
}

func (r *Repository) GetRankById(id string) (model.Rank, error) {
	var rank model.Rank
	if err := r.DB.Where("id = ?", id).Find(&rank).Error; err != nil {
		return model.Rank{}, err
	}

	return rank, nil
}

func (r *Repository) CreateRank(rank model.Rank) error {
	if err := r.DB.Create(&rank).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateRank(rank *model.Rank) (error) {
	if err := r.DB.Where("id = ?", rank.Id).Updates(&rank).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteRank(id string) error {
	if err := r.DB.Delete(&model.Rank{}, id).Error; err != nil {
		return err
	}

	return nil
}
