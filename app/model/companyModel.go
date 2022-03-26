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
	companyFieldNames          = builder.RawFieldNames(&Company{})
	companyRows                = strings.Join(companyFieldNames, ",")
	companyRowsExpectAutoSet   = strings.Join(stringx.Remove(companyFieldNames, "`cid`", "`company_create_time`", "`company_update_time`"), ",")
	companyRowsWithPlaceHolder = strings.Join(stringx.Remove(companyFieldNames, "`cid`", "`company_create_time`", "`company_update_time`"), "=?,") + "=?"

	cacheWorkbenchCompanyCidPrefix = "cache:workbench:company:cid:"
)

type (
	CompanyModel interface {
		Insert(data *Company) (sql.Result, error)
		FindOne(cid int64) (*Company, error)
		Update(data *Company) error
		Delete(cid int64) error
		baseCompanyModel
	}

	defaultCompanyModel struct {
		sqlc.CachedConn
		table string
	}

	Company struct {
		Cid                int64     `db:"cid"`                 // 公司id，主键
		CompanyName        string    `db:"company_name"`        // 公司名称
		CompanyAddress     string    `db:"company_address"`     // 公司地址
		CompanyBossId      int64     `db:"company_boss_id"`     // 公司负责人id
		CompanyStatus      int64     `db:"company_status"`      // 公司状态，默认0，需要审核过后才能用
		CompanyCreateTime  time.Time `db:"company_create_time"` // 创建时间
		CompanyUpdateTime  time.Time `db:"company_update_time"` // 更新时间
		CompanyDescription string    `db:"company_description"` // 公司简介
		CompanyLogo        string    `db:"company_logo"`        // 公司logo
		CompanyWebsite     string    `db:"company_website"`
	}
)

func NewCompanyModel(conn sqlx.SqlConn, c cache.CacheConf) CompanyModel {
	return &defaultCompanyModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`company`",
	}
}

func (m *defaultCompanyModel) Insert(data *Company) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, companyRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.CompanyName, data.CompanyAddress, data.CompanyBossId, data.CompanyStatus, data.CompanyCreateTime, data.CompanyUpdateTime, data.CompanyDescription, data.CompanyLogo, data.CompanyWebsite)

	return ret, err
}

func (m *defaultCompanyModel) FindOne(cid int64) (*Company, error) {
	workbenchCompanyCidKey := fmt.Sprintf("%s%v", cacheWorkbenchCompanyCidPrefix, cid)
	var resp Company
	err := m.QueryRow(&resp, workbenchCompanyCidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `cid` = ? limit 1", companyRows, m.table)
		return conn.QueryRow(v, query, cid)
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

func (m *defaultCompanyModel) Update(data *Company) error {
	workbenchCompanyCidKey := fmt.Sprintf("%s%v", cacheWorkbenchCompanyCidPrefix, data.Cid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `cid` = ?", m.table, companyRowsWithPlaceHolder)
		return conn.Exec(query, data.CompanyName, data.CompanyAddress, data.CompanyBossId, data.CompanyStatus, data.CompanyCreateTime, data.CompanyUpdateTime, data.CompanyDescription, data.CompanyLogo, data.CompanyWebsite, data.Cid)
	}, workbenchCompanyCidKey)
	return err
}

func (m *defaultCompanyModel) Delete(cid int64) error {

	workbenchCompanyCidKey := fmt.Sprintf("%s%v", cacheWorkbenchCompanyCidPrefix, cid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `cid` = ?", m.table)
		return conn.Exec(query, cid)
	}, workbenchCompanyCidKey)
	return err
}

func (m *defaultCompanyModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchCompanyCidPrefix, primary)
}

func (m *defaultCompanyModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `cid` = ? limit 1", companyRows, m.table)
	return conn.QueryRow(v, query, primary)
}
