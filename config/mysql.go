package config

import "fmt"

type Mysql struct {
	Host         string `mapstructure:"Host" json:"Host" yaml:"Host"`                         // 服务器地址
	Port         int    `mapstructure:"Port" json:"Port" yaml:"Port"`                         // 端口
	Params       string `mapstructure:"Params" json:"Params" yaml:"Params"`                   // 高级配置
	Dbname       string `mapstructure:"Dbname" json:"Dbname" yaml:"Dbname"`                   // 数据库名
	Username     string `mapstructure:"Username" json:"Username" yaml:"Username"`             // 数据库用户名
	Password     string `mapstructure:"Password" json:"Password" yaml:"Password"`             // 数据库密码
	MaxIdleConns int    `mapstructure:"MaxIdleConns" json:"MaxIdleConns" yaml:"MaxIdleConns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"MaxOpenConns" json:"MaxOpenConns" yaml:"MaxOpenConns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"LogMode" json:"LogMode" yaml:"LogMode"`                // Gorm全局日志级别: silent|error|warn|info(默认)
	LogZap       bool   `mapstructure:"LogZap" json:"LogZap" yaml:"LogZap"`                   // 是否通过zap写入日志文件
}

func (m *Mysql) DataSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.Dbname, m.Params)
	// return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Params
}
