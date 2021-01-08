package dataservice_test

import (
	"net/url"
	"testing"

	"github.com/piopon/domesticity/services/text-event-service/src/dataservice"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func TestNewFiltersShouldCreateMongoFilters(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	if filters == nil {
		t.Error("Mongo filters should be created but filters = nil")
	}
	filterNames := filters.GetAvailable()
	filtersNo := len(filterNames)
	if filtersNo != 8 {
		t.Errorf("Received incorrect filters number. Available filters %d", filtersNo)
	}
	expected := []string{"limit", "offset", "title", "owner", "dayStart", "dayStop", "category", "content"}
	for i := range filterNames {
		exists := isPresent(expected, filterNames[i])
		if !exists {
			t.Error("Filter names does not match")
		}
	}
}

func TestNewFiltersShouldCreateMemoryFilters(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "memory"}
	filters := dataservice.NewFilters(&config)
	if filters == nil {
		t.Error("Memory filters should be created but filters = nil")
	}
	filterNames := filters.GetAvailable()
	filtersNo := len(filterNames)
	if filtersNo != 1 {
		t.Errorf("Received incorrect filters number. Available filters %d", filtersNo)
	}
	expected := []string{"internal"}
	for i := range filterNames {
		exists := isPresent(expected, filterNames[i])
		if !exists {
			t.Error("Filter names does not match")
		}
	}
}

func TestGetOptionsShouldReturnNilWhenQueryEmpty(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	options, error := filters.GetOptions(url.Values{})
	if error != nil {
		t.Errorf("Get optons returned error on empty query: %s", error.Error())
	}
	if options != nil {
		t.Error("Get options returned non-empty option value")
	}
}

func TestGetOptionsShouldFailOnInvalidFilterName(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	options, error := filters.GetOptions(url.Values{"incorrect": []string{"filter"}})
	if error == nil {
		t.Error("Get optons should return error on incorrect filter name")
	}
	if options != nil {
		t.Error("Get options should be nil on incorrect filter name")
	}
}

func TestGetOptionsShouldReturnCorrectResultOnLimit(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	options, error := filters.GetOptions(url.Values{"limit": []string{"10"}})
	if error != nil {
		t.Errorf("Get optons returned error on empty query: %s", error.Error())
	}
	if options == nil {
		t.Error("Get options returned empty option value")
	}
	if *options.Limit != int64(10) {
		t.Errorf("Get options returned bad limit value %d", options.Limit)
	}
}

func TestGetOptionsShouldFailOnIncorrectLimitNumber(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	options, error := filters.GetOptions(url.Values{"limit": []string{"zly"}})
	if error == nil {
		t.Error("Get optons should return error on incorrect limit value")
	}
	if options != nil {
		t.Error("Get options should be nil on incorrect limit value")
	}
}

func TestGetOptionsShouldReturnCorrectResultOnOffset(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	options, error := filters.GetOptions(url.Values{"offset": []string{"10"}})
	if error != nil {
		t.Errorf("Get optons returned error on empty query: %s", error.Error())
	}
	if options == nil {
		t.Error("Get options returned empty option value")
	}
	if *options.Skip != int64(10) {
		t.Errorf("Get options returned bad limit value %d", options.Limit)
	}
}

func TestGetOptionsShouldSkipNonTypeOptionFilter(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	testQuery := url.Values{
		"owner":   []string{"Admin"},
		"dayStop": []string{"2020-03-03"},
	}
	options, error := filters.GetOptions(testQuery)
	if error != nil {
		t.Errorf("Get optons returned error on empty query: %s", error.Error())
	}
	if options == nil {
		t.Error("Get options returned empty option value")
	}
	if options.Skip != nil || options.Limit != nil {
		t.Errorf("Get options returned bad options: %v", options)
	}
}

func TestGetOptionsShouldReturnCorrectOptions(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	testQuery := url.Values{
		"owner":   []string{"Admin"},
		"limit":   []string{"2"},
		"dayStop": []string{"2020-03-03"},
		"offset":  []string{"5"},
	}
	options, error := filters.GetOptions(testQuery)
	if error != nil {
		t.Errorf("Get optons returned error on empty query: %s", error.Error())
	}
	if options == nil {
		t.Error("Get options returned empty option value")
	}
	if *options.Skip != int64(5) || *options.Limit != int64(2) {
		t.Errorf("Get options returned bad options: %v", options)
	}
}

func TestGetOptionsShouldFailOnInternalQuery(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "memory"}
	filters := dataservice.NewFilters(&config)
	testQuery := url.Values{"internal": []string{""}}
	options, error := filters.GetOptions(testQuery)
	if error == nil {
		t.Error("Get optons should return error on internal query")
	}
	if options != nil {
		t.Error("Get options should be nil on internal query")
	}
}

func TestGetFilterShouldReturnEmptyFilterIfQueryEmpty(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	result, error := filters.GetFilters(url.Values{})
	if error != nil {
		t.Errorf("Get filters returned error on empty query: %s", error.Error())
	}
	if result == nil {
		t.Error("Get filters returned nil bson")
	}
	if len(result.(bson.M)) != 0 {
		t.Error("Get filters returned non empty bson result")
	}
}

func TestGetFiltersShouldSkipNonTypeFilter(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	testQuery := url.Values{
		"limit":  []string{"2"},
		"offset": []string{"3"},
	}
	result, error := filters.GetFilters(testQuery)
	if error != nil {
		t.Errorf("Get filter returned error on options query: %s", error.Error())
	}
	if result == nil {
		t.Error("Get filter returned nil bson")
	}
	if len(result.(bson.M)) != 0 {
		t.Error("Get filters returned non empty bson result")
	}
}

func TestGetFiltersShouldReturnCorrectFilters(t *testing.T) {
	config := utils.ConfigServer{TypeDB: "mongo"}
	filters := dataservice.NewFilters(&config)
	testQuery := url.Values{
		"owner":   []string{"Admin"},
		"title":   []string{"test"},
		"dayStop": []string{"2020-03-03"},
		"offset":  []string{"5"},
	}
	result, error := filters.GetFilters(testQuery)
	if error != nil {
		t.Errorf("Get filter returned error on filters query: %s", error.Error())
	}
	if result == nil {
		t.Error("Get filter returned nil bson")
	}
	if len(result.(bson.M)) != 1 {
		t.Errorf("Get filters returned empty bson result: %v", result)
	}
}

func isPresent(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
