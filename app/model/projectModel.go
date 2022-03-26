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
	projectFieldNames          = builder.RawFieldNames(&Project{})
	projectRows                = strings.Join(projectFieldNames, ",")
	projectRowsExpectAutoSet   = strings.Join(stringx.Remove(projectFieldNames, "`pid`", "`project_create_time`", "`project_update_time`"), ",")
	projectRowsWithPlaceHolder = strings.Join(stringx.Remove(projectFieldNames, "`pid`", "`project_create_time`", "`project_update_time`"), "=?,") + "=?"

	cacheWorkbenchProjectPidPrefix = "cache:workbench:project:pid:"
)

type (
	ProjectModel interface {
		baseProjectModel
		Insert(data *Project) (sql.Result, error)
		FindOne(pid int64) (*Project, error)
		Update(data *Project) error
		Delete(pid int64) error
	}

	defaultProjectModel struct {
		sqlc.CachedConn
		table string
	}

	Project struct {
		Pid                int64     `db:"pid"`                 // 主键id
		ProjectName        string    `db:"project_name"`        // 项目名称
		ProjectAddress     string    `db:"project_address"`     // 所在位置
		ProjectCompanyId   int64     `db:"project_company_id"`  // 所属公司id
		ProjectAdminId     int64     `db:"project_admin_id"`    // 项目负责人/项目经理
		ProjectStatus      int64     `db:"project_status"`      // 项目状态，默认值0，管理员审核后才能用
		ProjectCreateTime  time.Time `db:"project_create_time"` // 创建时间
		ProjectUpdateTime  time.Time `db:"project_update_time"` // 更新时间
		ProjectDescription string    `db:"project_description"` // 说明
	}
)

func NewProjectModel(conn sqlx.SqlConn, c cache.CacheConf) ProjectModel {
	return &defaultProjectModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`project`",
	}
}

func (m *defaultProjectModel) Insert(data *Project) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, projectRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.ProjectName, data.ProjectAddress, data.ProjectCompanyId, data.ProjectAdminId, data.ProjectStatus, data.ProjectCreateTime, data.ProjectUpdateTime, data.ProjectDescription)

	return ret, err
}

func (m *defaultProjectModel) FindOne(pid int64) (*Project, error) {
	workbenchProjectPidKey := fmt.Sprintf("%s%v", cacheWorkbenchProjectPidPrefix, pid)
	var resp Project
	err := m.QueryRow(&resp, workbenchProjectPidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `pid` = ? limit 1", projectRows, m.table)
		return conn.QueryRow(v, query, pid)
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

func (m *defaultProjectModel) Update(data *Project) error {
	workbenchProjectPidKey := fmt.Sprintf("%s%v", cacheWorkbenchProjectPidPrefix, data.Pid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `pid` = ?", m.table, projectRowsWithPlaceHolder)
		return conn.Exec(query, data.ProjectName, data.ProjectAddress, data.ProjectCompanyId, data.ProjectAdminId, data.ProjectStatus, data.ProjectCreateTime, data.ProjectUpdateTime, data.ProjectDescription, data.Pid)
	}, workbenchProjectPidKey)
	return err
}

func (m *defaultProjectModel) Delete(pid int64) error {

	workbenchProjectPidKey := fmt.Sprintf("%s%v", cacheWorkbenchProjectPidPrefix, pid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `pid` = ?", m.table)
		return conn.Exec(query, pid)
	}, workbenchProjectPidKey)
	return err
}

func (m *defaultProjectModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchProjectPidPrefix, primary)
}

func (m *defaultProjectModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `pid` = ? limit 1", projectRows, m.table)
	return conn.QueryRow(v, query, primary)
}
