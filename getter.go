package emailtemplate

import (
	"sync"
)

// Key allows callers to load, set and get templates at a predictable
// key in a templates map.
type Key string

// Getter is a map of template keys to Template that allows looking up a
// given template by key.
type Getter struct {
	m  map[Key]*Template
	mu sync.RWMutex // mu protects m.
}

// newGetter makes a new initialized Getter.
func newGetter() *Getter {
	return &Getter{m: make(map[Key]*Template)}
}

// Get returns the Template at key k, if it exists.
func (g *Getter) Get(k Key) (*Template, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	t, ok := g.m[k]
	return t, ok
}

// set puts the Template in the map at key k.
func (g *Getter) set(k Key, t *Template) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.m[k] = t
}

// len returns the number of templates in the map.
func (g *Getter) len() int {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return len(g.m)
}
