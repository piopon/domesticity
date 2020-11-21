package utils_test

import (
	"os"
	"testing"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/utils"
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
	}
	if scriptInfo.ModTime().YearDay() != time.Now().YearDay() {
		t.Errorf("Created script is not currently created")
	}
	cleanDefaultConfig()
}

func getProductionConfig() *utils.Config {
	return utils.NewConfig("../scripts")
}

func getDefaultConfig() *utils.Config {
	return utils.NewConfig("")
}

func cleanDefaultConfig() {
	os.Remove("settings.yaml")
}
