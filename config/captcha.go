package config

type Captcha struct {
	KeyLong   int `mapstructure:"KeyLong" json:"KeyLong" yaml:"KeyLong"`       // 验证码长度
	ImgWidth  int `mapstructure:"ImgWidth" json:"ImgWidth" yaml:"ImgWidth"`    // 验证码宽度
	ImgHeight int `mapstructure:"ImgHeight" json:"ImgHeight" yaml:"ImgHeight"` // 验证码高度
}
