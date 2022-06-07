package global

import "douyin/pkg/setting"

var (
	ServerSetting *setting.ServerSettingS //存放server全局数据
	RedisSetting  *setting.RedisSettingS  //存放redis全局数据
	MysqlSetting  *setting.MysqlSettingS  //存放mysql全局数据
	LoggerSetting *setting.LoggerSettingS //存放log的全局数据
	OSSetting     *setting.OSSettingS
)
