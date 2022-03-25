package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	menuFieldNames          = builder.RawFieldNames(&Menu{})
	menuRows                = strings.Join(menuFieldNames, ",")
	menuRowsExpectAutoSet   = strings.Join(stringx.Remove(menuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	menuRowsWithPlaceHolder = strings.Join(stringx.Remove(menuFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchMenuIdPrefix = "cache:workbench:menu:id:"
)

type (
	MenuModel interface {
		Insert(data *Menu) (sql.Result, error)
		FindOne(id int64) (*Menu, error)
		Update(data *Menu) error
		Delete(id int64) error
		baseMenuModel
	}

	defaultMenuModel struct {
		sqlc.CachedConn
		table string
	}

	Menu struct {
		Id         int64     `db:"id"`          // 主键
		Name       string    `db:"name"`        // 菜单名称
		Path       string    `db:"path"`        // 菜单路径
		Status     int64     `db:"status"`      // 菜单状态，默认关闭0
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
		ParentId   int64     `db:"parent_id"`   // 父菜单id
	}
)

func NewMenuModel(conn sqlx.SqlConn, c cache.CacheConf) MenuModel {
	return &defaultMenuModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`menu`",
	}
}

func (m *defaultMenuModel) Insert(data *Menu) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, menuRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Name, data.Path, data.Status, data.ParentId)

	return ret, err
}

func (m *defaultMenuModel) FindOne(id int64) (*Menu, error) {
	workbenchMenuIdKey := fmt.Sprintf("%s%v", cacheWorkbenchMenuIdPrefix, id)
	var resp Menu
	err := m.QueryRow(&resp, workbenchMenuIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", menuRows, m.table)
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

func (m *defaultMenuModel) Update(data *Menu) error {
	workbenchMenuIdKey := fmt.Sprintf("%s%v", cacheWorkbenchMenuIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, menuRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Path, data.Status, data.ParentId, data.Id)
	}, workbenchMenuIdKey)
	return err
}

func (m *defaultMenuModel) Delete(id int64) error {

	workbenchMenuIdKey := fmt.Sprintf("%s%v", cacheWorkbenchMenuIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchMenuIdKey)
	return err
}

func (m *defaultMenuModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchMenuIdPrefix, primary)
}

func (m *defaultMenuModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", menuRows, m.table)
	return conn.QueryRow(v, query, primary)
}
