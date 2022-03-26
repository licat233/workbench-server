package model

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"workbench-server/common/utils"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	cacheWorkbenchSysLogIpPrefix = "cache:workbench:sysLog:ip:"
	// cacheWorkbenchSysLogPlatformPrefix = "cache:workbench:sysLog:platform:"
)

type baseSysLogModel interface {
	Find(path, start, end, ip string, Current, PageSize int64) ([]*SysLog, error)
	FindLoginFailCount(ip string) (int64, error)
	FindCount() (int64, error)
	FindLoginCountByUid(uid int64) (int64, error)
	FindLastLoginByUid(uid int64) (*SysLog, error)
	FindLastLogin(uid int64, platform string) (*SysLog, error)
}

func (m *defaultSysLogModel) Find(path, start, end, ip string, Current, PageSize int64) ([]*SysLog, error) {
	var resp = make([]*SysLog, 0)
	var err error
	path = strings.TrimSpace(path)
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)
	ip = strings.TrimSpace(ip)
	hasDate := len(start) != 0 && len(end) != 0
	hasIp := len(ip) != 0
	hasPath := len(path) != 0

	//页码应该由第一页开始
	if Current <= 0 {
		Current = 1
	}

	//每页数量最多只能显示 1000 条
	if PageSize > 1000 {
		PageSize = 1000
	}
	//每页数量至少显示 5 条
	if PageSize <= 0 {
		PageSize = 5
	}

	//start和end时间处理
	if hasDate {
		if start, err = utils.StringToTimeString(start); err != nil {
			return nil, err
		}
		if end, err = utils.StringToTimeString(end); err != nil {
			return nil, err
		}
		endDate, err := utils.StringToDate(end)
		if err != nil {
			return nil, err
		}
		oneday, _ := time.ParseDuration("24h")
		endDate = endDate.Add(oneday)               //mysql查询不包含endDate，所以需要加上1天时间
		end = endDate.Format("2006-01-02 15:04:05") //格式化成mysql需要的时间字符串
	}
	var sqlquery, pathWhere string

	if hasPath {
		pathWhere = fmt.Sprintf("AND `path`=%s", path)
	}

	if !hasDate && !hasIp {
		var pathStr string
		if hasPath {
			pathStr = fmt.Sprintf("WHERE `path`='%s'", path)
		}
		sqlquery = fmt.Sprintf("select %s from %s %s limit %v,%v ", sysLogRows, m.table, pathStr, (Current-1)*PageSize, PageSize)
	} else if !hasDate && hasIp {
		sqlquery = fmt.Sprintf("select %s from %s where `ip`='%s' %s limit %v,%v", sysLogRows, m.table, ip, pathWhere, (Current-1)*PageSize, PageSize)
	} else if hasDate && !hasIp {
		sqlquery = fmt.Sprintf("select %s from %s where `create_time` between %s AND %s %s limit %v,%v", sysLogRows, m.table, start, end, pathWhere, (Current-1)*PageSize, PageSize)
	} else {
		sqlquery = fmt.Sprintf("select %s from %s where `create_time` between %s AND %s %s AND `ip`='%s' limit %v,%v", sysLogRows, m.table, start, end, pathWhere, ip, (Current-1)*PageSize, PageSize)
	}

	err = m.QueryRowsNoCache(&resp, sqlquery)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//FindLoginFailCount 获取指定IP当天登录错误的次数
func (m *defaultSysLogModel) FindLoginFailCount(ip string) (int64, error) {
	ip = strings.TrimSpace(ip)
	if len(ip) == 0 {
		return 0, errors.New("查询ip不能为空")
	}
	workbenchSysLogIPKey := fmt.Sprintf("%s%s", cacheWorkbenchSysLogIpPrefix, ip)
	var resp int64
	err := m.QueryRow(&resp, workbenchSysLogIPKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select COUNT(*) from %s WHERE `ip`= ? AND `status`=0 AND to_days(`create_time`) = to_days(CURDATE())", m.table)
		return conn.QueryRow(v, query, ip)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultSysLogModel) FindCount() (int64, error) {
	var count int64
	query := fmt.Sprintf("SELECT count(*) FROM %s", m.table)
	err := m.QueryRowNoCache(&count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultSysLogModel) FindLoginCountByUid(uid int64) (int64, error) {
	var count int64
	query := fmt.Sprintf("SELECT count(*) FROM %s WHERE path='/user/login' AND user_id = %d", m.table, uid)
	err := m.QueryRowNoCache(&count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultSysLogModel) FindLastLoginByUid(uid int64) (*SysLog, error) {
	var resp = &SysLog{}
	query := fmt.Sprintf("SELECT  %s FROM %s WHERE `user_id`=%d AND `status`=1 AND `path`='/user/login' ORDER BY create_time DESC LIMIT 1", sysLogRows, m.table, uid)
	err := m.QueryRowNoCache(resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysLogModel) FindLastLogin(uid int64, platform string) (*SysLog, error) {
	var resp = &SysLog{}
	query := fmt.Sprintf("SELECT  %s FROM %s WHERE `user_id`=%d AND `platform`='%s' AND `status`=1 AND `path`='/user/login' ORDER BY create_time DESC LIMIT 1", sysLogRows, m.table, uid, platform)
	err := m.QueryRowNoCache(resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
