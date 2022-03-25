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
	adCampaignCostFieldNames          = builder.RawFieldNames(&AdCampaignCost{})
	adCampaignCostRows                = strings.Join(adCampaignCostFieldNames, ",")
	adCampaignCostRowsExpectAutoSet   = strings.Join(stringx.Remove(adCampaignCostFieldNames, "`create_time`", "`update_time`"), ",")
	adCampaignCostRowsWithPlaceHolder = strings.Join(stringx.Remove(adCampaignCostFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchAdCampaignCostIdPrefix = "cache:workbench:adCampaignCost:id:"
)

type (
	AdCampaignCostModel interface {
		baseAdCampaignCostModel
		Insert(data *AdCampaignCost) (sql.Result, error)
		FindOne(id int64) (*AdCampaignCost, error)
		Update(data *AdCampaignCost) error
		Delete(id int64) error
	}

	defaultAdCampaignCostModel struct {
		sqlc.CachedConn
		table string
	}

	AdCampaignCost struct {
		Id          int64     `db:"id"`           // id主键
		CampaignId  int64     `db:"campaign_id"`  // 活动id，外键
		ADDate      time.Time `db:"AD_date"`      // 广告日期
		ADCost      float64   `db:"AD_cost"`      // 广告消耗
		CurrentCost float64   `db:"current_cost"` // 当日当前消耗
		NodeCost    float64   `db:"node_cost"`    // 广告日节点消耗
		YdayCost    float64   `db:"yday_cost"`    // 昨天全天消耗
		CreateTime  time.Time `db:"create_time"`  // 插入日期
	}
)

func NewAdCampaignCostModel(conn sqlx.SqlConn, c cache.CacheConf) AdCampaignCostModel {
	return &defaultAdCampaignCostModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`ad_campaign_cost`",
	}
}

func (m *defaultAdCampaignCostModel) Insert(data *AdCampaignCost) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, adCampaignCostRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.CampaignId, data.ADDate, data.ADCost, data.CurrentCost, data.NodeCost, data.YdayCost)

	return ret, err
}

func (m *defaultAdCampaignCostModel) FindOne(id int64) (*AdCampaignCost, error) {
	workbenchAdCampaignCostIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdCampaignCostIdPrefix, id)
	var resp AdCampaignCost
	err := m.QueryRow(&resp, workbenchAdCampaignCostIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adCampaignCostRows, m.table)
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

func (m *defaultAdCampaignCostModel) Update(data *AdCampaignCost) error {
	workbenchAdCampaignCostIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdCampaignCostIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adCampaignCostRowsWithPlaceHolder)
		return conn.Exec(query, data.CampaignId, data.ADDate, data.ADCost, data.CurrentCost, data.NodeCost, data.YdayCost, data.Id)
	}, workbenchAdCampaignCostIdKey)
	return err
}

func (m *defaultAdCampaignCostModel) Delete(id int64) error {

	workbenchAdCampaignCostIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdCampaignCostIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchAdCampaignCostIdKey)
	return err
}

func (m *defaultAdCampaignCostModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchAdCampaignCostIdPrefix, primary)
}

func (m *defaultAdCampaignCostModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adCampaignCostRows, m.table)
	return conn.QueryRow(v, query, primary)
}
