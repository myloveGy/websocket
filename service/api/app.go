package api

import (
	"errors"
	"websocket/entity"
	"websocket/models"
	"websocket/repo"
	"websocket/request/admin"
	"websocket/utils"
)

type AppService struct {
	appRepo *repo.App
}

func NewAppService(appRepo *repo.App) *AppService {
	return &AppService{appRepo: appRepo}
}

func (a *AppService) List(param *admin.AppSearch) ([]*models.App, int64, error) {
	userList := make([]*models.App, 0)
	query := a.appRepo.Builder(&userList)
	if param.AppId != "" {
		query.Where("app_id", param.AppId)
	}

	if param.Status > 0 {
		query.Where("status", param.Status)
	} else {
		query.Where("status", "!=", entity.TableStatusDelete)
	}

	if param.AppName != "" {
		query.Where("app_name", "like", "%"+param.AppName+"%")
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
func (a *AppService) Create(param *admin.AppCreate) (*models.App, error) {

	user := &models.App{
		AppId:     utils.Unique(""),
		AppName:   param.AppName,
		AppSecret: param.AppSecret,
		Status:    entity.AppStatusActivate,
	}

	// 修改数据
	if err := a.appRepo.Create(user); err != nil {
		return nil, errors.New(entity.ErrAppCreate)
	}

	return user, nil
}

func (a *AppService) Update(param *admin.AppUpdate) error {
	// 查询数据
	if err := a.findOrFail(param.Id); err != nil {
		return err
	}

	// 修改数据
	if _, err := a.appRepo.Update(&models.App{
		Id:        param.Id,
		AppName:   param.AppName,
		AppSecret: param.AppSecret,
	}); err != nil {
		return errors.New(entity.ErrAppUpdate)
	}

	return nil
}

func (a *AppService) Online(param *admin.AppIdStruct) error {
	// 查询数据
	if err := a.findOrFail(param.Id); err != nil {
		return err
	}
	// 修改数据
	if _, err := a.appRepo.Update(&models.App{Id: param.Id, Status: entity.AppStatusActivate}); err != nil {
		return errors.New(entity.ErrAppUpdateStatus)
	}

	return nil
}

func (a *AppService) Offline(param *admin.AppIdStruct) error {
	// 查询数据
	if err := a.findOrFail(param.Id); err != nil {
		return err
	}
	// 修改数据
	if _, err := a.appRepo.Update(&models.App{Id: param.Id, Status: entity.AppStatusDisabled}); err != nil {
		return errors.New(entity.ErrAppUpdateStatus)
	}

	return nil
}

func (a *AppService) Delete(param *admin.AppIdStruct) error {
	// 查询数据
	if err := a.findOrFail(param.Id); err != nil {
		return err
	}

	// 删除数据
	if _, err := a.appRepo.Update(&models.App{Id: param.Id, Status: entity.TableStatusDelete}); err != nil {
		return errors.New(entity.ErrAppDelete)
	}

	return nil
}

func (a *AppService) findOrFail(id int64) error {
	// 查询数据
	if err := a.appRepo.Find(&models.App{Id: id}); err != nil {
		return errors.New(entity.ErrAppNoTExists)
	}

	return nil
}
