package model

type MongoDatabase struct {
	Uri string
	Database string
	CollectionName string
}

type Configuration struct {
	MongoCollections map[string]MongoDatabase
	Settings map[string]string
}