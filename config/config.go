package config

import (
	"github.com/alimgiray/guido/database"
)

var defaultAppName = "Guido"

type ConfigurationManager struct {
	db             *database.Database
	configurations map[string]string
}

func NewConfigurationManager(DB *database.Database) *ConfigurationManager {
	return &ConfigurationManager{
		db: DB,
	}
}

func (c *ConfigurationManager) GetMeta(description, keywords string) *Meta {
	return &Meta{
		Description: description,
		Keywords:    keywords,
	}
}

func (c *ConfigurationManager) GetHeader(username string, loggedIn bool) *Header {
	header := &Header{
		AppName:    c.GetAppName(),
		IsLoggedIn: loggedIn,
		Username:   username,
	}
	return header
}

func (c *ConfigurationManager) GetAppName() string {
	if val, ok := c.configurations["appName"]; ok {
		return val
	}
	return defaultAppName
}
