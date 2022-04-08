package jwtx

import (
	"workbench-server/common/global"
	"workbench-server/common/jwtx/model"
)

// var JwtWhitelist model.JwtWhitelistsModel

func NewJwtWhitelist() model.JwtWhitelistsModel {
	return model.NewJwtWhitelistsModel(global.SQLConn, global.CONFIG.Cache.Redis)
}
