package tt

// A Message is an IRC message
type Message struct {
	Body string // The body of the message
	Room *Room  // The room the message came from, if there is one
	User *User  // The user the message came from
}
