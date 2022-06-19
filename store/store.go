package store

import (
    "errors"
    "github.com/ivo-skyway/my-redis/common"
)

// Store is simple map string -> string
// store is oblivious to transactions
type Store map[string]string

// store state
type store struct {
    s       Store
    deleted []string
}

// NewStore creates new store
func NewStore() *store {
    return &store{s: make(Store, 0)}
}

// Set key - value
func (m *store) Set(key string, value string) {
    m.s[key] = value
}

// Unset key deletes key from the store
func (m *store) Unset(key string) error {
    if _, ok := m.s[key]; ok {
        delete(m.s, key)
        m.deleted = append(m.deleted, key)
        return nil
    }
    return errors.New("key not found")
}

// Get returns value of the key or Nil if not found
func (m *store) Get(key string) string {
    v, found := m.s[key]
    if found {
        return v
    }
    return common.Nil
}

// Freq returns frequency counter for a given value - how many keys have the value
// If value is not found - the counter is 0
func (m *store) Freq(val string) int {
    c := 0
    // loop can be avoided with value -> frequency counter
    for _, v := range m.s {
        if v == val {
            c++
        }
    }
    return c
}

// Cascade copies key/values from source store to the current store
// existing keys are updated, new keys are merged, deleted keys are Unset (may cascade down)
func (m *store) Cascade(s *store) {
    for k, v := range s.s {
        m.s[k] = v
    }
    for _, k := range s.deleted {
        m.Unset(k)
    }
}
