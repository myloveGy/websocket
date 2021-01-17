package api

import (
	"errors"
	"websocket/entity"
	"websocket/models"
	"websocket/repo"
	"websocket/request/admin"
	"websocket/utils"
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
	} else {
		query.Where("status", "!=", entity.TableStatusDelete)
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
func (u *UserService) Create(param *admin.UserCreate) (*models.User, error) {

	if err := u.ValidateAndHandleParams(param, 0); err != nil {
		return nil, err
	}

	user := &models.User{
		Username: param.Username,
		Phone:    param.Phone,
		Password: param.Password,
		Status:   entity.UserStatusActivate,
	}

	// 修改数据
	if err := u.userRepo.Create(user); err != nil {
		return nil, errors.New(entity.ErrUserCreate)
	}

	return user, nil
}

func (u *UserService) Update(param *admin.UserUpdate) error {
	// 查询数据
	if _, err := u.findOrFail(param.UserId.UserId); err != nil {
		return err
	}

	// 验证并处理参数
	if err := u.ValidateAndHandleParams(&param.UserCreate, param.UserId.UserId); err != nil {
		return err
	}

	// 修改数据
	if _, err := u.userRepo.Update(&models.User{
		UserId:   param.UserId.UserId,
		Username: param.Username,
		Phone:    param.Phone,
		Password: param.UserCreate.Password,
	}); err != nil {
		return errors.New(entity.ErrUserUpdate)
	}

	return nil
}

func (u *UserService) Online(param *admin.UserId) error {
	// 查询数据
	if _, err := u.findOrFail(param.UserId); err != nil {
		return err
	}

	// 修改数据
	if _, err := u.userRepo.Update(&models.User{UserId: param.UserId, Status: entity.UserStatusActivate}); err != nil {
		return errors.New(entity.ErrUserUpdateStatus)
	}

	return nil
}

func (u *UserService) Offline(param *admin.UserId) error {
	// 查询数据
	if _, err := u.findOrFail(param.UserId); err != nil {
		return err
	}

	// 修改数据
	if _, err := u.userRepo.Update(&models.User{UserId: param.UserId, Status: entity.UserStatusDisabled}); err != nil {
		return errors.New(entity.ErrUserUpdateStatus)
	}

	return nil
}

func (u *UserService) Delete(param *admin.UserId) error {
	// 查询数据
	if _, err := u.findOrFail(param.UserId); err != nil {
		return err
	}

	// 删除数据
	if _, err := u.userRepo.Update(&models.User{UserId: param.UserId, Status: entity.TableStatusDelete}); err != nil {
		return errors.New(entity.ErrUserUpdateStatus)
	}

	return nil
}

func (u *UserService) ValidateUsernameAndPhone(username, phone string, userId int64) error {
	if u.userRepo.ExistsUsername(username, userId) {
		return errors.New(entity.ErrUserUsernameExists)
	}

	if u.userRepo.ExistsPhone(phone, userId) {
		return errors.New(entity.ErrUserPhoneExists)
	}

	return nil
}

func (u *UserService) ValidateAndHandleParams(param *admin.UserCreate, userId int64) error {
	// 如果密码不为空，那么需要加密密码
	if password, err := u.handleValidatePassword(param.Password); err != nil {
		return err
	} else {
		param.Password = password
	}

	// 验证用户账号名称和手机号
	if err := u.ValidateUsernameAndPhone(param.Username, param.Phone, userId); err != nil {
		return err
	}

	return nil
}

func (u *UserService) handleValidatePassword(password string) (string, error) {
	// 如果密码不为空，那么需要加密密码
	if password != "" {
		if len(password) < 6 {
			return "", errors.New(entity.ErrUserPasswordLength)
		}

		// 加密密码
		if ps, err := utils.GeneratePassword(password); err != nil {
			return "", errors.New(entity.ErrUserPasswordEncode)
		} else {
			return string(ps), nil
		}
	}

	return "", nil
}

func (u *UserService) findOrFail(userId int64) (*models.User, error) {
	// 查询数据
	user := &models.User{UserId: userId}
	if err := u.userRepo.Find(user); err != nil {
		return nil, errors.New(entity.ErrUserNotExists)
	}

	return user, nil
}
