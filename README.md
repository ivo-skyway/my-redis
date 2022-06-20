# In Memory Database Test

my-redis is small exercise app simulating Redis set/get key-value store written in go.

See REQUIREMENTS.txt for the requirements

# Installation

Assuming working installation of Go 1.7 or 1.18, simply run 

```
source run-me.sh
```

The script will:

1. run unit tests
2. build my-redis binary
3. execute my-redis with input from test1.txt

## Alternative installation (from github)

```
# install from github
go install github.com/ivo-skyway/my-redis@latest
my-redis

# or pull the source

git clone https://github.com/ivo-skyway/my-redis.git
cd my-redis
./run-me.sh
```

## Note: 
Commands are case-insensitive - i.e. get and GET will work the same way,
keys and values are case-sensitive.





