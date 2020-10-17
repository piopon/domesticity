package data

// Event defines the structure for an API event
type Event struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Owner    string `json:"owner"`
	Category string `json:"category"`
	Content  string `json:"content"`
}

// Events is a type definition for slice of Event pointers
type Events []*Event

// GetEvents returns all events stored in DB
func GetEvents() Events {
	return eventList
}

var eventList = Events{
	&Event{
		ID:       1,
		Title:    "This is my first event",
		Owner:    "Admin",
		Category: "Notes",
		Content:  "Test event number 1",
	},
	&Event{
		ID:       2,
		Title:    "2nd event",
		Owner:    "Admin",
		Category: "Stuff",
		Content:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	},
}
