package config

import (
	"github.com/alimgiray/guido/database"
)

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
