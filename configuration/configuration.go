package configuration

import (
	"fmt"
	"github.com/pkg/errors"
)

type Configuration struct {
	ConnectionStrings map[string]string `yaml:"connection_strings"`
	AppSettings map[string]interface{} `yaml:"app_settings"`
}

func (c *Configuration) GetConnectionString(name string) (string, error) {
	cs := c.ConnectionStrings[name]
	if cs == "" {
		return "", errors.New(fmt.Sprintf("Error resolving connection string with name %v\n", name))
	}

	return cs, nil
}

func (c *Configuration) GetAppSettingString(name string) (string, error) {
	setting := c.AppSettings[name]
	if setting == nil {
		return "", errors.New(fmt.Sprintf("App setting %v is not present\n", name))
	}

	if s, ok := setting.(string); !ok {
		return "", errors.New(fmt.Sprintf("Apps setting %v could not be converted to a string\n", name))
	} else {
		return s, nil
	}
}

func (c *Configuration) GetAppSettingInt(name string) (int, error) {
	setting := c.AppSettings[name]
	if setting == nil {
		return -1, errors.New(fmt.Sprintf("App setting %v is not present\n", name))
	}

	if i, ok := setting.(int); !ok {
		return -1, errors.New(fmt.Sprintf("App setting %v could not be converted to an int\n", name))
	} else {
		return i, nil
	}
}

func (c *Configuration) GetAppSettingInt32(name string) (int32, error) {
	setting := c.AppSettings[name]
	if setting == nil {
		return -1, errors.New(fmt.Sprintf("App setting %v is not present\n", name))
	}

	if i, ok := setting.(int32); !ok {
		return -1, errors.New(fmt.Sprintf("App setting %v could not be converted to an int32\n", name))
	} else {
		return i, nil
	}
}

func (c *Configuration) GetAppSettingInt64(name string) (int64, error) {
	setting := c.AppSettings[name]
	if setting == nil {
		return -1, errors.New(fmt.Sprintf("App setting %v is not present\n", name))
	}

	if i, ok := setting.(int64); !ok {
		return -1, errors.New(fmt.Sprintf("App setting %v could not be converted to an int64\n", name))
	} else {
		return i, nil
	}
}

func (c *Configuration) GetAppSettingBool(name string) (bool, error) {
	setting := c.AppSettings[name]
	if setting == nil {
		return false, errors.New(fmt.Sprintf("App setting %v is not present\n", name))
	}

	if b, ok := setting.(bool); !ok {
		return false, errors.New(fmt.Sprintf("App setting %v could not be converted to a bool\n", name))
	} else {
		return b, nil
	}
}