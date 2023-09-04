package config

import (
	"fmt"
	"github.com/LiveAlone/go-util/util"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var SupportConfigFileType = []string{"yml", "yaml"}

type Loader struct{}

func NewConfigLoader() *Loader {
	return &Loader{}
}

func (l *Loader) LoadConfigToEntity(path string, entity any) error {
	paths := strings.Split(path, ".")
	if !supportFile(paths[len(paths)-1]) {
		return fmt.Errorf("not support file type %s", path)
	}

	confContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(confContent, entity)
	if err != nil {
		return err
	}
	return nil
}

// LoadConfigFromUrlToEntity 从url加载配置信息到实体
func (l *Loader) LoadConfigFromUrlToEntity(url string, entity any) (err error) {
	content, err := util.Get(url, nil)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, entity)
	if err != nil {
		return err
	}
	return nil
}

func supportFile(fileType string) bool {
	for _, s := range SupportConfigFileType {
		if s == fileType {
			return true
		}
	}
	return false
}
