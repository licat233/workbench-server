package jwtx

import (
	"workbench-server/common/global"
	"workbench-server/common/jwtx/model"
)

var JwtBlacklist = NewJwtBlacklist()

func NewJwtBlacklist() model.JwtBlacklistsModel {
	return model.NewJwtBlacklistsModel(global.SQLConn, global.CONFIG.Cache.Redis)
}
