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
	m            map[Key]*Template
	sync.RWMutex // RWMutex protects m.
}

// newGetter makes a new initialized Getter.
func newGetter() *Getter {
	return &Getter{m: make(map[Key]*Template)}
}

// Get returns the Template at key k, if it exists.
func (ts *Getter) Get(k Key) (*Template, bool) {
	ts.RLock()
	t, ok := ts.m[k]
	ts.RUnlock()
	return t, ok
}

// set puts the Template in the map at key k.
func (ts *Getter) set(k Key, t *Template) {
	ts.Lock()
	ts.m[k] = t
	ts.Unlock()
}

// len returns the number of templates in the map.
func (ts *Getter) len() int {
	if ts == nil {
		return 0
	}
	ts.RLock()
	l := len(ts.m)
	ts.RUnlock()
	return l
}
