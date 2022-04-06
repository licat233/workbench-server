package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CustomJwtWhitelistModel interface {
	FindOneByJwt(ctx context.Context, jwt string) (*JwtWhitelists, error)
	FindAllByUidAndplatform(Uid int64, platform string) ([]*JwtWhitelists, error)
}

var cacheWorkbenchJwtWhitelistsJwtPrefix = "cache:workbench:jwtWhitelists:jwt:"

func (m *defaultJwtWhitelistsModel) FindOneByJwt(ctx context.Context, jwt string) (*JwtWhitelists, error) {
	workbenchJwtWhitelistsJwtKey := fmt.Sprintf("%s%s", cacheWorkbenchJwtWhitelistsJwtPrefix, jwt)
	var resp JwtWhitelists
	err := m.QueryRowCtx(ctx, &resp, workbenchJwtWhitelistsJwtKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `jwt` = ? limit 1", jwtWhitelistsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, jwt)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//FindAllHandler 多查询操作，不走缓存
func (m *defaultJwtWhitelistsModel) FindAllHandler(q string, args ...interface{}) ([]*JwtWhitelists, error) {
	resp := make([]*JwtWhitelists, 0)
	err := m.QueryRowsNoCache(&resp, q, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//FindAllByUidAndplatform 查找所有jwt，条件 uid 与 platform
func (m *defaultJwtWhitelistsModel) FindAllByUidAndplatform(Uid int64, platform string) ([]*JwtWhitelists, error) {
	query := fmt.Sprintf("select %s from %s where `platform` = ? ", jwtWhitelistsRows, m.table)
	return m.FindAllHandler(query, platform)
}
