package config

type CORS struct {
	Mode      string          `mapstructure:"Mode" json:"Mode" yaml:"Mode"`
	Whitelist []CORSWhitelist `mapstructure:"Whitelist" json:"Whitelist" yaml:"Whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"AllowOrigin" json:"AllowOrigin" yaml:"AllowOrigin"`
	AllowMethods     string `mapstructure:"AllowMethods" json:"AllowMethods" yaml:"AllowMethods"`
	AllowHeaders     string `mapstructure:"AllowHeaders" json:"AllowHeaders" yaml:"AllowHeaders"`
	ExposeHeaders    string `mapstructure:"ExposeHeaders" json:"ExposeHeaders" yaml:"ExposeHeaders"`
	AllowCredentials bool   `mapstructure:"AllowCredentials" json:"AllowCredentials" yaml:"AllowCredentials"`
}
