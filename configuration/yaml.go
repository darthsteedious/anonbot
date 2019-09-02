package configuration

import (
	"bufio"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadYamlConfig(path string) (*Configuration, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(errors.WithStack(err), "Error opening yaml configuration file: %v", path)
	}

	decoder := yaml.NewDecoder(bufio.NewReader(file))

	config := &Configuration{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, errors.Wrapf(errors.WithStack(err), "Error decoding yaml configuration file: %v", path)
	}

	if err := file.Close(); err != nil {
		return nil, errors.Wrapf(errors.WithStack(err), "Error closing yaml file: %v", path)
	}

	return config, nil
}