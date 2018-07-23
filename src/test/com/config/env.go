package config

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/haozzzzzzzz/go-rapid-development/utils/yaml"
	"github.com/sirupsen/logrus"
)

type EnvFormat struct {
	Debug bool   `json:"debug" yaml:"debug"`
	Stage string `json:"stage" yaml:"stage" validate:"required"`
}

func (m *EnvFormat) WithStagePrefix(strVal string) string {
	return fmt.Sprintf("%s_%s", m.Stage, strVal)
}

var EnvConfig EnvFormat

func init() {
	var err error
	err = yaml.ReadYamlFromFile("./config/env.yaml", &EnvConfig)
	if nil != err {
		logrus.Errorf("read env config file failed. %s.", err)
		return
	}

	err = validator.New().Struct(&EnvConfig)
	if nil != err {
		logrus.Fatalf("validate env config failed. %s", err)
		return
	}

}
