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
	articleFieldNames          = builder.RawFieldNames(&Article{})
	articleRows                = strings.Join(articleFieldNames, ",")
	articleRowsExpectAutoSet   = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	articleRowsWithPlaceHolder = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchArticleIdPrefix = "cache:workbench:article:id:"
)

type (
	ArticleModel interface {
		baseArticleModel
		Insert(data *Article) (sql.Result, error)
		FindOne(id int64) (*Article, error)
		Update(data *Article) error
		Delete(id int64) error
	}

	defaultArticleModel struct {
		sqlc.CachedConn
		table string
	}

	Article struct {
		Id          int64          `db:"id"`           // Id主键
		Title       string         `db:"title"`        // 标题
		Keywords    string         `db:"keywords"`     // SEO中的关键词
		Description string         `db:"description"`  // SEO中的描述
		Html        string         `db:"html"`         // 内容
		StatusId    int64          `db:"status_id"`    // 状态，0表示未发布/草稿，1表示已发布，2表示已删除，默认 0
		ReleaseTime sql.NullTime   `db:"release_time"` // 发布时间
		CreateTime  time.Time      `db:"create_time"`  // 创建时间
		UpdateTime  time.Time      `db:"update_time"`  // 更新时间
		Screenshoot string         `db:"screenshoot"`  // 截图路径
		ProjectId   int64          `db:"project_id"`   // 所属项目id，外键
		GradeId     int64          `db:"grade_id"`     // 等级，外键，1优，2良，3中，4差
		AuthorId    int64          `db:"author_id"`    // 作者id，外键
		Safehtml    sql.NullString `db:"safehtml"`     // 加密html代码
		Remark      string         `db:"remark"`       // 对此文章的评论
	}
)

func NewArticleModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleModel {
	return &defaultArticleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`article`",
	}
}

func (m *defaultArticleModel) Insert(data *Article) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, articleRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Title, data.Keywords, data.Description, data.Html, data.StatusId, data.ReleaseTime, data.Screenshoot, data.ProjectId, data.GradeId, data.AuthorId, data.Safehtml, data.Remark)

	return ret, err
}

func (m *defaultArticleModel) FindOne(id int64) (*Article, error) {
	workbenchArticleIdKey := fmt.Sprintf("%s%v", cacheWorkbenchArticleIdPrefix, id)
	var resp Article
	err := m.QueryRow(&resp, workbenchArticleIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleRows, m.table)
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

func (m *defaultArticleModel) Update(data *Article) error {
	workbenchArticleIdKey := fmt.Sprintf("%s%v", cacheWorkbenchArticleIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, articleRowsWithPlaceHolder)
		return conn.Exec(query, data.Title, data.Keywords, data.Description, data.Html, data.StatusId, data.ReleaseTime, data.Screenshoot, data.ProjectId, data.GradeId, data.AuthorId, data.Safehtml, data.Remark, data.Id)
	}, workbenchArticleIdKey)
	return err
}

func (m *defaultArticleModel) Delete(id int64) error {

	workbenchArticleIdKey := fmt.Sprintf("%s%v", cacheWorkbenchArticleIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, workbenchArticleIdKey)
	return err
}

func (m *defaultArticleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchArticleIdPrefix, primary)
}

func (m *defaultArticleModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleRows, m.table)
	return conn.QueryRow(v, query, primary)
}
