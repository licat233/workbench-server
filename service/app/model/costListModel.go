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
	costListFieldNames          = builder.RawFieldNames(&CostList{})
	costListRows                = strings.Join(costListFieldNames, ",")
	costListRowsExpectAutoSet   = strings.Join(stringx.Remove(costListFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	costListRowsWithPlaceHolder = strings.Join(stringx.Remove(costListFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchCostListIdPrefix = "cache:workbench:costList:id:"
)

type (
	CostListModel interface {
		baseCostListModel
		Insert(data *CostList) (sql.Result, error)
		FindOne(id int64) (*CostList, error)
		Update(data *CostList) error
		Delete(id int64) error
	}

	defaultCostListModel struct {
		sqlc.CachedConn
		table string
	}

	CostList struct {
		Id          int64     `db:"id"`            // Id主键
		CostDate    time.Time `db:"cost_date"`     // 消耗日期
		CreateTime  time.Time `db:"create_time"`   // 创建时间
		UpdateTime  time.Time `db:"update_time"`   // 更新时间
		CostTwd     float64   `db:"cost_twd"`      // 消耗台币
		CostCny     float64   `db:"cost_cny"`      // 消耗人民币
		Fans        int64     `db:"fans"`          // 加粉数
		Cpa         float64   `db:"cpa"`           // 每个粉的成本
		Remark      string    `db:"remark"`        // 备注
		ProjectId   int64     `db:"project_id"`    // 所属项目id，外键
		ChannelId   int64     `db:"channel_id"`    // 所属渠道id，外键
		AdAccountId int64     `db:"ad_account_id"` // 所属账户id，外键
		UserId      int64     `db:"user_id"`       // 负责人/操户人_id，外键
		BrandId     int64     `db:"brand_id"`
		Status      int64     `db:"status"` // 状态，默认1
	}
)

func NewCostListModel(conn sqlx.SqlConn, c cache.CacheConf) CostListModel {
	return &defaultCostListModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`cost_list`",
	}
}

func (m *defaultCostListModel) Insert(data *CostList) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, costListRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.CostDate, data.CostTwd, data.CostCny, data.Fans, data.Cpa, data.Remark, data.ProjectId, data.ChannelId, data.AdAccountId, data.UserId, data.BrandId, data.Status)

	return ret, err
}

func (m *defaultCostListModel) FindOne(id int64) (*CostList, error) {
	workbenchCostListIdKey := fmt.Sprintf("%s%v", cacheWorkbenchCostListIdPrefix, id)
	var resp CostList
	err := m.QueryRow(&resp, workbenchCostListIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", costListRows, m.table)
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

func (m *defaultCostListModel) Update(data *CostList) error {
	workbenchCostListIdKey := fmt.Sprintf("%s%v", cacheWorkbenchCostListIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, costListRowsWithPlaceHolder)
		return conn.Exec(query, data.CostDate, data.CostTwd, data.CostCny, data.Fans, data.Cpa, data.Remark, data.ProjectId, data.ChannelId, data.AdAccountId, data.UserId, data.BrandId, data.Status, data.Id)
	}, workbenchCostListIdKey)
	return err
}

func (m *defaultCostListModel) Delete(id int64) error {

	workbenchCostListIdKey := fmt.Sprintf("%s%v", cacheWorkbenchCostListIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchCostListIdKey)
	return err
}

func (m *defaultCostListModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchCostListIdPrefix, primary)
}

func (m *defaultCostListModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", costListRows, m.table)
	return conn.QueryRow(v, query, primary)
}
