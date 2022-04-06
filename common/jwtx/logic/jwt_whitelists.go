package logic

import (
	"workbench-server/common/global"
	"workbench-server/common/jwtx/model"
)

var JwtWhitelist = NewJwtWhitelist()

func NewJwtWhitelist() model.JwtWhitelistsModel {
	return model.NewJwtWhitelistsModel(global.SQLConn, global.CONFIG.Cache.Redis)
}
