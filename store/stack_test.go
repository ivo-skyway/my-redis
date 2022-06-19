package store

import (
    "github.com/ivo-skyway/my-redis/common"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestStack(t *testing.T) {
    s := NewStack()
    assert.NotNil(t, s)
    assert.Equal(t, 0, s.Len())
    assert.Nil(t, s.Top())
    m := NewStore()
    m.Set(key, val)
    s.Push(*m)
    assert.NotNil(t, s.Top())
    assert.Equal(t, 1, s.Len())
    m2 := s.Pop()
    assert.NotNil(t, m2)
    assert.Equal(t, m.Get(key), m2.Get(key))
    assert.Equal(t, 0, s.Len())
    assert.Nil(t, s.Top())

    s.Exec("set", key, val)
    v, err := s.Exec("get", key, "")
    assert.NoError(t, err)
    assert.Equal(t, val, v)
}

func TestCommit(t *testing.T) {
    s := NewStack()
    v, err := s.Exec("set", key, val)
    v, err = s.Exec("set", key2, val2)
    assert.NoError(t, err)
    //>>>
    v, err = s.Exec("begin", "", "")
    assert.NoError(t, err)
    v, err = s.Exec("get", key, "")
    assert.Equal(t, val, v)
    s.Exec("set", key, val2)
    v, _ = s.Exec("get", key, "")
    // unset should propagate
    s.Exec("unset", key2, "")
    s.Exec("commit", "", "")
    //<<<
    v, _ = s.Exec("get", key, "")
    assert.Equal(t, val2, v)
    v, _ = s.Exec("get", key2, "")
    assert.Equal(t, common.Nil, v)
}

func TestRollback(t *testing.T) {
    s := NewStack()
    s.Exec("set", key, val)
    s.Exec("set", key2, val2)
    //>>>
    s.Exec("begin", "", "")
    v, _ := s.Exec("get", key, "")
    // unset should not propagate
    s.Exec("unset", key2, "")
    s.Exec("rollback", "", "")
    //<<<
    v, _ = s.Exec("get", key, "")
    assert.Equal(t, val, v)
    v, _ = s.Exec("get", key2, "")
    assert.Equal(t, val2, v)
}

func TestNestedTx(t *testing.T) {
    s := NewStack()
    s.Exec("set", key, "100")
    //>>
    s.Exec("begin", "", "")
    s.Exec("set", key, val)
    //>>>
    s.Exec("begin", "", "")
    s.Exec("set", key, val2)
    v, _ := s.Exec("get", key, "")
    s.Exec("rollback", "", "")
    //<<
    v, _ = s.Exec("get", key, "")
    s.Exec("commit", "", "")
    //<<<

    assert.Equal(t, val, v)
}
