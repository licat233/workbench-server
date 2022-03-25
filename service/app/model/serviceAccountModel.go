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
	serviceAccountFieldNames          = builder.RawFieldNames(&ServiceAccount{})
	serviceAccountRows                = strings.Join(serviceAccountFieldNames, ",")
	serviceAccountRowsExpectAutoSet   = strings.Join(stringx.Remove(serviceAccountFieldNames, "`said`", "`create_time`", "`update_time`"), ",")
	serviceAccountRowsWithPlaceHolder = strings.Join(stringx.Remove(serviceAccountFieldNames, "`said`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheWorkbenchServiceAccountSaidPrefix = "cache:workbench:serviceAccount:said:"
)

type (
	ServiceAccountModel interface {
		baseServiceAccountModel
		Insert(data *ServiceAccount) (sql.Result, error)
		FindOne(said int64) (*ServiceAccount, error)
		Update(data *ServiceAccount) error
		Delete(said int64) error
	}

	defaultServiceAccountModel struct {
		sqlc.CachedConn
		table string
	}

	ServiceAccount struct {
		Said          int64     `db:"said"`            // Id主键
		SaUserId      int64     `db:"sa_user_id"`      // 使用者id，外键
		SaBrandId     int64     `db:"sa_brand_id"`     // 所属品牌id，外键，一般不会修改
		SaAccountId   int64     `db:"sa_account_id"`   // 所在的广告账户id，外键
		SaShowAccount string    `db:"sa_show_account"` // 展示业务号，常用于Line渠道，每个业务号可额外起个昵称
		SaAccount     string    `db:"sa_account"`      // 业务号ID
		SaPassword    string    `db:"sa_password"`     // 业务号密码
		SaCreateTime  time.Time `db:"sa_create_time"`  // 创建时间
		SaUpdateTime  time.Time `db:"sa_update_time"`  // 更新时间
		SaAddLinkPc   string    `db:"sa_add_link_pc"`  // 加入好友电脑链接
		SaAddLinkMb   string    `db:"sa_add_link_mb"`  // 加入好友手机链接
		SaStatus      int64     `db:"sa_status"`       // 账号是否可用，0不可用（需要审核），1可用，2满粉了，3出问题了
		SaOnline      int64     `db:"sa_online"`       // 在线？默认0，默认离线状态
		SaRemark      string    `db:"sa_remark"`       // 备注
		SaLuckyValue  int64     `db:"sa_lucky_value"`  // 幸运值，默认 0
		SaBindMobile  string    `db:"sa_bind_mobile"`
		SaBindEmail   string    `db:"sa_bind_email"`
		SaTypeId      int64     `db:"sa_type_id"` // 类型id
	}
)

func NewServiceAccountModel(conn sqlx.SqlConn, c cache.CacheConf) ServiceAccountModel {
	return &defaultServiceAccountModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`service_account`",
	}
}

func (m *defaultServiceAccountModel) Insert(data *ServiceAccount) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, serviceAccountRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.SaUserId, data.SaBrandId, data.SaAccountId, data.SaShowAccount, data.SaAccount, data.SaPassword, data.SaCreateTime, data.SaUpdateTime, data.SaAddLinkPc, data.SaAddLinkMb, data.SaStatus, data.SaOnline, data.SaRemark, data.SaLuckyValue, data.SaBindMobile, data.SaBindEmail, data.SaTypeId)

	return ret, err
}

func (m *defaultServiceAccountModel) FindOne(said int64) (*ServiceAccount, error) {
	workbenchServiceAccountSaidKey := fmt.Sprintf("%s%v", cacheWorkbenchServiceAccountSaidPrefix, said)
	var resp ServiceAccount
	err := m.QueryRow(&resp, workbenchServiceAccountSaidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `said` = ? limit 1", serviceAccountRows, m.table)
		return conn.QueryRow(v, query, said)
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

func (m *defaultServiceAccountModel) Update(data *ServiceAccount) error {
	workbenchServiceAccountSaidKey := fmt.Sprintf("%s%v", cacheWorkbenchServiceAccountSaidPrefix, data.Said)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `said` = ?", m.table, serviceAccountRowsWithPlaceHolder)
		return conn.Exec(query, data.SaUserId, data.SaBrandId, data.SaAccountId, data.SaShowAccount, data.SaAccount, data.SaPassword, data.SaCreateTime, data.SaUpdateTime, data.SaAddLinkPc, data.SaAddLinkMb, data.SaStatus, data.SaOnline, data.SaRemark, data.SaLuckyValue, data.SaBindMobile, data.SaBindEmail, data.SaTypeId, data.Said)
	}, workbenchServiceAccountSaidKey)
	return err
}

func (m *defaultServiceAccountModel) Delete(said int64) error {

	workbenchServiceAccountSaidKey := fmt.Sprintf("%s%v", cacheWorkbenchServiceAccountSaidPrefix, said)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `said` = ?", m.table)
		return conn.Exec(query, said)
	}, workbenchServiceAccountSaidKey)
	return err
}

func (m *defaultServiceAccountModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheWorkbenchServiceAccountSaidPrefix, primary)
}

func (m *defaultServiceAccountModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `said` = ? limit 1", serviceAccountRows, m.table)
	return conn.QueryRow(v, query, primary)
}
