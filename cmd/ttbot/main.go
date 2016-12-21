// Package main provides the ttbot binary implementation
package main

import (
	"context"
	"fmt"

	"github.com/kris-nova/T_T/pkg/tt"
	"github.com/kris-nova/T_T/pkg/tt/mux"
)

type exampleMessageWriter struct {
	items []string
}

func (w *exampleMessageWriter) Write(b string) {
	w.items = append(w.items, b)
}

func main() {

	ctx := context.Background()
	w := exampleMessageWriter{}
	contains := containsPlugin{}

	m := mux.New()
	m.Register(&contains)

	m.ServeIRC(ctx, testMessage("hello world"), &w)
	m.ServeIRC(ctx, testMessage("world !Kris"), &w)
	m.ServeIRC(ctx, testMessage("world !@Kris"), &w)

	fmt.Printf("%v\n", w.items)
}

func testMessage(msg string) *tt.Message {
	return &tt.Message{
		Body: msg,
		Room: nil,
		User: nil,
	}
}
