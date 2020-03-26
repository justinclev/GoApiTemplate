package config

import (
	"os"
	"testing"
)

func runBefore() {

	//Running with debug mode set to true, will cause unit test to load from env. 
	//Useful for testing a production issue that might be related to a missing config value
	if os.Getenv("APIDebugMode") != "true" {
		os.Setenv("MongoDataBaseCollection_Admin", "Collection")
		os.Setenv("MongoDataBaseUri_Admin", "Uri")
		os.Setenv("MongoDataBaseDB_Admin", "Db")

		os.Setenv("MongoDataBaseCollection_User", "Collection1")
		os.Setenv("MongoDataBaseUri_User", "Uri1")
		os.Setenv("MongoDataBaseDB_User", "Db1")
	}
}

func Test_LoadConfiguration(t *testing.T) {
	runBefore()

	config, err := LoadConfiguration()
	if err != nil {
		t.Error(err)
	}

	//Change to actual database details
	if config.Database.Uri != "Uri" &&
	config.Database.Database != "Db" {
		t.Error("Error loading config Values")
	}
}