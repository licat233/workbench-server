package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	websiteConfigFieldNames          = builder.RawFieldNames(&WebsiteConfig{})
	websiteConfigRows                = strings.Join(websiteConfigFieldNames, ",")
	websiteConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(websiteConfigFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	websiteConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(websiteConfigFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchWebsiteConfigIdPrefix = "cache:workbench:websiteConfig:id:"
)

type (
	WebsiteConfigModel interface {
		Insert(data *WebsiteConfig) (sql.Result, error)
		FindOne(id int64) (*WebsiteConfig, error)
		Update(data *WebsiteConfig) error
		Delete(id int64) error
	}

	defaultWebsiteConfigModel struct {
		sqlc.CachedConn
		table string
	}

	WebsiteConfig struct {
		Id            int64        `db:"id"`             // id主键
		Title         string       `db:"title"`          // 标题
		Keywords      string       `db:"keywords"`       // 关键词
		Description   string       `db:"description"`    // 描述
		Domain        string       `db:"domain"`         // 域名
		Beian         string       `db:"beian"`          // 备案号
		DomainDuetime sql.NullTime `db:"domain_duetime"` // 域名到期日
		SslDuetime    sql.NullTime `db:"ssl_duetime"`    // SSL证书到期日
		Email         string       `db:"email"`          // 管理员邮箱
		AdminerId     int64        `db:"adminer_id"`     // 指定网站管理员id，外键
	}
)

func NewWebsiteConfigModel(conn sqlx.SqlConn, c cache.CacheConf) WebsiteConfigModel {
	return &defaultWebsiteConfigModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`website_config`",
	}
}

func (m *defaultWebsiteConfigModel) Insert(data *WebsiteConfig) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, websiteConfigRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Title, data.Keywords, data.Description, data.Domain, data.Beian, data.DomainDuetime, data.SslDuetime, data.Email, data.AdminerId)

	return ret, err
}

func (m *defaultWebsiteConfigModel) FindOne(id int64) (*WebsiteConfig, error) {
	workbenchWebsiteConfigIdKey := fmt.Sprintf("%s%v", cacheWorkbenchWebsiteConfigIdPrefix, id)
	var resp WebsiteConfig
	err := m.QueryRow(&resp, workbenchWebsiteConfigIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", websiteConfigRows, m.table)
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

func (m *defaultWebsiteConfigModel) Update(data *WebsiteConfig) error {
	workbenchWebsiteConfigIdKey := fmt.Sprintf("%s%v", cacheWorkbenchWebsiteConfigIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, websiteConfigRowsWithPlaceHolder)
		return conn.Exec(query, data.Title, data.Keywords, data.Description, data.Domain, data.Beian, data.DomainDuetime, data.SslDuetime, data.Email, data.AdminerId, data.Id)
	}, workbenchWebsiteConfigIdKey)
	return err
}

func (m *defaultWebsiteConfigModel) Delete(id int64) error {

	workbenchWebsiteConfigIdKey := fmt.Sprintf("%s%v", cacheWorkbenchWebsiteConfigIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchWebsiteConfigIdKey)
	return err
}

func (m *defaultWebsiteConfigModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchWebsiteConfigIdPrefix, primary)
}

func (m *defaultWebsiteConfigModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", websiteConfigRows, m.table)
	return conn.QueryRow(v, query, primary)
}
