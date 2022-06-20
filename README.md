# In Memory Database Test

my-redis is small exercise app simulating Redis set/get key-value store written in go.

See REQUIREMENTS.txt for the requirements

# Installation

Assuming a working installation of Go 1.7 or 1.18 and source in the current directory, simply run 

```
source run-me.sh
```

The script will:

1. Run unit tests;
2. Build my-redis binary in the current directory;
3. Execute my-redis with input from test1.txt file.


## Alternative installation (from github)

```
go install github.com/ivo-skyway/my-redis@latest
my-redis

# or pull the source

git clone https://github.com/ivo-skyway/my-redis.git
cd my-redis
./run-me.sh
```

## Note: 

my-redis commands are case-insensitive - i.e. get and GET will work in the same way, while keys and values are case-sensitive.





