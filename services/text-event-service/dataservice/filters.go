package dataservice

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
