echo Unit tests
go test ./cmd
go test ./store

echo building my-redis
go build .

echo Testing
echo expected output lines: Nil, 100, 120, 150, 120, Nil
./my-redis < test1.txt
