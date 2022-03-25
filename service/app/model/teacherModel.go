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
	teacherFieldNames          = builder.RawFieldNames(&Teacher{})
	teacherRows                = strings.Join(teacherFieldNames, ",")
	teacherRowsExpectAutoSet   = strings.Join(stringx.Remove(teacherFieldNames, "`teacher_create_time`", "`teacher_update_time`"), ",")
	teacherRowsWithPlaceHolder = strings.Join(stringx.Remove(teacherFieldNames, "`tid`", "`teacher_create_time`", "`teacher_update_time`"), "=?,") + "=?"

	cacheWorkbenchTeacherTidPrefix = "cache:workbench:teacher:tid:"
)

type (
	TeacherModel interface {
		Insert(data *Teacher) (sql.Result, error)
		FindOne(tid int64) (*Teacher, error)
		Update(data *Teacher) error
		Delete(tid int64) error
		baseTeacherModel
	}

	defaultTeacherModel struct {
		sqlc.CachedConn
		table string
	}

	Teacher struct {
		Tid                int64          `db:"tid"`                 // 主键
		TeacherName        string         `db:"teacher_name"`        // 老师姓名
		TeacherEmail       string         `db:"teacher_email"`       // 老师邮箱
		TeacherMobile      string         `db:"teacher_mobile"`      // 老师电话
		TeacherCreateTime  time.Time      `db:"teacher_create_time"` // 创建时间
		TeacherUpdateTime  time.Time      `db:"teacher_update_time"` // 更新时间
		TeacherAvatar      string         `db:"teacher_avatar"`      // 老师头像
		TeacherGender      int64          `db:"teacher_gender"`      // 老师的性别，默认1（女性）
		TeacherDescription sql.NullString `db:"teacher_description"` // 老师简介
		TeacherWebsite     string         `db:"teacher_website"`     // 老师主页url
		TeacherStatus      int64          `db:"teacher_status"`      // 老师状态，默认0（不可用，审核过后才能用）
	}
)

func NewTeacherModel(conn sqlx.SqlConn, c cache.CacheConf) TeacherModel {
	return &defaultTeacherModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`teacher`",
	}
}

func (m *defaultTeacherModel) Insert(data *Teacher) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, teacherRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Tid, data.TeacherName, data.TeacherEmail, data.TeacherMobile, data.TeacherCreateTime, data.TeacherUpdateTime, data.TeacherAvatar, data.TeacherGender, data.TeacherDescription, data.TeacherWebsite, data.TeacherStatus)

	return ret, err
}

func (m *defaultTeacherModel) FindOne(tid int64) (*Teacher, error) {
	workbenchTeacherTidKey := fmt.Sprintf("%s%v", cacheWorkbenchTeacherTidPrefix, tid)
	var resp Teacher
	err := m.QueryRow(&resp, workbenchTeacherTidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `tid` = ? limit 1", teacherRows, m.table)
		return conn.QueryRow(v, query, tid)
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

func (m *defaultTeacherModel) Update(data *Teacher) error {
	workbenchTeacherTidKey := fmt.Sprintf("%s%v", cacheWorkbenchTeacherTidPrefix, data.Tid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `tid` = ?", m.table, teacherRowsWithPlaceHolder)
		return conn.Exec(query, data.TeacherName, data.TeacherEmail, data.TeacherMobile, data.TeacherCreateTime, data.TeacherUpdateTime, data.TeacherAvatar, data.TeacherGender, data.TeacherDescription, data.TeacherWebsite, data.TeacherStatus, data.Tid)
	}, workbenchTeacherTidKey)
	return err
}

func (m *defaultTeacherModel) Delete(tid int64) error {

	workbenchTeacherTidKey := fmt.Sprintf("%s%v", cacheWorkbenchTeacherTidPrefix, tid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `tid` = ?", m.table)
		return conn.Exec(query, tid)
	}, workbenchTeacherTidKey)
	return err
}

func (m *defaultTeacherModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchTeacherTidPrefix, primary)
}

func (m *defaultTeacherModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `tid` = ? limit 1", teacherRows, m.table)
	return conn.QueryRow(v, query, primary)
}
