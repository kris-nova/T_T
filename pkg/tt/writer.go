package tt

// MessageWriter allows writing IRC messages. The MessageWriter
// passed to the handler is contextual based on the incoming message,
// and will write to the room on room messages or to the user on
// direct messages.
type MessageWriter interface {
	Write(msg string)
}
