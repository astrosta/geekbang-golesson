/*
	解题思路：当遇到sql.ErrNoRows的时候，代表未查询到数据。在数据库中未查询到数据不应当作一个错误处理，应将错误降级，返回查询结果为空
*/

package main

import (
	"database/sql"

	"github.com/pkg/errors"
)

type QueryResultMock struct {
}

func queryMock() ([]QueryResultMock, error) {
	err := queryRowMock()
	if err != nil {
		//若数据库未查询到结果，错误降级，返回查询结果为nil
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "query error")
	}

	//模拟处理查询结果
	result := make([]QueryResultMock, 1)

	return result, nil
}

//模拟查询数据库
func queryRowMock() error {
	return sql.ErrNoRows
}
