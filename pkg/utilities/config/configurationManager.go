package config

import (
	"GoApiExample/pkg/utilities/config/model"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func LoadConfiguration() (*model.Configuration, error) {
	databases, err := LoadDatabases()
	if err != nil {
		return nil, err
	}

	settings, err := LoadSettings()
	if err != nil {
		return nil, err
	}

	return &model.Configuration {
		Database: databases,
		Settings: settings,
	}, nil
}

func LoadDatabases()(model.MongoDatabase, error) {
	collections := make(map[string]*mongo.Collection)

	collections[os.Getenv("MongoDataBaseCollection_Admin")] = nil
	collections[os.Getenv("MongoDataBaseCollection_User")] = nil
	

	databases := model.MongoDatabase{
		Uri: os.Getenv("MongoDataBaseUri_Admin"),
		Database: os.Getenv("MongoDataBaseDB_Admin"),
		Collections: collections,
	}

	return databases, nil
}

func LoadSettings()(map[string]string, error) {

	settings := make(map[string]string)

	return settings, nil
}