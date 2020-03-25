package config

import (
	"GoApiExample/pkg/utilities/config/model"
	"os"
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
		MongoCollections: databases,
		Settings: settings,
	}, nil
}

func LoadDatabases()(map[string]model.MongoDatabase, error) {
	databases := make(map[string]model.MongoDatabase)

	databases[os.Getenv("MongoDataBaseCollection_Admin")] = model.MongoDatabase{
		Uri: os.GetEnv("MongoDataBaseUri_Admin"),
		Database: os.GetEnv("MongoDataBaseDB_Admin"),
		CollectionName: os.GetEnv("MongoDataBaseCollection_Admin"),
	}
	databases[os.Getenv("MongoDataBaseCollection_User")] = model.MongoDatabase{
		Uri: os.GetEnv("MongoDataBaseUri_User"),
		Database: os.GetEnv("MongoDataBaseDB_User"),
		CollectionName: os.GetEnv("MongoDataBaseCollection_User"),
	}

	return databases, nil
}

func LoadSettings()(map[string]string, error) {

	settings := make(map[string]string)

	return settings, nil
}