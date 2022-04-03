package main

import (
	"database/sql"
	"errors"
	"fmt"
)

/*
我觉得需要Wrap这个error抛给上层.
因为我们调用的第三方库的SQL查询,在进行查询时,它抛出的error不足以让我们定位和排查问题,我们应该wrap尽量以脱敏的msg来帮助定位代码有问题的区域
*/

func sqlErr() error {
	var err error
	err = sql.ErrNoRows
	if err != nil {
		return fmt.Errorf("failed to sql querry : package Name with lineNum and some maybe Not sensitive msg, rootError:%w", err)
	}
	return nil
}
func main() {
	err := sqlErr()
	if errors.Is(err, sql.ErrNoRows) {
		//solution
	}
}
