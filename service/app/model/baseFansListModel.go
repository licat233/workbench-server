package model

import (
	"fmt"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
)

type baseFansListModel interface {
	FindAll(Current, PageSize int64) ([]*FansList, error)
	FindCount() (int64, error)
}

func (m *defaultFansListModel) FindAll(Current, PageSize int64) ([]*FansList, error) {
	resp := make([]*FansList, 0)
	query := fmt.Sprintf("select %s from %s limit %v,%v", fansListRows, m.table, (Current-1)*PageSize, PageSize)
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

func (m *defaultFansListModel) FindCount() (int64, error) {
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
