package model

import (
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
	adChannelFieldNames          = builder.RawFieldNames(&AdChannel{})
	adChannelRows                = strings.Join(adChannelFieldNames, ",")
	adChannelRowsExpectAutoSet   = strings.Join(stringx.Remove(adChannelFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	adChannelRowsWithPlaceHolder = strings.Join(stringx.Remove(adChannelFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchAdChannelIdPrefix = "cache:workbench:adChannel:id:"
)

type (
	AdChannelModel interface {
		Insert(data *AdChannel) (sql.Result, error)
		FindOne(id int64) (*AdChannel, error)
		Update(data *AdChannel) error
		Delete(id int64) error
		baseAdChannelModel
	}

	defaultAdChannelModel struct {
		sqlc.CachedConn
		table string
	}

	AdChannel struct {
		Id          int64     `db:"id"`          // 渠道id，主键
		Name        string    `db:"name"`        // 渠道昵称
		Description string    `db:"description"` // 说明
		CreateTime  time.Time `db:"create_time"` // 创建时间
		UpdateTime  time.Time `db:"update_time"` // 更新时间
		Status      int64     `db:"status"`      // 状态，默认1
	}
)

func NewAdChannelModel(conn sqlx.SqlConn, c cache.CacheConf) AdChannelModel {
	return &defaultAdChannelModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`ad_channel`",
	}
}

func (m *defaultAdChannelModel) Insert(data *AdChannel) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, adChannelRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Name, data.Description, data.Status)

	return ret, err
}

func (m *defaultAdChannelModel) FindOne(id int64) (*AdChannel, error) {
	workbenchAdChannelIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdChannelIdPrefix, id)
	var resp AdChannel
	err := m.QueryRow(&resp, workbenchAdChannelIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adChannelRows, m.table)
		return conn.QueryRow(v, query, id)
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

func (m *defaultAdChannelModel) Update(data *AdChannel) error {
	workbenchAdChannelIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdChannelIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adChannelRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Description, data.Status, data.Id)
	}, workbenchAdChannelIdKey)
	return err
}

func (m *defaultAdChannelModel) Delete(id int64) error {

	workbenchAdChannelIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdChannelIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchAdChannelIdKey)
	return err
}

func (m *defaultAdChannelModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchAdChannelIdPrefix, primary)
}

func (m *defaultAdChannelModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adChannelRows, m.table)
	return conn.QueryRow(v, query, primary)
}
