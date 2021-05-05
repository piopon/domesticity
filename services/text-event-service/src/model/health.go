package model

// Health is a health check object returned by a server
type Health struct {
	Status string `json:"status"`
}

// ToString is a metod to return health result in a string variable
func (health Health) ToString() string {
	return health.Status
}
