package global

import (
	"crypto/md5"
	"encoding/hex"

	"go.uber.org/zap"
)

var (
	Secrete = string("douyinxiangmu") //用于MD5加盐
	Logger  *zap.Logger               //全局日志
)

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(Secrete))
	return string(hex.EncodeToString(m.Sum([]byte(str))))
}
