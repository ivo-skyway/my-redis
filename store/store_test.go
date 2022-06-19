package store

import (
    "github.com/ivo-skyway/my-redis/common"
    "github.com/stretchr/testify/assert"
    "testing"
)

const (
    key  = "foo"
    key2 = "foo2"
    val  = "bar"
    val2 = ""
)

func TestCrud(t *testing.T) {
    m := NewStore()
    assert.NotNil(t, m)
    assert.Equal(t, common.Nil, m.Get(key))
    m.Set(key, val)
    assert.Equal(t, val, m.Get(key))
    m.Set(key, val2)
    assert.Equal(t, val2, m.Get(key))
    m.Unset(key)
    assert.Equal(t, common.Nil, m.Get(key))
}

func TestFreq(t *testing.T) {
    m := NewStore()
    assert.Equal(t, 0, m.Freq(key))
    m.Set(key, val)
    assert.Equal(t, 1, m.Freq(val))
    m.Set(key2, val)
    assert.Equal(t, 2, m.Freq(val))
    m.Set(key, val2)
    assert.Equal(t, 1, m.Freq(val))
    m.Unset(key2)
    assert.Equal(t, 0, m.Freq(val))
}
