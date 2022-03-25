package model

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
)

type baseUserModel interface {
	FindOneByMobile(mobile string) (*User, error)
	IsHasByMobile(mobile string) (bool, error)
	FindAllByMobile(Current, PageSize int64, mobile string) ([]*UserInfo, error)

	FindOneByEmail(email string) (*User, error)
	IsHasByEmail(email string) (bool, error)
	FindAllByEmail(Current, PageSize int64, mobile string) ([]*UserInfo, error)

	FindOneByPassport(passport string) (*User, error)
	IsHasByPassport(passport string) (bool, error)
	FindAllByPassport(Current, PageSize int64, passport string) ([]*UserInfo, error)

	FindOneByRealName(name string) (*UserInfo, error)
	IsHasByRealName(name string) (bool, error)
	FindAllByRealName(Current, PageSize int64, name string) ([]*UserInfo, error)

	FindOneById(uid int64) (*UserInfo, error)
	InsertUserRegister(data *User) error
	BaseUpdate(data *UserPutData) error
	DeleteUser(uid int64) error
	FindAll(r *http.Request, Current, PageSize int64) ([]*User, error)
	UpdatePassword(uid int64, password string) error
	FindCount() (int64, error)
	UpdateLoginInfo(uid, loginCount int64, lastTime time.Time) error
	GenerateQueryCountSQL(r *http.Request) string
	GenerateQuerySQL(r *http.Request, Current, PageSize int64) string
	UpdateMobile(uid int64, mobile string) error
	UpdateEmail(uid int64, email string) error
}

//用于查询所有信息
type UserInfo struct {
	Uid          int64  `db:"uid"`            // UID主键
	UserRealName string `db:"user_real_name"` // 用户真实姓名
	UserAvatar   string `db:"user_avatar"`    // 用户头像
	UserGender   int64  `db:"user_gender"`    // 用户性别，男是1，女士0
	Passport     string `db:"passport"`       // 账号
	// Password        string       `db:"password"`         // 密码，不可见，因为均为md5密文
	UserEmail       string       `db:"user_email"`       // 邮箱地址
	UserMobile      string       `db:"user_mobile"`      // 电话号码
	UserRoleId      int64        `db:"user_role_id"`     // 外键，角色id
	UserRoleName    string       `db:"role_name"`        // 角色名称
	UserCompanyId   int64        `db:"user_company_id"`  // 所属公司id
	UserCompanyName string       `db:"company_name"`     // 所属公司名称
	UserQq          string       `db:"user_qq"`          // 用户QQ
	UserWechat      string       `db:"user_wechat"`      // 用户微信
	UserStatus      int64        `db:"user_status"`      // 账号状态(默认0不可用，需审核)
	UserCreateTime  time.Time    `db:"user_create_time"` // 创建时间
	UserUpdateTime  time.Time    `db:"user_update_time"` // 更新时间
	UserDescription string       `db:"user_description"` // 说明
	UserLoginCount  int64        `db:"user_login_count"` // 登录次数
	LastLoginTime   sql.NullTime `db:"last_login_time"`  // 最后一次登录时间
}

//用户更新自己信息
type UserPutData struct {
	Uid             int64  `db:"uid"`              // UID主键
	UserRealName    string `db:"user_real_name"`   // 用户真实姓名
	UserRoleId      int64  `db:"user_role_id"`     // 外键，角色id
	UserGender      int64  `db:"user_gender"`      // 用户性别，男是1，女士0
	UserQq          string `db:"user_qq"`          // 用户QQ
	UserWechat      string `db:"user_wechat"`      // 用户微信
	UserDescription string `db:"user_description"` // 说明
}

