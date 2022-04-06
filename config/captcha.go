package config

type Captcha struct {
	KeyLong   int `mapstructure:"KeyLong" json:"keyLong" yaml:"KeyLong"`       // 验证码长度
	ImgWidth  int `mapstructure:"ImgWidth" json:"imgWidth" yaml:"ImgWidth"`    // 验证码宽度
	ImgHeight int `mapstructure:"ImgHeight" json:"imgHeight" yaml:"ImgHeight"` // 验证码高度
}
