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
	jwtBlacklistsFieldNames          = builder.RawFieldNames(&JwtBlacklists{})
	jwtBlacklistsRows                = strings.Join(jwtBlacklistsFieldNames, ",")
	jwtBlacklistsRowsExpectAutoSet   = strings.Join(stringx.Remove(jwtBlacklistsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	jwtBlacklistsRowsWithPlaceHolder = strings.Join(stringx.Remove(jwtBlacklistsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchJwtBlacklistsIdPrefix = "cache:workbench:jwtBlacklists:id:"
)

type (
	JwtBlacklistsModel interface {
		Insert(ctx context.Context, data *JwtBlacklists) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*JwtBlacklists, error)
		Update(ctx context.Context, data *JwtBlacklists) error
		Delete(ctx context.Context, id int64) error
		CustomJwtBlacklistModel
	}

	defaultJwtBlacklistsModel struct {
		sqlc.CachedConn
		table string
	}

	JwtBlacklists struct {
		Id          int64        `db:"id"`
		Uid         int64        `db:"uid"`          // 用户id
		Jwt         string       `db:"jwt"`          // jwt-token
		Platform    string       `db:"platform"`     // 绑定平台
		Ip          string       `db:"ip"`           // 绑定ip
		CreatedTime sql.NullTime `db:"created_time"` // 创建时间
		UpdatedTime time.Time    `db:"updated_time"` // 更新时间
		ExpireTime  sql.NullTime `db:"expire_time"`  // 过期时间
	}
)

func NewJwtBlacklistsModel(conn sqlx.SqlConn, c cache.CacheConf) JwtBlacklistsModel {
	return &defaultJwtBlacklistsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`jwt_blacklists`",
	}
}

func (m *defaultJwtBlacklistsModel) Insert(ctx context.Context, data *JwtBlacklists) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, jwtBlacklistsRowsExpectAutoSet)
	ret, err := m.ExecNoCacheCtx(ctx, query, data.Uid, data.Jwt, data.Platform, data.Ip, data.CreatedTime, data.UpdatedTime, data.ExpireTime)

	return ret, err
}

func (m *defaultJwtBlacklistsModel) FindOne(ctx context.Context, id int64) (*JwtBlacklists, error) {
	workbenchJwtBlacklistsIdKey := fmt.Sprintf("%s%v", cacheWorkbenchJwtBlacklistsIdPrefix, id)
	var resp JwtBlacklists
	err := m.QueryRowCtx(ctx, &resp, workbenchJwtBlacklistsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", jwtBlacklistsRows, m.table)
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

func (m *defaultJwtBlacklistsModel) Update(ctx context.Context, data *JwtBlacklists) error {
	workbenchJwtBlacklistsIdKey := fmt.Sprintf("%s%v", cacheWorkbenchJwtBlacklistsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, jwtBlacklistsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Jwt, data.Platform, data.Ip, data.CreatedTime, data.UpdatedTime, data.ExpireTime, data.Id)
	}, workbenchJwtBlacklistsIdKey)
	return err
}

func (m *defaultJwtBlacklistsModel) Delete(ctx context.Context, id int64) error {
	workbenchJwtBlacklistsIdKey := fmt.Sprintf("%s%v", cacheWorkbenchJwtBlacklistsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, workbenchJwtBlacklistsIdKey)
	return err
}

func (m *defaultJwtBlacklistsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchJwtBlacklistsIdPrefix, primary)
}

func (m *defaultJwtBlacklistsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", jwtBlacklistsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
