package middleware

import (
	"context"
	"douyin/controller"
	"douyin/global"
	"douyin/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

/**
  用户认证拦截器中间件。主要做网关使用，拦截未登录请求。无需放在每个接口进行验证
*/

// AuthHandler 用户鉴权过滤器中间件
func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if len(token) == 0 {
			//不存在token
			c.JSON(http.StatusOK, controller.Response{
				StatusCode: 402,
				StatusMsg:  controller.CodeEmptyAuth.Msg(),
			})
			fmt.Println("用户未登陆被拦截！")
			c.Abort()
		} else {
			exists, _ := global.RedisEngine.Exists(context.Background(), global.TokenPrefix+token).Result()
			//验证成功
			if exists > 0 {
				global.RedisEngine.Expire(context.Background(), token, 24*time.Hour)
				userJson, _ := global.RedisEngine.Get(context.Background(), global.TokenPrefix+token).Result()
				//将用户信息加入上下文信息中,注意是value是user指针
				_, exists := c.Get("auth")
				if !exists {
					authUser := models.User{}
					if err := json.Unmarshal([]byte(userJson), &authUser); err != nil {
						global.Logger.Error("unmarshal wrong", zap.Error(err))
						c.Abort()
					}
					c.Set("auth", authUser)
				}
				fmt.Println("用户请求已放行！")
				c.Next()
			} else {
				//token过期
				c.JSON(http.StatusOK, controller.Response{
					StatusCode: 402,
					StatusMsg:  controller.CodeMap[controller.CodeInvalidAuth],
				})
				fmt.Println("用户token失效被拦截！")
				c.Abort()
			}
		}
	}
}
