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
	landingPageFieldNames          = builder.RawFieldNames(&LandingPage{})
	landingPageRows                = strings.Join(landingPageFieldNames, ",")
	landingPageRowsExpectAutoSet   = strings.Join(stringx.Remove(landingPageFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	landingPageRowsWithPlaceHolder = strings.Join(stringx.Remove(landingPageFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchLandingPageIdPrefix = "cache:workbench:landingPage:id:"
)

type (
	LandingPageModel interface {
		baseLandingPageModel
		Insert(data *LandingPage) (sql.Result, error)
		FindOne(id int64) (*LandingPage, error)
		Update(data *LandingPage) error
		Delete(id int64) error
	}

	defaultLandingPageModel struct {
		sqlc.CachedConn
		table string
	}

	LandingPage struct {
		Id         int64          `db:"id"`
		AccountId  sql.NullInt64  `db:"account_id"`  // 所属广告账号id，外键
		Name       string         `db:"name"`        // 标题
		Url        string         `db:"url"`         // 链接
		ArticleId  sql.NullInt64  `db:"article_id"`  // 文章id，外键
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime time.Time      `db:"update_time"` // 更新时间
		JsCode     sql.NullString `db:"js_code"`     // Js代码
		JsFile     sql.NullString `db:"js_file"`     // Js文件路径
	}
)

func NewLandingPageModel(conn sqlx.SqlConn, c cache.CacheConf) LandingPageModel {
	return &defaultLandingPageModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`landing_page`",
	}
}

func (m *defaultLandingPageModel) Insert(data *LandingPage) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, landingPageRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.AccountId, data.Name, data.Url, data.ArticleId, data.JsCode, data.JsFile)

	return ret, err
}

func (m *defaultLandingPageModel) FindOne(id int64) (*LandingPage, error) {
	workbenchLandingPageIdKey := fmt.Sprintf("%s%v", cacheWorkbenchLandingPageIdPrefix, id)
	var resp LandingPage
	err := m.QueryRow(&resp, workbenchLandingPageIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", landingPageRows, m.table)
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

func (m *defaultLandingPageModel) Update(data *LandingPage) error {
	workbenchLandingPageIdKey := fmt.Sprintf("%s%v", cacheWorkbenchLandingPageIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, landingPageRowsWithPlaceHolder)
		return conn.Exec(query, data.AccountId, data.Name, data.Url, data.ArticleId, data.JsCode, data.JsFile, data.Id)
	}, workbenchLandingPageIdKey)
	return err
}

func (m *defaultLandingPageModel) Delete(id int64) error {

	workbenchLandingPageIdKey := fmt.Sprintf("%s%v", cacheWorkbenchLandingPageIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchLandingPageIdKey)
	return err
}

func (m *defaultLandingPageModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchLandingPageIdPrefix, primary)
}

func (m *defaultLandingPageModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", landingPageRows, m.table)
	return conn.QueryRow(v, query, primary)
}
