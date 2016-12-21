package tt

// A Room is an IRC room
type Room struct {
	Name string // Name of the room
}

// Users gets the list of usernames for a room
func (r *Room) Users() []string {
	panic("Not implemented")
}

// Topic gets the topic for the room
func (r *Room) Topic() *Topic {
	return &Topic{
		room: r,
	}
}