var (
	userinfoFieldNames = builder.RawFieldNames(&UserInfo{})
	userinfoRows       = strings.Join(userinfoFieldNames, ",")
	//var userRowsAllDataSet  = strings.Join(append(userFieldNames, "`role_name`", "`company_name`"), ",")
	defaultMoveField                 = []string{"`uid`", "`user_create_time`", "`user_update_time`", "`last_login_time`", "`user_login_count`"}
	userRowsForRegisterSet           = strings.Join(stringx.Remove(userFieldNames, "`uid`", "`user_create_time`", "`user_update_time`", "`last_login_time`", "`user_login_count`", "`user_description`", "`user_status`"), ",")
	UserPutDataFieldNames            = builder.RawFieldNames(&UserPutData{})
	UserPutDataFieldNamesRows        = strings.Join(UserPutDataFieldNames, ",")
	UserPutDataRowsWithPlaceHolder   = strings.Join(stringx.Remove(UserPutDataFieldNames, "`uid`"), "=?,") + "=?"
	cacheWorkbenchUserEmailPrefix    = "cache:workbench:user:email:"
	cacheWorkbenchUserMobilePrefix   = "cache:workbench:user:mobile:"
	cacheWorkbenchUserPassportPrefix = "cache:workbench:user:passport:"
	cacheWorkbenchUserRealNamePrefix = "cache:workbench:user:realname:"
)

func (m *defaultUserModel) formatMobile(mobile string) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchUserMobilePrefix, mobile)
}

func (m *defaultUserModel) formatPassport(passport string) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchUserPassportPrefix, passport)
}

func (m *defaultUserModel) formatRealname(realname string) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchUserRealNamePrefix, realname)
}

func (m *defaultUserModel) formatEmail(email string) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchUserEmailPrefix, email)
}

//用户数据有变动时，需要更新的缓存键
func (m *defaultUserModel) updateCacheKeys(user *User) []string {
	return []string{m.formatPrimary(user.Uid), m.formatMobile(user.UserMobile), m.formatPassport(user.Passport), m.formatRealname(user.UserRealName), m.formatEmail(user.UserEmail)}
}

//GenerateQueryCountSQL 生成用于分页查询用户数量的SQL语句
func (m *defaultUserModel) GenerateQueryCountSQL(r *http.Request) string {
	var query string = fmt.Sprintf("select count(*) from %s", m.table)
	var hasField bool = false
	var build strings.Builder
	build.WriteString(query)
	var queryFeilds = map[string]string{
		"Uid":           "uid",
		"UserRealName":  "user_real_name",
		"UserGender":    "user_gender",
		"UserRoleId":    "user_role_id",
		"UserCompanyId": "user_company_id",
		"UserEmail":     "user_email",
		"UserMobile":    "user_mobile",
		"UserQq":        "user_qq",
		"UserWechat":    "user_wechat",
	}
	for form, field := range queryFeilds {
		if r.Form.Has(form) {
			v := strings.TrimSpace(r.Form.Get(form))
			if len(v) != 0 {
				if !hasField {
					hasField = true
					build.WriteString(" where")
				}
				build.WriteString(fmt.Sprintf(" %s='%s'", field, v))
			}
		}
	}
	return build.String()
}

//GenerateQuerySQL 生成用于分页查询用户数据的SQL语句
func (m *defaultUserModel) GenerateQuerySQL(r *http.Request, Current, PageSize int64) string {
	var query string = fmt.Sprintf("select %s from %s", userRows, m.table)
	var hasField bool = false
	var build strings.Builder
	build.WriteString(query)
	var queryFeilds = map[string]string{
		"Uid":           "uid",
		"UserRealName":  "user_real_name",
		"UserGender":    "user_gender",
		"UserRoleId":    "user_role_id",
		"UserCompanyId": "user_company_id",
		"UserEmail":     "user_email",
		"UserMobile":    "user_mobile",
		"UserQq":        "user_qq",
		"UserWechat":    "user_wechat",
	}
	for form, field := range queryFeilds {
		if r.Form.Has(form) {
			v := strings.TrimSpace(r.Form.Get(form))
			if len(v) != 0 {
				if !hasField {
					hasField = true
					build.WriteString(" where")
				}
				build.WriteString(fmt.Sprintf(" %s='%s'", field, v))
			}
		}
	}
	build.WriteString(fmt.Sprintf(" limit %v,%v ", (Current-1)*PageSize, PageSize))
	return build.String()
}

//FindCount 查找所有用户量，不走缓存
func (m *defaultUserModel) FindCount() (int64, error) {
	var count int64
	query := fmt.Sprintf("SELECT count(*) FROM %s", m.table)
	err := m.QueryRowNoCache(&count, query)
	//当使用 count 查询的时候，不存在 sqlc.ErrNotFound
	return count, err
}

