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
	sysLogFieldNames          = builder.RawFieldNames(&SysLog{})
	sysLogRows                = strings.Join(sysLogFieldNames, ",")
	sysLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchSysLogIdPrefix = "cache:workbench:sysLog:id:"
)

type (
	SysLogModel interface {
		Insert(data *SysLog) (sql.Result, error)
		FindOne(id int64) (*SysLog, error)
		Update(data *SysLog) error
		Delete(id int64) error
		baseSysLogModel
	}

	defaultSysLogModel struct {
		sqlc.CachedConn
		table string
	}

	SysLog struct {
		Id         int64     `db:"id"`
		Ip         string    `db:"ip"`
		Method     string    `db:"method"`      // 请求方法
		Path       string    `db:"path"`        // 路径，资源
		Status     int64     `db:"status"`      // 状态，0失败(默认)，1成功
		Message    string    `db:"message"`     // 信息
		Body       string    `db:"body"`        // 请求内容
		Resp       string    `db:"resp"`        // 应答内容
		UserId     int64     `db:"user_id"`     // 用户id
		CreateTime time.Time `db:"create_time"` // 创建时间
		JwtToken   string    `db:"jwt_token"`   // 用户jwt_token
		Platform   string    `db:"platform"`    // 设备平台
	}
)

func NewSysLogModel(conn sqlx.SqlConn, c cache.CacheConf) SysLogModel {
	return &defaultSysLogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_log`",
	}
}

func (m *defaultSysLogModel) Insert(data *SysLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysLogRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Ip, data.Method, data.Path, data.Status, data.Message, data.Body, data.Resp, data.UserId, data.JwtToken, data.Platform)

	return ret, err
}

func (m *defaultSysLogModel) FindOne(id int64) (*SysLog, error) {
	workbenchSysLogIdKey := fmt.Sprintf("%s%v", cacheWorkbenchSysLogIdPrefix, id)
	var resp SysLog
	err := m.QueryRow(&resp, workbenchSysLogIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLogRows, m.table)
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

func (m *defaultSysLogModel) Update(data *SysLog) error {
	workbenchSysLogIdKey := fmt.Sprintf("%s%v", cacheWorkbenchSysLogIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLogRowsWithPlaceHolder)
		return conn.Exec(query, data.Ip, data.Method, data.Path, data.Status, data.Message, data.Body, data.Resp, data.UserId, data.JwtToken, data.Platform, data.Id)
	}, workbenchSysLogIdKey)
	return err
}

func (m *defaultSysLogModel) Delete(id int64) error {

	workbenchSysLogIdKey := fmt.Sprintf("%s%v", cacheWorkbenchSysLogIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchSysLogIdKey)
	return err
}

func (m *defaultSysLogModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchSysLogIdPrefix, primary)
}

func (m *defaultSysLogModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLogRows, m.table)
	return conn.QueryRow(v, query, primary)
}
