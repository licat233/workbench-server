package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	jwtWhitelistsFieldNames          = builder.RawFieldNames(&JwtWhitelists{})
	jwtWhitelistsRows                = strings.Join(jwtWhitelistsFieldNames, ",")
	jwtWhitelistsRowsExpectAutoSet   = strings.Join(stringx.Remove(jwtWhitelistsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	jwtWhitelistsRowsWithPlaceHolder = strings.Join(stringx.Remove(jwtWhitelistsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchJwtWhitelistsIdPrefix = "cache:workbench:jwtWhitelists:id:"
)

type (
	JwtWhitelistsModel interface {
		Insert(ctx context.Context, data *JwtWhitelists) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*JwtWhitelists, error)
		Update(ctx context.Context, data *JwtWhitelists) error
		Delete(ctx context.Context, id int64) error
		CustomJwtWhitelistModel
	}

	defaultJwtWhitelistsModel struct {
		sqlc.CachedConn
		table string
	}

	JwtWhitelists struct {
		Id          int64     `db:"id"`
		Uid         int64     `db:"uid"`          // 用户id
		Jwt         string    `db:"jwt"`          // jwt-token
		Platform    string    `db:"platform"`     // 绑定平台
		Ip          string    `db:"ip"`           // 绑定ip
		CreatedTime time.Time `db:"created_time"` // 创建时间
		UpdatedTime time.Time `db:"updated_time"` // 更新时间
		ExpireTime  time.Time `db:"expire_time"`  // 过期时间
	}
)

func NewJwtWhitelistsModel(conn sqlx.SqlConn, c cache.CacheConf) JwtWhitelistsModel {
	return &defaultJwtWhitelistsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`jwt_whitelists`",
	}
}

func (m *defaultJwtWhitelistsModel) Insert(ctx context.Context, data *JwtWhitelists) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, jwtWhitelistsRowsExpectAutoSet)
	ret, err := m.ExecNoCacheCtx(ctx, query, data.Uid, data.Jwt, data.Platform, data.Ip, data.CreatedTime, data.UpdatedTime, data.ExpireTime)

	return ret, err
}

func (m *defaultJwtWhitelistsModel) FindOne(ctx context.Context, id int64) (*JwtWhitelists, error) {
	workbenchJwtWhitelistsIdKey := fmt.Sprintf("%s%v", cacheWorkbenchJwtWhitelistsIdPrefix, id)
	var resp JwtWhitelists
	err := m.QueryRowCtx(ctx, &resp, workbenchJwtWhitelistsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", jwtWhitelistsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultJwtWhitelistsModel) Update(ctx context.Context, data *JwtWhitelists) error {
	workbenchJwtWhitelistsIdKey := fmt.Sprintf("%s%v", cacheWorkbenchJwtWhitelistsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, jwtWhitelistsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Jwt, data.Platform, data.Ip, data.CreatedTime, data.UpdatedTime, data.ExpireTime, data.Id)
	}, workbenchJwtWhitelistsIdKey)
	return err
}

func (m *defaultJwtWhitelistsModel) Delete(ctx context.Context, id int64) error {
	workbenchJwtWhitelistsIdKey := fmt.Sprintf("%s%v", cacheWorkbenchJwtWhitelistsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, workbenchJwtWhitelistsIdKey)
	return err
}

func (m *defaultJwtWhitelistsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchJwtWhitelistsIdPrefix, primary)
}

func (m *defaultJwtWhitelistsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", jwtWhitelistsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
