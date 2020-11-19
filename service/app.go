package service

import (
	"sync"
	"websocket/models"
	"websocket/request"
)

type App struct {
	Id      int64
	AppId   string
	AppName string
	Online  int64
	Users   map[string][]*Client
	Groups  map[string]string
	mu      sync.RWMutex
}

func NewApp(app *models.App) *App {
	return &App{
		Id:      app.Id,
		AppId:   app.AppId,
		AppName: app.AppName,
		Users:   make(map[string][]*Client),
		Groups:  make(map[string]string),
	}
}

func (a *App) register(c *Client) {

	a.mu.Lock()

	if a.Users == nil {
		a.Users = make(map[string][]*Client)
	}

	a.Users[c.UserId] = append(a.Users[c.UserId], c)

	a.mu.Unlock()
}

func (a *App) unRegister(c *Client) {
	a.mu.Lock()
	var tmpUsers []*Client
	for _, v := range a.Users[c.UserId] {
		if v != c {
			tmpUsers = append(tmpUsers, v)
		}
	}

	if len(tmpUsers) == 0 {
		delete(a.Users, c.UserId)
	} else {
		a.Users[c.UserId] = tmpUsers
	}

	a.mu.Unlock()
}

func (a *App) ToItem() *request.AppItem {
	users := []string{}
	for userId := range a.Users {
		users = append(users, userId)
	}

	return &request.AppItem{
		AppId:       a.AppId,
		AppName:     a.AppName,
		OnlineUser:  len(a.Users),
		OnlineGroup: len(a.Groups),
		OnlineUsers: users,
	}
}
