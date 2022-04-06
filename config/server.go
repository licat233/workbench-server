package config

const ProjectDir = "workbench-server"
const ProjectConfigFile = "deploy/server/config.yaml"

//ServerConfig 主服务配置
type ServerConfig struct {
	RootPath string `json:"RootPath" yaml:"RootPath"` //项目根目录
	Domain   string `json:"Domain" yaml:"Domain"`     //绑定域名
	Mode     string `json:"Mode" yaml:"Mode"`         //项目开发模式
}

//StaticConfig 静态资源服务配置
type StaticConfig struct {
	Domain string `json:"Domain" yaml:"Domain"` //绑定域名
	Root   struct {
		Path   string `json:"Path" yaml:"Path"` //文件仓库目录
		Upload struct {
			Path   string `json:"Path" yaml:"Path"` //上传文件目录
			Images struct {
				Path   string `json:"Path" yaml:"Path"` //上传图片目录
				Avatar struct {
					Path string `json:"Path" yaml:"Path"` //上传头像目录
				} `json:"Avatar" yaml:"Avatar"`
			} `json:"Images" yaml:"Images"`
			Collect struct {
				Path string `json:"Path" yaml:"Path"` //上传收藏夹目录
			} `json:"Collect" yaml:"Collect"`
		} `json:"Upload" yaml:"Upload"`
		Static struct {
			Path string `json:"Path" yaml:"Path"` //静态文件
			Js   struct {
				Path string `json:"Path" yaml:"Path"` // js文件
			} `json:"Js" yaml:"Js"`
			Css struct {
				Path string `json:"Path" yaml:"Path"` // js文件
			} `json:"Css" yaml:"Css"`
			Images struct {
				Path string `json:"Path" yaml:"Path"` // img文件
			} `json:"Images" yaml:"Images"`
			Svg struct {
				Path string `json:"Path" yaml:"Path"` // svg文件
			} `json:"Svg" yaml:"Svg"`
			Json struct {
				Path string `json:"Path" yaml:"Path"` // svg文件
			} `json:"Json" yaml:"Json"`
			Txt struct {
				Path string `json:"Path" yaml:"Path"` // txt文件
			} `json:"Txt" yaml:"Txt"`
		} `json:"Static" yaml:"Static"`
	} `json:"Root" yaml:"Root"`
}
