package config

import (
	"log"

	"github.com/alimgiray/guido/database"
)

var defaultAppName = "Guido"
var defaultUserType = "user"

type ConfigurationManager struct {
	db             *database.Database
	configurations map[string]string
}

func NewConfigurationManager(DB *database.Database) *ConfigurationManager {
	return &ConfigurationManager{
		db:             DB,
		configurations: load(DB),
	}
}

func load(DB *database.Database) map[string]string {
	configs := make(map[string]string)
	rows, err := DB.Connection.Query("SELECT name, value FROM config")
	if err != nil {
		log.Println("Couldn't load configurations", err.Error())
		return configs
	}
	for true {
		if !rows.Next() {
			break
		}
		var name string
		var value string
		rows.Scan(&name, &value)
		configs[name] = value
	}
	return configs
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

func (c *ConfigurationManager) GetNewUserType() string {
	if val, ok := c.configurations["userType"]; ok {
		return val
	}
	return defaultUserType
}
