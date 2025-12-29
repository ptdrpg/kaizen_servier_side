package repository

import "KageNoEn/model"

func (r *Repository) GetAllFriends(userId string) ([]model.FriendWithStatus, error) {
	var friends []model.FriendWithStatus
	if err := r.DB.Table("friend_lists f").
		Select("friend.id, friend.username, friend.is_online, f.status").
		Joins("JOIN users friend ON friend.id = CASE WHEN f.sender = ? THEN f.receiver ELSE f.sender END", userId).
		Where("f.status = ? AND (f.sender = ? OR f.receiver = ?)", "accepted", userId, userId).
		Scan(&friends).Error; err != nil {
		return nil, err
	}
	return friends, nil
}

func (r *Repository) GetFriendRequest(receiverId string) ([]model.FriendRequestType, error) {
	var requests []model.FriendRequestType
	if err := r.DB.Table("friend_lists f").
		Select("f.id, u.username").
		Joins("JOIN users u ON u.id = f.sender").
		Where("f.receiver = ? AND f.status = ?", receiverId, "pending").
		Scan(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r *Repository) GetFilteredSearch(myID string, username string) ([]model.FriendWithStatus, error) {
	var users []model.FriendWithStatus

	err := r.DB.
		Table("users u").
		Select(`
			u.id,
			u.username,
			u.is_online,
			f.status
		`).
		Joins(`
			LEFT JOIN friend_lists f
			  ON (
			       (f.sender = ? AND f.receiver = u.id)
			    OR (f.receiver = ? AND f.sender = u.id)
			  )
		`, myID, myID).
		Where("u.username ILIKE ?", "%"+username+"%").
		Where("u.id != ?", myID).
		Scan(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}


func (r *Repository) AddFriend(data model.FriendList) error {
	if err := r.DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetInvitationByID(id string) (model.FriendList, error) {
	var invit model.FriendList
	if err := r.DB.Where("id = ?", id).First(&invit).Error; err != nil {
		return model.FriendList{}, err
	}
	return invit, nil
}

func (r *Repository) ConfirmFriend(invit model.FriendList) error {
	if err := r.DB.Where("id = ?", invit.Id).Updates(&invit).Error; err != nil {
		return err
	}
	return nil
}
