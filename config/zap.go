package config

type Zap struct {
	Level         string `mapstructure:"Level" json:"Level" yaml:"Level"`                         // 级别
	Format        string `mapstructure:"Format" json:"Format" yaml:"Format"`                      // 输出格式：json | console(默认)
	Prefix        string `mapstructure:"Prefix" json:"Prefix" yaml:"Prefix"`                      // 日志前缀
	Director      string `mapstructure:"Director" json:"Director"  yaml:"Director"`               // 日志文件夹
	ShowLine      bool   `mapstructure:"ShowLine" json:"ShowLine" yaml:"ShowLine"`                // 显示行
	EncodeLevel   string `mapstructure:"EncodeLevel" json:"EncodeLevel" yaml:"EncodeLevel"`       // 编码级
	StacktraceKey string `mapstructure:"StacktraceKey" json:"StacktraceKey" yaml:"StacktraceKey"` // 栈名
	LogInConsole  bool   `mapstructure:"LogInConsole" json:"LogInConsole" yaml:"LogInConsole"`    // 输出控制台
}
