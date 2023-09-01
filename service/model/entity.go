package model

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
