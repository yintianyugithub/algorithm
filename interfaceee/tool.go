package interfaceee

import (
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

func str2Int64(i string) int64 {
	v, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		logx.Errorf("str2Int64 err: %v", err)
		return 0
	}

	return v
}
