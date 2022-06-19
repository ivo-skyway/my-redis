package cmd

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type testCase struct {
    name    string
    input   string
    wantErr bool
}

func TestParser(t *testing.T) {
    tests := []testCase{
        {"0", "", true},
        {"1", "GET", true},
        {"2", "get foo", false},
        {"3", "Get foo", false},
        {"4", "get foo bar", true},
        {"5", "foo bar", true},
        {"6", "SET", true},
        {"7", "set foo", true},
        {"8", "  set   foo   bar   ", false},
        {"9", "unset foo", false},
        {"10", "UNSET", true},
        {"20", "numequalto", true},
        {"22", "numequalto bar", false},
        {"30", "BEGINxxx", true},
        {"31", "begin foo", true},
        {"32", "begin", false},
        {"33", "commit", false},
        {"34", "commit", true},
        {"35", "begin", false},
        {"36", "begin", false},
        {"40", "rollback", false},
    }

    p := NewParser()
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := p.Parse(tt.input); (err != nil) != tt.wantErr {
                t.Errorf("%v Parse('%v') error = %v, wantErr %v", tt.name, tt.input, err, tt.wantErr)
            }
        })
    }
}

func TestNesting(t *testing.T) {
    p := NewParser()
    err := p.Parse("begin")
    assert.NoError(t, err)
    err = p.Parse("begin")
    assert.NoError(t, err)
    err = p.Parse("rollback")
    assert.NoError(t, err)
    err = p.Parse("commit")
    assert.NoError(t, err)
    err = p.Parse("commit")
    assert.Error(t, err)
}
