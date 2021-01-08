package utils_test

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/src/utils"
)

func TestConvertsCorrectStructToJson(t *testing.T) {
	test := jsonStruct{
		ID:    123,
		Name:  "test-struct-name",
		Price: 99.99,
	}
	var buffer bytes.Buffer
	error := utils.ToJSON(test, &buffer)
	if error != nil {
		t.Errorf("Could not convert test structure to JSON")
	}
	expected := createBufferJsonString(test.ID, test.Name, test.Price)
	if buffer.String() != expected {
		t.Errorf("Test structure was incorrectly parsed to JSON")
		t.Errorf("- expected: %s", expected)
		t.Errorf("- got:      %s", buffer.String())
	}
}

func TestGetsCorrectStructFromJson(t *testing.T) {
	test := jsonStruct{}
	expectedID := 999
	expectedName := "my_name"
	expectedPrice := 123.456
	var buffer bytes.Buffer
	buffer.WriteString(createBufferJsonString(expectedID, expectedName, expectedPrice))
	error := utils.FromJSON(&test, &buffer)
	if error != nil {
		t.Errorf("Could not receive test structure from JSON")
	}
	if test.ID != expectedID {
		t.Errorf("Test structure ID was incorrectly parsed from JSON")
	}
	if test.Name != expectedName {
		t.Errorf("Test structure NAME was incorrectly parsed from JSON")
	}
	if test.Price != expectedPrice {
		t.Errorf("Test structure PRICE was incorrectly parsed from JSON")
	}
}

type jsonStruct struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func createBufferJsonString(id int, name string, price float64) string {
	idParsed := strconv.Itoa(id)
	priceParsed := fmt.Sprintf("%g", price)
	return "{\"id\":" + idParsed + ",\"name\":\"" + name + "\",\"price\":" + priceParsed + "}\n"
}
