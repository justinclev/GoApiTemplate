package model

import "go.mongodb.org/mongo-driver/mongo"

type MongoDatabase struct {
	Uri string
	Name string
	Collections map[string]*mongo.Collection
}

type Configuration struct {
	Database MongoDatabase
	Settings map[string]string
}