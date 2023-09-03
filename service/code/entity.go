package code

// Config dest/model.yml 获取配置信息
type Config struct {
	Db     *DbConfig `yaml:"db"`
	Target *Target   `yaml:"target"`
}

type DbConfig struct {
	Url      string `yaml:"url"`
	DataBase string `yaml:"dataBase"`
	Tables   string `yaml:"tables"`
}

type Target struct {
	Lang    string      `yaml:"lang"`
	Package string      `yaml:"package"`
	Java    *JavaConfig `yaml:"java"`
	Go      *GoConfig   `yaml:"go"`
}

type JavaConfig struct {
}

type GoConfig struct {
}

// ClientConfig dest/client.yml 获取配置信息
type ClientConfig struct {
	Schema string        `yaml:"schema"`
	Source *ClientSource `yaml:"source"`
}

// ClientSource 资源获取方式
type ClientSource struct {
	Type stirng `yaml:"type"`
	Url  string `yaml:"url"`
}
