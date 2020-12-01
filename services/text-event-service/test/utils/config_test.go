package utils_test

import (
	"os"
	"testing"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/src/utils"
)

func TestNewConfigShouldReadExistingScript(t *testing.T) {
	config := getProductionConfig()
	_, err := os.Stat("settings.yaml")
	scriptPresent := !os.IsNotExist(err)
	if scriptPresent {
		t.Errorf("Created default script instead of reading from exsiting one")
	}
	expectedName := "text-event-service"
	if config.Name != expectedName {
		t.Errorf("Expected '%s', got '%s'", expectedName, config.Name)
	}
}

func TestNewConfigShouldCreateMissingScript(t *testing.T) {
	getDefaultConfig()
	scriptInfo, err := os.Stat("settings.yaml")
	scriptPresent := !os.IsNotExist(err)
	if !scriptPresent {
		t.Errorf("Could not find default script")
		return
	}
	if scriptInfo.ModTime().YearDay() != time.Now().YearDay() {
		t.Errorf("Created script is not currently created")
	}
	cleanDefaultConfig()
}

func TestDefaultConfigShouldCreateCorrectServiceConfig(t *testing.T) {
	config := getDefaultConfig()
	expectedName := "text-event-service"
	if config.Name != expectedName {
		t.Errorf("Expected '%s', got '%s'", expectedName, config.Name)
	}
	expectedVerbose := false
	if config.Verbose != expectedVerbose {
		t.Errorf("Expected '%v', got '%v'", expectedVerbose, config.Verbose)
	}
	cleanDefaultConfig()
}

func TestDefaultConfigShouldCreateCorrectServerConfig(t *testing.T) {
	config := getDefaultConfig()
	expectedIP := ""
	if config.Server.IP != expectedIP {
		t.Errorf("Expected '%s', got '%s'", expectedIP, config.Server.IP)
	}
	expectedPort := "9999"
	if config.Server.Port != expectedPort {
		t.Errorf("Expected '%s', got '%s'", expectedPort, config.Server.Port)
	}
	expectedDbType := "mongo"
	if config.Server.TypeDB != expectedDbType {
		t.Errorf("Expected '%s', got '%s'", expectedDbType, config.Server.TypeDB)
	}
	expectedTimeouts := utils.ConfigServerTimeout{Idle: 600, Read: 10, Write: 10}
	if config.Server.Timeout != expectedTimeouts {
		t.Errorf("Expected '%v', got '%v'", expectedTimeouts, config.Server.Timeout)
	}
	cleanDefaultConfig()
}

func TestDefaultConfigShouldCreateCorrectMongoConfig(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log("Working dir: " + wd)

	config := getDefaultConfig()

	expectedScheme := "mongodb://"
	if config.MongoDB.Scheme != expectedScheme {
		t.Errorf("Expected '%s', got '%s'", expectedScheme, config.MongoDB.Scheme)
	}
	expectedIP := "127.0.0.1"
	if config.MongoDB.IP != expectedIP {
		t.Errorf("Expected '%s', got '%s'", expectedIP, config.MongoDB.IP)
	}
	expectedPort := "27017"
	if config.MongoDB.Port != expectedPort {
		t.Errorf("Expected '%s', got '%s'", expectedPort, config.MongoDB.Port)
	}
	expectedDB := utils.ConfigMongoData{"event-service", "events"}
	if config.MongoDB.Database != expectedDB {
		t.Errorf("Expected '%v', got '%v'", expectedDB, config.MongoDB.Database)
	}
	expectedTimeouts := utils.ConfigMongoTimeout{10, 10, 5, 5, 3}
	if config.MongoDB.Timeout != expectedTimeouts {
		t.Errorf("Expected '%v', got '%v'", expectedTimeouts, config.MongoDB.Timeout)
	}
	cleanDefaultConfig()
}

func getProductionConfig() *utils.Config {
	return utils.NewConfig("../../resources")
}

func getDefaultConfig() *utils.Config {
	return utils.NewConfig("")
}

func cleanDefaultConfig() {
	os.Remove("settings.yaml")
}
