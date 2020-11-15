package dataservice

import (
	"time"

	"github.com/piopon/domesticity/services/text-event-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	"limit":    {typeOption, "limit", limitQuery},
	"offset":   {typeOption, "offset", offsetQuery},
	"title":    {typeFilter, "title", regexQuery},
	"owner":    {typeFilter, "owner", exactQuery},
	"dayStart": {typeFilter, "date.start", dateQuery},
	"dayStop":  {typeFilter, "date.stop", dateQuery},
	"category": {typeFilter, "category", exactQuery},
	"content":  {typeFilter, "content", regexQuery},
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

func dateQuery(dbField string, value []string) interface{} {
	day, _ := time.Parse("2006-02-01", value[0])
	minDayTime := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.UTC)
	maxDayTime := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 9999999, time.UTC)
	return bson.M{dbField: bson.M{"$gte": minDayTime, "$lte": maxDayTime}}
}

func exactQuery(dbField string, value []string) interface{} {
	return bson.M{dbField: value[0]}
}

func regexQuery(dbField string, value []string) interface{} {
	return bson.M{dbField: primitive.Regex{Pattern: value[0], Options: ""}}
}

func limitQuery(dest *options.FindOptions, src int64) {
	return
}

func offsetQuery(dest *options.FindOptions, src int64) {
	return
}
