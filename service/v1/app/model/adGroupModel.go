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
	adGroupFieldNames          = builder.RawFieldNames(&AdGroup{})
	adGroupRows                = strings.Join(adGroupFieldNames, ",")
	adGroupRowsExpectAutoSet   = strings.Join(stringx.Remove(adGroupFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	adGroupRowsWithPlaceHolder = strings.Join(stringx.Remove(adGroupFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchAdGroupIdPrefix = "cache:workbench:adGroup:id:"
)

type (
	AdGroupModel interface {
		baseAdGroupModel
		Insert(data *AdGroup) (sql.Result, error)
		FindOne(id int64) (*AdGroup, error)
		Update(data *AdGroup) error
		Delete(id int64) error
	}

	defaultAdGroupModel struct {
		sqlc.CachedConn
		table string
	}

	AdGroup struct {
		Id               int64     `db:"id"`                 // 主键
		GroupId          string    `db:"group_id"`           // 广告组id
		GroupStatus      int64     `db:"group_status"`       // 广告组状态，0关闭，1打开，3被封
		GroupName        string    `db:"group_name"`         // 广告组名称
		GroupCampaignUid string    `db:"group_campaign_uid"` // 广告组所属活动id
		GroupCpa         float64   `db:"group_cpa"`          // 广告组CPA
		GroupMaxCap      float64   `db:"group_max_cap"`      // 广告组CPA出价
		GroupDescription string    `db:"group_description"`  // 广告组描述
		GroupCost        float64   `db:"group_cost"`         // 广告组消耗
		GroupCtr         float64   `db:"group_ctr"`          // 广告组点击率
		GroupClick       int64     `db:"group_click"`        // 广告组点击数
		GroupCv          int64     `db:"group_cv"`           // 广告组转化数
		GroupCvr         float64   `db:"group_cvr"`          // 广告组转化率
		GroupBudget      float64   `db:"group_budget"`       // 广告组预算
		CreateTime       time.Time `db:"create_time"`        // 创建时间
		UpdateTime       time.Time `db:"update_time"`        // 更新时间
	}
)

func NewAdGroupModel(conn sqlx.SqlConn, c cache.CacheConf) AdGroupModel {
	return &defaultAdGroupModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`ad_group`",
	}
}

func (m *defaultAdGroupModel) Insert(data *AdGroup) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, adGroupRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.GroupId, data.GroupStatus, data.GroupName, data.GroupCampaignUid, data.GroupCpa, data.GroupMaxCap, data.GroupDescription, data.GroupCost, data.GroupCtr, data.GroupClick, data.GroupCv, data.GroupCvr, data.GroupBudget)

	return ret, err
}

func (m *defaultAdGroupModel) FindOne(id int64) (*AdGroup, error) {
	workbenchAdGroupIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdGroupIdPrefix, id)
	var resp AdGroup
	err := m.QueryRow(&resp, workbenchAdGroupIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adGroupRows, m.table)
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

func (m *defaultAdGroupModel) Update(data *AdGroup) error {
	workbenchAdGroupIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdGroupIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adGroupRowsWithPlaceHolder)
		return conn.Exec(query, data.GroupId, data.GroupStatus, data.GroupName, data.GroupCampaignUid, data.GroupCpa, data.GroupMaxCap, data.GroupDescription, data.GroupCost, data.GroupCtr, data.GroupClick, data.GroupCv, data.GroupCvr, data.GroupBudget, data.Id)
	}, workbenchAdGroupIdKey)
	return err
}

func (m *defaultAdGroupModel) Delete(id int64) error {

	workbenchAdGroupIdKey := fmt.Sprintf("%s%v", cacheWorkbenchAdGroupIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchAdGroupIdKey)
	return err
}

func (m *defaultAdGroupModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchAdGroupIdPrefix, primary)
}

func (m *defaultAdGroupModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adGroupRows, m.table)
	return conn.QueryRow(v, query, primary)
}
