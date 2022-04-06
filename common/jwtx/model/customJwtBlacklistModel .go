package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CustomJwtBlacklistModel interface {
	FindOneByJwt(ctx context.Context, jwt string) (*JwtBlacklists, error)
}

var cacheWorkbenchJwtBlacklistsJwtPrefix = "cache:workbench:jwtBlacklists:jwt:"

func (m *defaultJwtBlacklistsModel) FindOneByJwt(ctx context.Context, jwt string) (*JwtBlacklists, error) {
	workbenchJwtBlacklistsJwtKey := fmt.Sprintf("%s%s", cacheWorkbenchJwtBlacklistsJwtPrefix, jwt)
	var resp JwtBlacklists
	err := m.QueryRowCtx(ctx, &resp, workbenchJwtBlacklistsJwtKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `jwt` = ? limit 1", jwtBlacklistsRows, m.table)
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
