package tt

// A User is an IRC user
type User struct {
	Username string // Name of user
	SourceIP string // IP address of the user
}

// Authenticated gets whether the user is authenticated with nickserv or the bot
func (u *User) Authenticated() bool {
	panic("Not Implemented")
}
