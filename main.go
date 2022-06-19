package main

import "github.com/ivo-skyway/my-redis/cmd"

func main() {
    parser := cmd.NewParser()
    parser.Run()
}
