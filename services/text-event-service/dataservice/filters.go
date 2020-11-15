package dataservice

import (
	"github.com/piopon/domesticity/services/text-event-service/utils"
)

// Filters is a struct containing all settings for filtering DB data
type Filters struct {
	available availableFilters
}

// availableFilters is an alias for string and filterData map
type availableFilters map[string]filterData

// filter is an inner struct for defining all neccessary data
type filterData struct {
	Type    filterType
	FieldDB string
	Query   interface{}
}

// filterType is an alias defined for integer type
type filterType int

const (
	typeOption filterType = 0
	typeFilter filterType = 1
)

// mongoFilters is a map for defining available MongoDB filters
var mongoFilters = availableFilters{
	"limit":    {typeOption, "limit", nil},
	"offset":   {typeOption, "offset", nil},
	"title":    {typeFilter, "title", nil},
	"owner":    {typeFilter, "owner", nil},
	"dayStart": {typeFilter, "date.start", nil},
	"dayStop":  {typeFilter, "date.stop", nil},
	"category": {typeFilter, "category", nil},
	"content":  {typeFilter, "content", nil},
}

// memoryFilters is a map for defining available in memory DB filters
var memoryFilters = availableFilters{}

// NewFilters is a factory method for creating Filters structure
func NewFilters(config *utils.ConfigServer) *Filters {
	if config.TypeDB == "mongo" {
		return &Filters{mongoFilters}
	}
	return &Filters{memoryFilters}
}
