package mux

import (
	"context"
	"sync"

	"github.com/kris-nova/T_T/pkg/tt"
)

// Mux is a router for action-based plugins
type Mux interface {
	tt.Handler

	Register(p tt.Plugin)
}

// New creates a new Mux router
func New() Mux {
	return &mux{
		plugins: make(map[int64]tt.Plugin),
	}
}

type mux struct {
	l       sync.RWMutex
	rules   []tt.Rule
	plugins map[int64]tt.Plugin
}

func (mx *mux) Register(p tt.Plugin) {
	rl := p.Rules()

	mx.l.Lock()
	defer mx.l.Unlock()

	for _, r := range rl {
		mx.rules = append(mx.rules, r)
		idx := len(mx.rules)
		mx.plugins[int64(idx-1)] = p
	}
}

func (mx *mux) ServeIRC(ctx context.Context, m *tt.Message, w tt.MessageWriter) {
	mx.l.RLock()
	defer mx.l.RUnlock()

	var plugins = make(map[string]tt.Plugin)

	// rules are currently executed sequentially but this is super optional.
	// We can invoke plugins via channels, via a message queue, etc.

	for idx, r := range mx.rules {
		if ok := r.TestRule(m); ok {
			p := mx.plugins[int64(idx)]
			pluginID := p.ID()
			if _, ok := plugins[pluginID]; !ok {
				plugins[pluginID] = p
			}
		}
	}

	for _, p := range plugins {
		h := p.Handler()
		h.ServeIRC(ctx, m, w)
	}
}
