package cmd

import (
    "bufio"
    "errors"
    "fmt"
    "github.com/ivo-skyway/my-redis/common"
    "github.com/ivo-skyway/my-redis/store"
    "os"
    "strings"
)

//parser state
type parser struct {
    db     *store.Stack
    argLen map[string]int
    level  int
    cmd    string
    key    string
    val    string
}

//NewParser creates new instance of the parser
func NewParser() *parser {
    p := parser{level: 0}
    p.argLen = map[string]int{
        common.Get:      1,
        common.Set:      2,
        common.Unset:    1,
        common.Freq:     1,
        common.Begin:    0,
        common.Commit:   0,
        common.Rollback: 0,
        common.End:      0,
    }
    return &p
}

// Run is the main method of the parser
// It reads from standard input line by line, and parses each line for valid command
// Current state contains command cmd (mandatory), optional key, and optional values
// If there are errors in the current line the error is printed on stdout and the line is skipped.
// If there are no errors the execution is delegated to the Stack object
func (p *parser) Run() {
    var err error
    var res string
    scanner := bufio.NewScanner(os.Stdin)

    // stack represents "memory database" interface
    p.db = store.NewStack()

    // read line by line
    for scanner.Scan() {
        // parse and validate
        err = p.Parse(scanner.Text())
        if err == nil {
            if p.cmd == common.End {
                return
            }
            // if valid - execute cmd [key] [value] using Stack object
            res, err = p.db.Exec(p.cmd, p.key, p.val)
            if err == nil && res != "" {
                fmt.Println(res)
            }
            continue
        }
        if err.Error() == "empty" {
            // ignore
            continue
        }
        fmt.Println(err)
        // keep reading on errors
    }
}

// Parse parses one line of the input
// it returns error in case of syntax errors or state errors
func (p *parser) Parse(line string) error {
    line = strings.Trim(line, " \r\n")
    // fmt.Println(line)
    args := strings.Fields(line)
    p.key = ""
    p.val = ""
    n := len(args)
    if n == 0 || args[0] == "" {
        return errors.New("empty")
    }
    p.cmd = strings.ToLower(args[0])
    // check command name and number of arguments
    l, ok := p.argLen[p.cmd]
    if !ok {
        return errors.New("invalid command")
    }
    if l != n-1 {
        return errors.New("invalid number of arguments")
    }
    if l >= 1 {
        p.key = args[1]
    }
    if l >= 2 {
        p.val = args[2]
    }
    return p.Validate()
}

// Validate performs additional state checks after basic syntax check
// it checks for nesting errors like commit without begin
func (p *parser) Validate() error {
    // check nesting
    if p.cmd == common.Begin {
        p.level++
        return nil
    }
    if p.cmd == common.Commit || p.cmd == common.Rollback {
        p.level--
        if p.level < 0 {
            return errors.New("invalid nesting of transactions")
        }
    }
    return nil
}
