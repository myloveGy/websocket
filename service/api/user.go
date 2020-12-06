package api

import (
	"errors"
	"websocket/entity"
	"websocket/models"
	"websocket/repo"
	"websocket/request/admin"
)

type UserService struct {
	userRepo *repo.User
}

func NewUserService(userRepo *repo.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) List(param *admin.UserSearch) ([]*models.User, int64, error) {
	userList := make([]*models.User, 0)
	query := u.userRepo.Builder(&userList)
	if param.UserId != "" {
		query.Where("user_id", param.UserId)
	}

	if param.Status > 0 {
		query.Where("status", param.Status)
	}

	if param.Username != "" {
		query.Where("username", "like", "%"+param.Username+"%")
	}

	if param.SortOrder != "" && param.SortField != "" {
		query.OrderBy(param.SortField, param.SortOrder)
	}

	total, err := query.Paginate(param.Page, param.PageSize)
	if err != nil {
		return nil, 0, err
	}

	return userList, total, nil
}

func (u *UserService) Online(param *admin.UserId) error {
	// 查询数据
	user := &models.User{UserId: param.UserId}
	if err := u.userRepo.Find(user); err != nil {
		return errors.New(entity.ErrUserNotExists)
	}

	// 修改数据
	if _, err := u.userRepo.Update(&models.User{UserId: user.UserId, Status: entity.UserStatusActivate}); err != nil {
		return errors.New(entity.ErrUserUpdateStatus)
	}

	return nil
}

func (u *UserService) Offline(param *admin.UserId) error {
	// 查询数据
	user := &models.User{UserId: param.UserId}
	if err := u.userRepo.Find(user); err != nil {
		return errors.New(entity.ErrUserNotExists)
	}

	// 修改数据
	if _, err := u.userRepo.Update(&models.User{UserId: user.UserId, Status: entity.UserStatusDisabled}); err != nil {
		return errors.New(entity.ErrUserUpdateStatus)
	}

	return nil
}
