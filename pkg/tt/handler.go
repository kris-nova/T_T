package tt

import "context"

// Handler is the interface for handling IRC messages
type Handler interface {
	ServeIRC(ctx context.Context, msg *Message, w MessageWriter)
}

// HandlerFunc is the functional alias for the Handler interface
type HandlerFunc func(ctx context.Context, msg *Message, w MessageWriter)

// ServeIRC handles incoming IRC messages
func (h HandlerFunc) ServeIRC(ctx context.Context, msg *Message, w MessageWriter) {
	h(ctx, msg, w)
}
