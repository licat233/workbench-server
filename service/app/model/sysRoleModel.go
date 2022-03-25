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
	sysRoleFieldNames          = builder.RawFieldNames(&SysRole{})
	sysRoleRows                = strings.Join(sysRoleFieldNames, ",")
	sysRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleFieldNames, "`rid`", "`create_time`", "`update_time`"), ",")
	sysRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleFieldNames, "`rid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchSysRoleRidPrefix = "cache:workbench:sysRole:rid:"
)

type (
	SysRoleModel interface {
		Insert(data *SysRole) (sql.Result, error)
		FindOne(rid int64) (*SysRole, error)
		Update(data *SysRole) error
		Delete(rid int64) error
		baseSysRoleModel
	}

	defaultSysRoleModel struct {
		sqlc.CachedConn
		table string
	}

	SysRole struct {
		Rid             int64     `db:"rid"`              // 主键id
		RoleName        string    `db:"role_name"`        // 角色昵称，分为：游客、客服、客服组长、客服主管、美工、项目经理、文案、推广员、技术员、推广主管、人事、人事主管、会计员、跟单员、总经理
		RoleStatus      int64     `db:"role_status"`      // 状态，默认不可用
		RoleCreateTime  time.Time `db:"role_create_time"` // 创建时间
		RoleUpdateTime  time.Time `db:"role_update_time"` // 更新时间
		RoleDescription string    `db:"role_description"` // 说明
		RoleIcon        string    `db:"role_icon"`        // 角色图标
	}
)

func NewSysRoleModel(conn sqlx.SqlConn, c cache.CacheConf) SysRoleModel {
	return &defaultSysRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_role`",
	}
}

func (m *defaultSysRoleModel) Insert(data *SysRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysRoleRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.RoleName, data.RoleStatus, data.RoleCreateTime, data.RoleUpdateTime, data.RoleDescription, data.RoleIcon)

	return ret, err
}

func (m *defaultSysRoleModel) FindOne(rid int64) (*SysRole, error) {
	workbenchSysRoleRidKey := fmt.Sprintf("%s%v", cacheWorkbenchSysRoleRidPrefix, rid)
	var resp SysRole
	err := m.QueryRow(&resp, workbenchSysRoleRidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `rid` = ? limit 1", sysRoleRows, m.table)
		return conn.QueryRow(v, query, rid)
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

func (m *defaultSysRoleModel) Update(data *SysRole) error {
	workbenchSysRoleRidKey := fmt.Sprintf("%s%v", cacheWorkbenchSysRoleRidPrefix, data.Rid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `rid` = ?", m.table, sysRoleRowsWithPlaceHolder)
		return conn.Exec(query, data.RoleName, data.RoleStatus, data.RoleCreateTime, data.RoleUpdateTime, data.RoleDescription, data.RoleIcon, data.Rid)
	}, workbenchSysRoleRidKey)
	return err
}

func (m *defaultSysRoleModel) Delete(rid int64) error {

	workbenchSysRoleRidKey := fmt.Sprintf("%s%v", cacheWorkbenchSysRoleRidPrefix, rid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `rid` = ?", m.table)
		return conn.Exec(query, rid)
	}, workbenchSysRoleRidKey)
	return err
}

func (m *defaultSysRoleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchSysRoleRidPrefix, primary)
}

func (m *defaultSysRoleModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `rid` = ? limit 1", sysRoleRows, m.table)
	return conn.QueryRow(v, query, primary)
}
