package main

import (
	"context"
	"interface/interfaceee"
)

func main() {
	ctx := context.TODO()
	uid := int64(2080)

	interfaceee.UserLogin(ctx, uid)
	interfaceee.UserLogin(ctx, uid)
	interfaceee.UserLogin(ctx, uid)
	interfaceee.UserLogin(ctx, uid)
	interfaceee.UserLogin(ctx, uid)
	interfaceee.UserLogin(ctx, uid)
	interfaceee.CalculateLoginTime(ctx, uid)
}
