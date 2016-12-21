package tt

// A Plugin is an extension for routing messages
type Plugin interface {
	ID() string

	Rules() []Rule

	Handler() Handler
}