//FindAll 分页查询用户数据，，不走缓存
func (m *defaultUserModel) FindAll(r *http.Request, Current, PageSize int64) ([]*User, error) {
	resp := make([]*User, 0)
	query := m.GenerateQuerySQL(r, Current, PageSize)
	err := m.QueryRowsNoCache(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//FindAllHandler 多查询操作，不走缓存
func (m *defaultUserModel) FindAllHandler(q string, args ...interface{}) ([]*UserInfo, error) {
	resp := make([]*UserInfo, 0)
	err := m.QueryRowsNoCache(&resp, q, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//FindAllByMobile 查找用户所有信息，by手机号，不走缓存
func (m *defaultUserModel) FindAllByMobile(Current, PageSize int64, mobile string) ([]*UserInfo, error) {
	query := fmt.Sprintf("SELECT %s FROM %s AS u LEFT JOIN company AS c ON u.user_company_id = c.cid LEFT JOIN sys_role AS s ON u.user_role_id = s.rid WHERE u.user_mobile = ? limit %v,%v", userinfoRows, m.table, (Current-1)*PageSize, PageSize)
	return m.FindAllHandler(query, mobile)
}

//FindAllByEmail 查找用户所有信息，by邮箱，不走缓存
func (m *defaultUserModel) FindAllByEmail(Current, PageSize int64, email string) ([]*UserInfo, error) {
	query := fmt.Sprintf("SELECT %s FROM %s AS u LEFT JOIN company AS c ON u.user_company_id = c.cid LEFT JOIN sys_role AS s ON u.user_role_id = s.rid WHERE u.user_email = ? limit %v,%v", userinfoRows, m.table, (Current-1)*PageSize, PageSize)
	return m.FindAllHandler(query, email)
}

//FindAllByPassport 查找用户所有信息，by账号，不走缓存
func (m *defaultUserModel) FindAllByPassport(Current, PageSize int64, passport string) ([]*UserInfo, error) {
	query := fmt.Sprintf("SELECT %s FROM %s AS u LEFT JOIN company AS c ON u.user_company_id = c.cid LEFT JOIN sys_role AS s ON u.user_role_id = s.rid WHERE u.passport = ? limit %v,%v", userinfoRows, m.table, (Current-1)*PageSize, PageSize)
	return m.FindAllHandler(query, passport)
}

//FindAllByRealName 查找用户所有信息，by用户名，不走缓存
func (m *defaultUserModel) FindAllByRealName(Current, PageSize int64, name string) ([]*UserInfo, error) {
	query := fmt.Sprintf("SELECT %s FROM %s AS u LEFT JOIN company AS c ON u.user_company_id = c.cid LEFT JOIN sys_role AS s ON u.user_role_id = s.rid WHERE u.user_real_name = ? limit %v,%v", userinfoRows, m.table, (Current-1)*PageSize, PageSize)
	return m.FindAllHandler(query, name)
}

//单个条件查询，不走缓存
//FilterOneUser 返回基本信息
func (m *defaultUserModel) FilterOneUser(q string, args ...interface{}) (*User, error) {
	var resp User
	err := m.QueryRowNoCache(&resp, q, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//FilterOneUserInfo 返回详细信息
func (m *defaultUserModel) FilterOneUserInfo(q string, args ...interface{}) (*UserInfo, error) {
	var resp UserInfo
	err := m.QueryRowNoCache(&resp, q, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//FindOneById 获取用户详细的信息
func (m *defaultUserModel) FindOneById(uid int64) (*UserInfo, error) {
	query := fmt.Sprintf("SELECT %s FROM %s AS u LEFT JOIN company AS c ON u.user_company_id = c.cid LEFT JOIN sys_role AS s ON u.user_role_id = s.rid WHERE u.uid = ? LIMIT 1", userinfoRows, m.table)
	return m.FilterOneUserInfo(query, uid)
}

func (m *defaultUserModel) FindOneByRealName(name string) (*UserInfo, error) {
	query := fmt.Sprintf("SELECT %s FROM %s AS u LEFT JOIN `company` AS c ON u.user_company_id = c.cid LEFT JOIN `sys_role` AS s ON u.user_role_id = s.rid WHERE u.user_real_name = ? limit 1", userinfoRows, m.table)
	return m.FilterOneUserInfo(query, name)
}

func (m *defaultUserModel) FindOneByMobile(mobile string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `user_mobile` = ? limit 1", userRows, m.table)
	return m.FilterOneUser(query, mobile)
}

func (m *defaultUserModel) FindOneByEmail(email string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `user_email` = ? limit 1", userRows, m.table)
	return m.FilterOneUser(query, email)
}

func (m *defaultUserModel) FindOneByPassport(passport string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `passport` = ? limit 1", userRows, m.table)
	return m.FilterOneUser(query, passport)
}

//判断是否存在，不走缓存
//IsHasByMobile 判断手机号是否存在
func (m *defaultUserModel) IsExistHandler(err error) (bool, error) {
	switch err {
	case nil:
		return true, nil
	case sqlc.ErrNotFound:
		return false, nil
	default:
		return false, err
	}
}

func (m *defaultUserModel) IsHasByMobile(mobile string) (bool, error) {
	_, err := m.FindOneByMobile(mobile)
	return m.IsExistHandler(err)
}

//IsHasByPassport 判断账号是否存在
func (m *defaultUserModel) IsHasByPassport(passport string) (bool, error) {
	_, err := m.FindOneByPassport(passport)
	return m.IsExistHandler(err)
}

//IsHasByRealName 判断用户名是否存在
func (m *defaultUserModel) IsHasByRealName(name string) (bool, error) {
	_, err := m.FindOneByRealName(name)
	return m.IsExistHandler(err)
}

//IsHasByEmail 判断邮箱是否存在
func (m *defaultUserModel) IsHasByEmail(email string) (bool, error) {
	_, err := m.FindOneByEmail(email)
	return m.IsExistHandler(err)
}

//插入数据，不走缓存
//InsertUserRegister 插入用户数据，用于注册
func (m *defaultUserModel) InsertUserRegister(data *User) error {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsForRegisterSet)
	_, err := m.ExecNoCache(query, data.UserRealName, data.UserAvatar, data.UserGender, data.Passport, data.Password, data.UserEmail, data.UserMobile, data.UserRoleId, data.UserCompanyId, data.UserQq, data.UserWechat)
	return err
}

//删除数据，走缓存
//DeleteUser 删除用户，实际为更新操作，将user_status更改为2
func (m *defaultUserModel) DeleteUser(uid int64) error {
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `user_status`= 2 where `uid` = ?", m.table)
		return conn.Exec(query, uid)
	}, m.formatPrimary(uid))
	return err
}

//更新数据，走缓存
//UpdateLoginInfo 更新登录信息
func (m *defaultUserModel) UpdateLoginInfo(uid, loginCount int64, lastTime time.Time) error {
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `user_login_count` = ? ,`last_login_time` = ? where `uid` = ?", m.table)
		return conn.Exec(query, loginCount, lastTime, uid)
	}, m.formatPrimary(uid))
	return err
}

//UpdatePassword 更新密码
func (m *defaultUserModel) UpdatePassword(uid int64, password string) error {
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `password` = ? where `uid` = ?", m.table)
		return conn.Exec(query, password, uid)
	}, m.formatPrimary(uid))
	return err
}

//BaseUpdate 更新基础信息，用户自己更新自己的信息
func (m *defaultUserModel) BaseUpdate(data *UserPutData) error {
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `uid` = ?", m.table, UserPutDataRowsWithPlaceHolder)
		return conn.Exec(query, data.UserRealName, data.UserRoleId, data.UserGender, data.UserQq, data.UserWechat, data.UserDescription, data.Uid)
	}, m.formatPrimary(data.Uid))
	return err
}

//UpdateMobile 更改用户手机号
func (m *defaultUserModel) UpdateMobile(uid int64, mobile string) error {
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `user_mobile` = ? where `uid` = ?", m.table)
		return conn.Exec(query, mobile, uid)
	}, m.formatPrimary(uid))
	return err
}

//UpdateMobile 更改用户邮箱
func (m *defaultUserModel) UpdateEmail(uid int64, email string) error {
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `user_email` = ? where `uid` = ?", m.table)
		return conn.Exec(query, email, uid)
	}, m.formatPrimary(uid))
	return err
}
