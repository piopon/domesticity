package dataservice

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/piopon/domesticity/services/text-event-service/src/model"
	"github.com/piopon/domesticity/services/text-event-service/src/utils"
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
	typeInternal filterType = 0
	typeOption   filterType = 1
	typeFilter   filterType = 2
)

// mongoFilters is a map for defining available MongoDB filters
var mongoFilters = availableFilters{
	"limit":    {typeOption, "limit", limitQuery},
	"offset":   {typeOption, "offset", offsetQuery},
	"title":    {typeFilter, "title", regexQuery},
	"owner":    {typeFilter, "owner", exactQuery},
	"dayStart": {typeFilter, "date.start", dateQuery},
	"dayStop":  {typeFilter, "date.stop", dateQuery},
	"inMonth":  {typeFilter, "date.start", monthQuery},
	"category": {typeFilter, "category", exactQuery},
	"content":  {typeFilter, "content", regexQuery},
}

// memoryFilters is a map for defining available in memory DB filters
var memoryFilters = availableFilters{
	"internal": {typeInternal, "internal", internalQuery},
}

// NewFilters is a factory method for creating Filters structure
func NewFilters(config *utils.ConfigServer) *Filters {
	if config.TypeDB == "mongo" {
		return &Filters{mongoFilters}
	}
	return &Filters{memoryFilters}
}

// GetOptions is used to specify find request MongoDB options
func (filters Filters) GetOptions(queryParams url.Values) (*options.FindOptions, error) {
	if len(queryParams) == 0 {
		return nil, nil
	}
	queryOptions := options.FindOptions{}
	for key, value := range queryParams {
		if filter, ok := filters.available[key]; ok {
			if filter.Type == typeInternal {
				return nil, fmt.Errorf("Internal filters does not provide additional options")
			}
			if filter.Type != typeOption {
				continue
			}
			valueParsed, error := strconv.ParseInt(value[0], 10, 64)
			if error != nil {
				return nil, fmt.Errorf("Filter '"+key+"': cannot parse input value %s", value[0])
			}
			filter.Query.(func(*options.FindOptions, int64))(&queryOptions, valueParsed)
		} else {
			return nil, fmt.Errorf("Filter named '" + key + "' is not available")
		}
	}
	return &queryOptions, nil
}

// GetFilters is used to update bson interface to filter MongoDB results
func (filters Filters) GetFilters(queryParams url.Values) (interface{}, error) {
	if len(queryParams) == 0 {
		return bson.M{}, nil
	}
	queryFilter := []bson.M{}
	for key, value := range queryParams {
		if filter, ok := filters.available[key]; ok {
			if filter.Type == typeInternal {
				return filters.available[key].Query, nil
			}
			if filter.Type != typeFilter {
				continue
			}
			query := filter.Query.(func(string, []string) interface{})(filter.FieldDB, value)
			queryFilter = append(queryFilter, query.(bson.M))
		} else {
			return nil, fmt.Errorf("Filter named '" + key + "' is not available")
		}
	}
	if len(queryFilter) > 0 {
		return bson.M{"$and": queryFilter}, nil
	}
	return bson.M{}, nil
}

// GetAvailable returns list of available filter names
func (filters Filters) GetAvailable() []string {
	keys := make([]string, 0, len(filters.available))
	for k := range filters.available {
		keys = append(keys, k)
	}
	return keys
}

func dateQuery(dbField string, value []string) interface{} {
	day, _ := time.Parse("2006-02-01", value[0])
	minDayTime := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.UTC)
	maxDayTime := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 9999999, time.UTC)
	return bson.M{dbField: bson.M{"$gte": minDayTime, "$lte": maxDayTime}}
}

func monthQuery(dbField string, value []string) interface{} {
	day, _ := time.Parse("2006-02-01", value[0])
	daysInMonth := time.Date(day.Year(), day.Month() + 1, 0, 0, 0, 0, 0, time.UTC).Day()
	minDayTime := time.Date(day.Year(), day.Month(), 1, 0, 0, 0, 0, time.UTC)
	maxDayTime := time.Date(day.Year(), day.Month(), daysInMonth, 23, 59, 59, 9999999, time.UTC)
	return bson.M{dbField: bson.M{"$gte": minDayTime, "$lte": maxDayTime}}
}

func exactQuery(dbField string, value []string) interface{} {
	return bson.M{dbField: value[0]}
}

func regexQuery(dbField string, value []string) interface{} {
	return bson.M{dbField: primitive.Regex{Pattern: value[0], Options: ""}}
}

func limitQuery(dest *options.FindOptions, src int64) {
	dest.Limit = &src
}

func offsetQuery(dest *options.FindOptions, src int64) {
	dest.Skip = &src
}

func internalQuery(events model.Events, queryParams url.Values) (*model.Events, error) {
	return events.Filter(queryParams)
}
