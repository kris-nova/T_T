package tt

// A Rule is a piece of for testing a message for a specific
// rule
type Rule interface {
	TestRule(msg *Message) bool
}
