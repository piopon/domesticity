package dataservice

// Filters is a struct containing all settings for filtering DB data
type Filters struct {
	available availableFilters
}

// availableFilters is an alias for string and filterData map
type availableFilters map[string]string
