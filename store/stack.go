package store

import (
    "errors"
    "github.com/ivo-skyway/my-redis/common"
    "strconv"
)

// Stack is the main structure to represent nested transaction state
// It uses slice of store objects to represent each nesting level
type Stack struct {
    level int
    // each level is separate store indexed by level
    stores []store
}

// NewStack creates new instance of Stack
func NewStack() *Stack {
    s := Stack{level: 0,
        stores: make([]store, 0),
    }
    return &s
}

// Len returns number of nested transactions
// default is 0
func (stack *Stack) Len() int {
    return stack.level
}

// Push adds store to the stack and increments the level
func (stack *Stack) Push(store store) {
    stack.stores = append(stack.stores, store)
    stack.level++
}

// Pop removes top level  store from the stack and decrements the level
// proper nesting is already validated - nil is returned on unlikely violations
func (stack *Stack) Pop() *store {
    if stack.level <= 0 {
        return nil
    }
    stack.level--
    store := stack.stores[stack.level]
    stack.stores = stack.stores[:stack.level]
    return &store
}

// Top returns top level store
func (stack *Stack) Top() *store {
    if stack.level <= 0 {
        return nil
    }
    return &stack.stores[stack.level-1]
}

// Exec executes stateful operation over the stack and current top store
// If there is no transaction, implicit first level is created
// Begin creates new transaction level and copies current top to new level (including all key/values)
// Commit merges top level to the previous level and removes the level
// Rollback discards current top store and returns to previous level
func (stack *Stack) Exec(c string, key string, val string) (string, error) {
    m := stack.Top()
    // create first level
    if m == nil {
        m = NewStore()
        stack.Push(*m)
    }
    // command dispatch
    // commands with keys are delegated to current top store
    // begin / commit / rollback change the stack / store state
    // instead of switch we could use map string -> func
    switch c {
    case common.Get:
        return m.Get(key), nil
    case common.Set:
        m.Set(key, val)
        return "", nil
    case common.Unset:
        return "", m.Unset(key)
    case common.Freq:
        return strconv.Itoa(m.Freq(key)), nil

    case common.Begin:
        n := NewStore()
        n.Cascade(m)
        stack.Push(*n)
        return "", nil

    case common.Commit:
        if stack.Len() == 0 {
            return "", errors.New("commit outside transaction")
        }
        // merge with previous
        n := stack.Pop()
        stack.Top().Cascade(n)
        return "", nil

    case common.Rollback:
        if stack.Len() == 0 {
            return "", errors.New("rollback outside transaction")
        }
        // discard
        _ = stack.Pop()
        return "", nil

    default:
        return "", errors.New("not implemented")
    }
}
