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
	fansListFieldNames          = builder.RawFieldNames(&FansList{})
	fansListRows                = strings.Join(fansListFieldNames, ",")
	fansListRowsExpectAutoSet   = strings.Join(stringx.Remove(fansListFieldNames, "`fid`", "`create_time`", "`update_time`"), ",")
	fansListRowsWithPlaceHolder = strings.Join(stringx.Remove(fansListFieldNames, "`fid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchFansListFidPrefix = "cache:workbench:fansList:fid:"
)

type (
	FansListModel interface {
		baseFansListModel
		Insert(data *FansList) (sql.Result, error)
		FindOne(fid int64) (*FansList, error)
		Update(data *FansList) error
		Delete(fid int64) error
	}

	defaultFansListModel struct {
		sqlc.CachedConn
		table string
	}

	FansList struct {
		Fid                  int64     `db:"fid"`                     // Id主键
		FansDate             time.Time `db:"fans_date"`               // 加粉日期
		FansCreateTime       time.Time `db:"fans_create_time"`        // 创建时间
		FansUpdateTime       time.Time `db:"fans_update_time"`        // 更新时间
		DayFans              int64     `db:"day_fans"`                // 白天的加粉
		NightFans            int64     `db:"night_fans"`              // 夜间的加粉
		TodayFans            int64     `db:"today_fans"`              // 今天全部加粉
		ActiveFans           int64     `db:"active_fans"`             // 活跃的粉丝数
		FansRemark           string    `db:"fans_remark"`             // 备注
		FansServiceAccountId int64     `db:"fans_service_account_id"` // 所属账号id，外键
		FansAccountId        int64     `db:"fans_account_id"`         // 所属广告账户Id，外键
		FansChannelId        int64     `db:"fans_channel_id"`         // 所属广告渠道Id，外键
		FansProjectId        int64     `db:"fans_project_id"`         // 所属项目id，外键
		FansUserId           int64     `db:"fans_user_id"`            // 流量负责人
		FansStatus           int64     `db:"fans_status"`             // 状态，默认 1
	}
)

func NewFansListModel(conn sqlx.SqlConn, c cache.CacheConf) FansListModel {
	return &defaultFansListModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`fans_list`",
	}
}

func (m *defaultFansListModel) Insert(data *FansList) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, fansListRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.FansDate, data.FansCreateTime, data.FansUpdateTime, data.DayFans, data.NightFans, data.TodayFans, data.ActiveFans, data.FansRemark, data.FansServiceAccountId, data.FansAccountId, data.FansChannelId, data.FansProjectId, data.FansUserId, data.FansStatus)

	return ret, err
}

func (m *defaultFansListModel) FindOne(fid int64) (*FansList, error) {
	workbenchFansListFidKey := fmt.Sprintf("%s%v", cacheWorkbenchFansListFidPrefix, fid)
	var resp FansList
	err := m.QueryRow(&resp, workbenchFansListFidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `fid` = ? limit 1", fansListRows, m.table)
		return conn.QueryRow(v, query, fid)
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

func (m *defaultFansListModel) Update(data *FansList) error {
	workbenchFansListFidKey := fmt.Sprintf("%s%v", cacheWorkbenchFansListFidPrefix, data.Fid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `fid` = ?", m.table, fansListRowsWithPlaceHolder)
		return conn.Exec(query, data.FansDate, data.FansCreateTime, data.FansUpdateTime, data.DayFans, data.NightFans, data.TodayFans, data.ActiveFans, data.FansRemark, data.FansServiceAccountId, data.FansAccountId, data.FansChannelId, data.FansProjectId, data.FansUserId, data.FansStatus, data.Fid)
	}, workbenchFansListFidKey)
	return err
}

func (m *defaultFansListModel) Delete(fid int64) error {

	workbenchFansListFidKey := fmt.Sprintf("%s%v", cacheWorkbenchFansListFidPrefix, fid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `fid` = ?", m.table)
		return conn.Exec(query, fid)
	}, workbenchFansListFidKey)
	return err
}

func (m *defaultFansListModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchFansListFidPrefix, primary)
}

func (m *defaultFansListModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `fid` = ? limit 1", fansListRows, m.table)
	return conn.QueryRow(v, query, primary)
}
