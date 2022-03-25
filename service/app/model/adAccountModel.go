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
	adAccountFieldNames          = builder.RawFieldNames(&AdAccount{})
	adAccountRows                = strings.Join(adAccountFieldNames, ",")
	adAccountRowsExpectAutoSet   = strings.Join(stringx.Remove(adAccountFieldNames, "`aaid`", "`ad_account_create_time`", "`ad_account_update_time`"), ",")
	adAccountRowsWithPlaceHolder = strings.Join(stringx.Remove(adAccountFieldNames, "`aaid`", "`ad_account_create_time`", "`ad_account_update_time`"), "=?,") + "=?"

	cacheWorkbenchAdAccountAaidPrefix = "cache:workbench:adAccount:aaid:"
)

type (
	AdAccountModel interface {
		baseAdAccountModel
		Insert(data *AdAccount) (sql.Result, error)
		FindOne(aaid int64) (*AdAccount, error)
		Update(data *AdAccount) error
		Delete(aaid int64) error
	}

	defaultAdAccountModel struct {
		sqlc.CachedConn
		table string
	}

	AdAccount struct {
		Aaid                 int64          `db:"aaid"`                   // 账户id
		AdAccountChannelId   int64          `db:"ad_account_channel_id"`  // 所属渠道id，外键
		AdAccountPassport    string         `db:"ad_account_passport"`    // 账号
		AdAccountPassword    string         `db:"ad_account_password"`    // 密码
		AdAccountName        string         `db:"ad_account_name"`        // 账号名字
		AdAccountDescription string         `db:"ad_account_description"` // 说明
		AdAccountStatus      int64          `db:"ad_account_status"`      // 账号状态，默认值1
		AdAccountCreateTime  time.Time      `db:"ad_account_create_time"` // 创建时间
		AdAccountUpdateTime  time.Time      `db:"ad_account_update_time"` // 更新时间
		AdAccountCookie      sql.NullString `db:"ad_account_cookie"`      // cookie
		AdAccountToken       sql.NullString `db:"ad_account_token"`
		AdAccountReadCode    sql.NullString `db:"ad_account_read_code"`  // 阅读代码
		AdAccountCvCode      sql.NullString `db:"ad_account_cv_code"`    // 转化代码
		AdAccountOtherCode   sql.NullString `db:"ad_account_other_code"` // 其它代码
		AdAccountJsFile      string         `db:"ad_account_js_file"`    // Js文件地址
	}
)

func NewAdAccountModel(conn sqlx.SqlConn, c cache.CacheConf) AdAccountModel {
	return &defaultAdAccountModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`ad_account`",
	}
}

func (m *defaultAdAccountModel) Insert(data *AdAccount) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, adAccountRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.AdAccountChannelId, data.AdAccountPassport, data.AdAccountPassword, data.AdAccountName, data.AdAccountDescription, data.AdAccountStatus, data.AdAccountCreateTime, data.AdAccountUpdateTime, data.AdAccountCookie, data.AdAccountToken, data.AdAccountReadCode, data.AdAccountCvCode, data.AdAccountOtherCode, data.AdAccountJsFile)

	return ret, err
}

func (m *defaultAdAccountModel) FindOne(aaid int64) (*AdAccount, error) {
	workbenchAdAccountAaidKey := fmt.Sprintf("%s%v", cacheWorkbenchAdAccountAaidPrefix, aaid)
	var resp AdAccount
	err := m.QueryRow(&resp, workbenchAdAccountAaidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `aaid` = ? limit 1", adAccountRows, m.table)
		return conn.QueryRow(v, query, aaid)
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

func (m *defaultAdAccountModel) Update(data *AdAccount) error {
	workbenchAdAccountAaidKey := fmt.Sprintf("%s%v", cacheWorkbenchAdAccountAaidPrefix, data.Aaid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `aaid` = ?", m.table, adAccountRowsWithPlaceHolder)
		return conn.Exec(query, data.AdAccountChannelId, data.AdAccountPassport, data.AdAccountPassword, data.AdAccountName, data.AdAccountDescription, data.AdAccountStatus, data.AdAccountCreateTime, data.AdAccountUpdateTime, data.AdAccountCookie, data.AdAccountToken, data.AdAccountReadCode, data.AdAccountCvCode, data.AdAccountOtherCode, data.AdAccountJsFile, data.Aaid)
	}, workbenchAdAccountAaidKey)
	return err
}

func (m *defaultAdAccountModel) Delete(aaid int64) error {

	workbenchAdAccountAaidKey := fmt.Sprintf("%s%v", cacheWorkbenchAdAccountAaidPrefix, aaid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `aaid` = ?", m.table)
		return conn.Exec(query, aaid)
	}, workbenchAdAccountAaidKey)
	return err
}

func (m *defaultAdAccountModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchAdAccountAaidPrefix, primary)
}

func (m *defaultAdAccountModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `aaid` = ? limit 1", adAccountRows, m.table)
	return conn.QueryRow(v, query, primary)
}
