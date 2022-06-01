package database

import (
	"context"
	"douyin/controller"
	"fmt"
	"testing"
)

//TestGetRedisConnect Redis连接测试
func TestGetRedisConnect(t *testing.T) {
	rdb := GetRedisConnect()
	ctx := context.Background()
	fmt.Println("Redis连接状态" + rdb.String())

	// test the set op
	_ = rdb.Set(ctx, "tan", "1", 0)

	ret := rdb.Get(ctx, "tan")
	fmt.Println(ret.Val())
}

//TestGetDBConnect 数据库连接测试
func TestGetDBConnect(t *testing.T) {
	conn := GetDBConnect()
	var user controller.User
	conn.First(&user)
	fmt.Printf("%+v", user)
}
