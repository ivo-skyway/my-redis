echo Unit tests
go test ./cmd
go test ./store

echo building my-redis
go build .

echo Testing
echo expected lines on stdout: Nil, 100, 222, 120, 150, 120, Nil, Nil
./my-redis < test1.txt
