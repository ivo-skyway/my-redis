# In Memory Database Test

The following challenge is to implement an in-memory database akin to Redis. The program to be implemented should be a stand alone program of your own writing. Said program should take in commands through standard input (stdin) and should then write it’s responses to standard out (stdout).

## Directives
* Applicants should spend between 30 and 90 minutes to complete the program.
* Use any backend programming language that you would like. We recommend Go (golang).
* Applicants should limit use of 3rd party dependencies or packages. Please stick with the standard library of the language of your choice as much as possible.
* Focus on readability, testing, and completion of acceptance criteria; along with focusing on usability of your finished program.
* Please include instructions in the form of a Readme.md file for how your app is to be installed and used.

## Acceptance Criteria
Your finished program should  accept the following from standard in (stdin).
* `SET name value` - Set the corresponding name to the related value. Neither var names or values can contain spaces.
* `GET name` - Prints out the value of the variable name, or Nil if the variable is not set.
* `UNSET name` - unset the variable name. Meaning if a var is set, then unset - getting said var will return Nil.
* `NUMEQUALTO value` - returns the number of variables that are set to the value in question. Empty results should return 0.

The aforementioned commands will be fed to your program one at a time. Each command will be fed in on its own line. In turn the program should print out ending with a new line when it prints output.

#### Example:

Input

```
SET test-var-name 100
GET test-var-name             100 
UNSET test-var-name
GET test-var-name             Nil
SET test-var-name-1 50
SET test-var-name-2 50
NUMEQUALTO 50                 2
SET test-var-name-2 10
NUMEQUALTO 50                 1
END
```

### Transactions
Along with the aforementioned commands database transactions must also be implemented. Transactions will allow for nested execution of commands. Transactions are supported through the following commands through standard in (stdin).
* `BEGIN` - Opens a new transaction block. Transactions can be nested.
* `ROLLBACK` - Undo all of the commands in the contextual transaction block. If no transaction is open your program should print “NO TRANSACTION”.
* `COMMIT` - Close all open transaction blocks, committing all changes. Print nothing if the transaction committing is successful. Print “NO TRANSACTION” if a transaction block is not open.

#### Examples:
##### Basic Transaction:

Input

```
GET test-var-name             Nil 
BEGIN
SET test-var-name 100
GET test-var-name             100 
COMMIT
GET test-var-name             100
```

##### Nested Transaction:
Input 

```
GET test-var-name             Nil 
BEGIN
SET test-var-name 100
GET test-var-name             100 
BEGIN
SET test-var-name 120
GET test-var-name             120 
BEGIN
SET test-var-name 150
GET test-var-name             150 
ROLLBACK
COMMIT
GET test-var-name             120
```
