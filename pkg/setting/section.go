package setting

import "time"

type ServerSettingS struct {
	RunMode      string        `mapstructure:"runMode"`
	HttpPort     int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout time.Duration `mapstructure:"WriteTimeout"`
}

type RedisSettingS struct {
	Url      string
	Password string
	Database int
}

type MysqlSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
type LoggerSettingS struct {
	LogLevel      string `mapstructure:"level"`
	LogSavePath   string `mapstructure:"log_path"`
	LogFileName   string `mapstructure:"filename"`
	MaxPageSize   int    `mapstructure:"max_size"`
	MaxPageAge    int    `mapstructure:"max_age"`
	LogMaxBackups int    `mapstructure:"max_backups"`
}

type OSSettingS struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyId     string `mapstructure:"accessKeyId"`
	AccessKeySecret string `mapstructure:"accessKeySecret"`
	BucketName      string `mapstructure:"bucketName"`
	TargetPath      string `mapstructure:"targetPath"`
	TargetURL       string `mapstructure:"targetUrl"`
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
