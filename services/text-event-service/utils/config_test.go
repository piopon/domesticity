package utils_test

import (
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/utils"
)

func TestNewConfig(t *testing.T) {
	config := utils.NewConfig("../scripts")
	expected := "text-event-service"
	if config.Name != expected {
		t.Errorf("expected '%s', got '%s'", expected, config.Name)
	}
}
