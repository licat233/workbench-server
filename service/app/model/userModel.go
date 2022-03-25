package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, defaultMoveField...), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, defaultMoveField...), "=?,") + "=?"

	cacheWorkbenchUserUidPrefix = "cache:workbench:user:uid:"
)

type (
	UserModel interface {
		Insert(data *User) (sql.Result, error)
		FindOne(uid int64) (*User, error)
		Update(data *User) error
		Delete(uid int64) error
		baseUserModel
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Uid             int64        `db:"uid"`              // UID主键
		UserRealName    string       `db:"user_real_name"`   // 用户真实姓名
		UserAvatar      string       `db:"user_avatar"`      // 用户头像
		UserGender      int64        `db:"user_gender"`      // 用户性别，男是0，女是1
		Passport        string       `db:"passport"`         // 账号
		Password        string       `db:"password"`         // 密码
		UserEmail       string       `db:"user_email"`       // 邮箱地址
		UserMobile      string       `db:"user_mobile"`      // 电话号码
		UserRoleId      int64        `db:"user_role_id"`     // 外键，角色id
		UserCompanyId   int64        `db:"user_company_id"`  // 所属公司id
		UserQq          string       `db:"user_qq"`          // 用户QQ
		UserWechat      string       `db:"user_wechat"`      // 用户微信
		UserStatus      int64        `db:"user_status"`      // 账号状态(默认0不可用，需审核；1-可用，2-被删除)
		UserCreateTime  sql.NullTime `db:"user_create_time"` // 创建日期
		UserUpdateTime  sql.NullTime `db:"user_update_time"` // 升级日期
		UserDescription string       `db:"user_description"` // 说明
		UserLoginCount  int64        `db:"user_login_count"` // 登录次数
		LastLoginTime   sql.NullTime `db:"last_login_time"`  // 最后一次登录日期
		DeletedAt       sql.NullTime `db:"deleted_at"`       // 删除日期
	}
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserRealName, data.UserAvatar, data.UserGender, data.Passport, data.Password, data.UserEmail, data.UserMobile, data.UserRoleId, data.UserCompanyId, data.UserQq, data.UserWechat, data.UserStatus, data.UserCreateTime, data.UserUpdateTime, data.UserDescription, data.UserLoginCount, data.LastLoginTime, data.DeletedAt)

	return ret, err
}

func (m *defaultUserModel) FindOne(uid int64) (*User, error) {
	workbenchUserUidKey := fmt.Sprintf("%s%v", cacheWorkbenchUserUidPrefix, uid)
	var resp User
	err := m.QueryRow(&resp, workbenchUserUidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userRows, m.table)
		return conn.QueryRow(v, query, uid)
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

func (m *defaultUserModel) Update(data *User) error {
	workbenchUserUidKey := fmt.Sprintf("%s%v", cacheWorkbenchUserUidPrefix, data.Uid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `uid` = ?", m.table, userRowsWithPlaceHolder)
		return conn.Exec(query, data.UserRealName, data.UserAvatar, data.UserGender, data.Passport, data.Password, data.UserEmail, data.UserMobile, data.UserRoleId, data.UserCompanyId, data.UserQq, data.UserWechat, data.UserStatus, data.UserCreateTime, data.UserUpdateTime, data.UserDescription, data.UserLoginCount, data.LastLoginTime, data.DeletedAt, data.Uid)
	}, workbenchUserUidKey)
	return err
}

func (m *defaultUserModel) Delete(uid int64) error {

	workbenchUserUidKey := fmt.Sprintf("%s%v", cacheWorkbenchUserUidPrefix, uid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `uid` = ?", m.table)
		return conn.Exec(query, uid)
	}, workbenchUserUidKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchUserUidPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userRows, m.table)
	return conn.QueryRow(v, query, primary)
}
