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
	brandFieldNames          = builder.RawFieldNames(&Brand{})
	brandRows                = strings.Join(brandFieldNames, ",")
	brandRowsExpectAutoSet   = strings.Join(stringx.Remove(brandFieldNames, "`bid`", "`brand_create_time`", "`brand_update_time`"), ",")
	brandRowsWithPlaceHolder = strings.Join(stringx.Remove(brandFieldNames, "`bid`", "`brand_create_time`", "`brand_update_time`"), "=?,") + "=?"

	cacheWorkbenchBrandBidPrefix = "cache:workbench:brand:bid:"
)

type (
	BrandModel interface {
		baseBrandModel
		Insert(data *Brand) (sql.Result, error)
		FindOne(bid int64) (*Brand, error)
		Update(data *Brand) error
		Delete(bid int64) error
	}

	defaultBrandModel struct {
		sqlc.CachedConn
		table string
	}

	Brand struct {
		Bid                    int64     `db:"bid"`
		BrandName              string    `db:"brand_name"`                // 品牌名
		BrandTeacherId         int64     `db:"brand_teacher_id"`          // 品牌代言老师id
		BrandEmail             string    `db:"brand_email"`               // 联系邮箱
		BrandCreateTime        time.Time `db:"brand_create_time"`         // 品牌创建时间
		BrandUpdateTime        time.Time `db:"brand_update_time"`         // 品牌更新时间
		BrandOfficialAccountId int64     `db:"brand_official_account_id"` // 品牌官方账号
		BrandDescription       string    `db:"brand_description"`         // 品牌描述
		BrandWebsite           string    `db:"brand_website"`             // 品牌官网
		BrandProjectId         int64     `db:"brand_project_id"`          // 所属项目id，外键
		BrandStatus            int64     `db:"brand_status"`              // 状态，默认0，审核过后才能用
		BrandLogo              string    `db:"brand_logo"`                // 品牌logo链接
	}
)

func NewBrandModel(conn sqlx.SqlConn, c cache.CacheConf) BrandModel {
	return &defaultBrandModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`brand`",
	}
}

func (m *defaultBrandModel) Insert(data *Brand) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, brandRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.BrandName, data.BrandTeacherId, data.BrandEmail, data.BrandCreateTime, data.BrandUpdateTime, data.BrandOfficialAccountId, data.BrandDescription, data.BrandWebsite, data.BrandProjectId, data.BrandStatus, data.BrandLogo)

	return ret, err
}

func (m *defaultBrandModel) FindOne(bid int64) (*Brand, error) {
	workbenchBrandBidKey := fmt.Sprintf("%s%v", cacheWorkbenchBrandBidPrefix, bid)
	var resp Brand
	err := m.QueryRow(&resp, workbenchBrandBidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `bid` = ? limit 1", brandRows, m.table)
		return conn.QueryRow(v, query, bid)
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

func (m *defaultBrandModel) Update(data *Brand) error {
	workbenchBrandBidKey := fmt.Sprintf("%s%v", cacheWorkbenchBrandBidPrefix, data.Bid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `bid` = ?", m.table, brandRowsWithPlaceHolder)
		return conn.Exec(query, data.BrandName, data.BrandTeacherId, data.BrandEmail, data.BrandCreateTime, data.BrandUpdateTime, data.BrandOfficialAccountId, data.BrandDescription, data.BrandWebsite, data.BrandProjectId, data.BrandStatus, data.BrandLogo, data.Bid)
	}, workbenchBrandBidKey)
	return err
}

func (m *defaultBrandModel) Delete(bid int64) error {

	workbenchBrandBidKey := fmt.Sprintf("%s%v", cacheWorkbenchBrandBidPrefix, bid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `bid` = ?", m.table)
		return conn.Exec(query, bid)
	}, workbenchBrandBidKey)
	return err
}

func (m *defaultBrandModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchBrandBidPrefix, primary)
}

func (m *defaultBrandModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `bid` = ? limit 1", brandRows, m.table)
	return conn.QueryRow(v, query, primary)
}
