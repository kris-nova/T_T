package main

import (
	"context"
	"strings"

	"github.com/kris-nova/T_T/pkg/tt"
)

type containsPlugin struct {
}

func (c *containsPlugin) ID() string {
	return "WhoIsKris"
}

func (c *containsPlugin) Handler() tt.Handler {
	return tt.HandlerFunc(func(_ context.Context, m *tt.Message, w tt.MessageWriter) {
		resp := "Kris is awesome! She runs the server around here!"
		w.Write(resp)
	})
}

func (c *containsPlugin) Rules() []tt.Rule {
	return []tt.Rule{
		&containsRule{Value: "!Kris"},
		&containsRule{Value: "!@Kris"},
	}
}

type containsRule struct {
	Value string
}

func (r *containsRule) TestRule(m *tt.Message) bool {
	return strings.Contains(m.Body, r.Value)
}
