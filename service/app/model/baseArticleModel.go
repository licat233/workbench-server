package model

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type baseArticleModel interface {
	FindAll(Current, PageSize int64) ([]*Article, error)
	FindCount() (int64, error)
}

func (m *defaultArticleModel) FindAll(Current, PageSize int64) ([]*Article, error) {
	resp := make([]*Article, 0)
	query := fmt.Sprintf("select %s from %s limit %v,%v", articleRows, m.table, (Current-1)*PageSize, PageSize)
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

func (m *defaultArticleModel) FindCount() (int64, error) {
	var count int64
	query := fmt.Sprintf("SELECT count(*) FROM %s", m.table)
	err := m.QueryRowNoCache(&count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
