package data

// Event defines the structure for an API event
type Event struct {
	ID       int
	Title    string
	Owner    string
	Category string
	Content  string
}

var eventList = []*Event{
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
