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
	adCampaignFieldNames          = builder.RawFieldNames(&AdCampaign{})
	adCampaignRows                = strings.Join(adCampaignFieldNames, ",")
	adCampaignRowsExpectAutoSet   = strings.Join(stringx.Remove(adCampaignFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	adCampaignRowsWithPlaceHolder = strings.Join(stringx.Remove(adCampaignFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchAdCampaignIdPrefix = "cache:workbench:adCampaign:id:"
)

type (
	AdCampaignModel interface {
		baseAdCampaignModel
		Insert(data *AdCampaign) (sql.Result, error)
		FindOne(id int64) (*AdCampaign, error)
		Update(data *AdCampaign) error
		Delete(id int64) error
	}

	defaultAdCampaignModel struct {
		sqlc.CachedConn
		table string
	}

	AdCampaign struct {
		Id                    int64          `db:"id"`                     // 活动id
		CampaignUid           string         `db:"campaign_uid"`           // 活动id值，最大64位
		AdcampaignStatus      int64          `db:"adcampaign_status"`      // 状态，默认1
		AdcampaignName        string         `db:"adcampaign_name"`        // 活动名称
		AdcampaignAccountId   int64          `db:"adcampaign_account_id"`  // 所属账户id，外键
		AdcampaignCookie      sql.NullString `db:"adcampaign_cookie"`      // cookie
		AdcampaignDescription string         `db:"adcampaign_description"` // 说明
		AdcampaignCreateTime  time.Time      `db:"adcampaign_create_time"` // 创建时间
		AdcampaignUpdateTime  time.Time      `db:"adcampaign_update_time"` // 更新时间
	}
)

func NewAdCampaignModel(conn sqlx.SqlConn, c cache.CacheConf) AdCampaignModel {
	return &defaultAdCampaignModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`ad_campaign`",
	}
}

func (m *defaultAdCampaignModel) Insert(data *AdCampaign) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, adCampaignRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.CampaignUid, data.AdcampaignStatus, data.AdcampaignName, data.AdcampaignAccountId, data.AdcampaignCookie, data.AdcampaignDescription, data.AdcampaignCreateTime, data.AdcampaignUpdateTime)

	return ret, err
}

func (m *defaultAdCampaignModel) FindOne(id int64) (*AdCampaign, error) {
	workbenchAdCampaignIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdCampaignIdPrefix, id)
	var resp AdCampaign
	err := m.QueryRow(&resp, workbenchAdCampaignIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adCampaignRows, m.table)
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

func (m *defaultAdCampaignModel) Update(data *AdCampaign) error {
	workbenchAdCampaignIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdCampaignIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adCampaignRowsWithPlaceHolder)
		return conn.Exec(query, data.CampaignUid, data.AdcampaignStatus, data.AdcampaignName, data.AdcampaignAccountId, data.AdcampaignCookie, data.AdcampaignDescription, data.AdcampaignCreateTime, data.AdcampaignUpdateTime, data.Id)
	}, workbenchAdCampaignIdKey)
	return err
}

func (m *defaultAdCampaignModel) Delete(id int64) error {

	workbenchAdCampaignIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdCampaignIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchAdCampaignIdKey)
	return err
}

func (m *defaultAdCampaignModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchAdCampaignIdPrefix, primary)
}

func (m *defaultAdCampaignModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adCampaignRows, m.table)
	return conn.QueryRow(v, query, primary)
}
